// Copyright (C) 2022 Specter Ops, Inc.
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
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Token struct {
	accessToken  string
	expiresIn    int
	extExpiresIn int
	expires      time.Time
}

func (s Token) IsExpired() bool {
	return time.Now().After(s.expires.Add(-10 * time.Second))
}

func (s Token) String() string {
	return fmt.Sprintf("Bearer %s", s.accessToken)
}

// The code below uses this weird unmarshalling way for ExpiresIn and ExtExpiresIn
// because the metadata APIs used for obtaining a System Assigned Managed Identity return integer values which are quoted (i.e. {"expires_in" : "86400"})
// while the normal 'login.microsoft.com' APIs return expires_in as a proper integer (i.e. {"expires_in" : 86400}). The previous code was failing
// to parse the response of the metadata APIs since it was a string and not an int. This solves it for both.
func (s *Token) UnmarshalJSON(data []byte) error {
	var res struct {
		AccessToken  string      `json:"access_token"`   // The token to use in calls to Microsoft Graph API
		ExpiresIn    interface{} `json:"expires_in"`     // How long the access token is valid in seconds
		ExtExpiresIn interface{} `json:"ext_expires_in"` // How long the access token is valid in seconds
		TokenType    string      `json:"token_type"`     // Indicates the token type value. The only type currently supported by Azure AD is `bearer`
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return err
	} else {
		// convert ExpiresIn to int
		expiresIn, ok := res.ExpiresIn.(int)
		if !ok {
			// ExpiresIn is not an int
			// try to convert it from string to int
			str, ok := res.ExpiresIn.(string)
			if !ok {
				return nil
			}
			expiresIn, err = strconv.Atoi(str)
			if err != nil {
				return nil
			}
		}

		// convert ExtExpiresIn to int
		extExpiresIn, ok := res.ExtExpiresIn.(int)
		if !ok {
			// ExpiresIn is not an int
			// try to convert it from string to int
			str, ok := res.ExtExpiresIn.(string)
			if !ok {
				return nil
			}
			extExpiresIn, err = strconv.Atoi(str)
			if err != nil {
				return nil
			}
		}

		s.expiresIn = expiresIn
		s.extExpiresIn = extExpiresIn
		s.accessToken = res.AccessToken
		s.expires = time.Now().Add(time.Duration(expiresIn) * time.Second)
		return nil
	}
}
