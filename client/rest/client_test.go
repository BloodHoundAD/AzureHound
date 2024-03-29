// Copyright (C) 2024 Specter Ops, Inc.
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

package rest

import (
	"testing"
)

func TestClosedConnection(t *testing.T) {
	// var s *httptest.Server
	// var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("closing client connections")
	// 	s.CloseClientConnections()
	// }
	// s = httptest.NewServer(h)
	// defer s.Close()

	// res, err := RestClient.Get(gomock.Any(), s.URL)
	// if err == nil {
	// 	t.Fatalf("Something aint right, err should be nil %v", err)
	// }

	// fmt.Printf("res:%v,err:%v", res, err)
}
