// Copyright (C) 2022 The BloodHound Enterprise Team
//
// This file is part of AzureHound.
//
// AzureHound is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// AzureHound is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pipeline

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Result struct {
	Error error
	Ok    interface{}
}

// OrDone provides an explicit cancellation mechanism to ensure the encapsulated and downstream goroutines are cleaned
// up. This frees the caller from depending on the input channel to close in order to free the goroutine, thus
// preventing possible leaks.
func OrDone(done, in interface{}) <-chan interface{} {
	if !isReadable(done) || !isReadable(in) {
		panic(fmt.Errorf("channels must be readable"))
	}
	out := make(chan interface{})

	go func() {
		defer close(out)
		doneCase := reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(done),
		}

		outerCases := []reflect.SelectCase{
			doneCase,
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(in),
			},
		}

		innerCases := []reflect.SelectCase{
			doneCase,
			{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(out),
			},
		}
		for {
			if chosen, item, ok := reflect.Select(outerCases); chosen == 0 || !ok {
				// If received on done then return
				return
			} else {
				if !ok {
					return
				} else {
					innerCases[1].Send = item
					if chosen, _, _ := reflect.Select(innerCases); chosen == 0 || !ok {
						return
					}
				}
			}
		}
	}()
	return out
}

// Mux joins multiple channels and returns a channel as single stream of data.
func Mux(done interface{}, channels ...interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	muxer := func(channel interface{}) {
		defer wg.Done()
		for item := range OrDone(done, channel) {
			out <- item
		}
	}

	wg.Add(len(channels))
	for _, channel := range channels {
		go muxer(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Demux distributes the stream of data from a single channel across multiple channels to parallelize CPU use and I/O
func Demux(done interface{}, in interface{}, size int) []<-chan interface{} {
	// use reflection to dynamically create select statement
	outputs := make([]chan interface{}, size)
	readChans := []<-chan interface{}{}
	for i := range outputs {
		out := make(chan interface{})
		outputs[i] = out
		readChans = append(readChans, out)
	}

	closeOutputs := func() {
		for i := range outputs {
			close(outputs[i])
		}
	}

	cases := make([]reflect.SelectCase, len(outputs))
	for i := range cases {
		cases[i].Dir = reflect.SelectSend
		cases[i].Chan = reflect.ValueOf(outputs[i])
	}
	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(done),
	})

	go func() {
		defer closeOutputs()
		for item := range OrDone(done, in) {
			// send item to exactly once channel or cancel
			for i := range cases {
				if cases[i].Dir == reflect.SelectSend {
					cases[i].Send = reflect.ValueOf(item)
				}
			}

			reflect.Select(cases)
		}
	}()

	return readChans
}

// Tee copies the stream of data from a single channel to zero or more channels
func Tee(done interface{}, in interface{}, outputs ...chan<- interface{}) {
	// use reflection to dynamically create select block
	cases := make([]reflect.SelectCase, len(outputs))
	for i := range cases {
		cases[i].Dir = reflect.SelectSend
	}

	go func() {
		// Need to close outputs when goroutine exits to ensure we avoid deadlock
		defer func() {
			for i := range outputs {
				close(outputs[i])
			}
		}()

		for item := range OrDone(done, in) {
			// setup all possible select cases
			for i := range cases {
				cases[i].Chan = reflect.ValueOf(outputs[i])
				cases[i].Send = reflect.ValueOf(item)
			}

			// send item to each channel no more than once or cancel
			for range cases {
				chosen, _, _ := reflect.Select(cases)
				cases[chosen].Chan = reflect.ValueOf(nil)
			}
		}
	}()
}

func Batch(done interface{}, in interface{}, maxItems int, maxTimeout time.Duration) <-chan []interface{} {
	if !isReadable(done) || !isReadable(in) {
		panic(fmt.Errorf("channels must be readable"))
	}
	out := make(chan []interface{})

	go func() {
		defer close(out)

		doneCase := reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(done),
		}

		itemCase := reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(in),
		}

		timeoutCase := reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(time.After(maxTimeout)),
		}

		var batch []interface{}
		for {
			if chosen, item, ok := reflect.Select([]reflect.SelectCase{doneCase, itemCase, timeoutCase}); chosen == 0 || !ok {
				// Flush and return when canceled or closed
				if len(batch) > 0 {
					out <- batch
					batch = nil
				}
				return
			} else if chosen == 1 {
				// Add to batch
				batch = append(batch, item.Interface())

				// Flush if limit is reached
				if len(batch) >= maxItems {
					out <- batch
					batch = nil
					timeoutCase.Chan = reflect.ValueOf(time.After(maxTimeout))
				}
			} else {
				// Timeout triggered, flush and reset
				if len(batch) > 0 {
					out <- batch
					batch = nil
				}
				timeoutCase.Chan = reflect.ValueOf(time.After(maxTimeout))
			}
		}
	}()

	return out
}

func FormatJson(done interface{}, in interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for item := range OrDone(done, in) {
			if bytes, err := json.Marshal(item); err != nil {
				panic(err)
			} else {
				out <- string(bytes)
			}
		}
	}()

	return out
}

func isReadable(channel interface{}) bool {
	channelType := reflect.TypeOf(channel)
	return channelType.Kind() == reflect.Chan && channelType.ChanDir() != reflect.SendDir
}
