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
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/youmark/pkcs8"
)

func Decode(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	defer io.ReadAll(body) // must read all; streaming to the json decoder does not read to EOF making the connection unavailable for reuse
	return json.NewDecoder(body).Decode(v)
}

func NewClientAssertion(tokenUrl string, clientId string, clientCert string, signingKey string, keyPassphrase string) (string, error) {
	if key, err := parseRSAPrivateKey(signingKey, keyPassphrase); err != nil {
		return "", fmt.Errorf("Unable to parse private key: %w", err)
	} else if jti, err := uuid.NewV4(); err != nil {
		return "", fmt.Errorf("Unable to generate JWT ID: %w", err)
	} else if thumbprint, err := x5t(clientCert); err != nil {
		return "", fmt.Errorf("Unable to create X.509 certificate thumbprint: %w", err)
	} else {
		iat := time.Now()
		exp := iat.Add(1 * time.Minute)
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.StandardClaims{
			Audience:  tokenUrl,
			ExpiresAt: exp.Unix(),
			Issuer:    clientId,
			Id:        jti.String(),
			NotBefore: iat.Unix(),
			Subject:   clientId,
			IssuedAt:  iat.Unix(),
		})

		token.Header = map[string]interface{}{
			"alg": "RS256",
			"typ": "JWT",
			"x5t": thumbprint,
		}

		if signedToken, err := token.SignedString(key); err != nil {
			return "", fmt.Errorf("Unable to sign JWT: %w", err)
		} else {
			return signedToken, nil
		}
	}
}

func ParseBody(accessToken string) (map[string]interface{}, error) {
	var (
		body  = make(map[string]interface{})
		parts = strings.Split(accessToken, ".")
	)

	if len(parts) != 3 {
		return body, fmt.Errorf("invalid access token")
	} else if bytes, err := base64.RawStdEncoding.DecodeString(parts[1]); err != nil {
		return body, err
	} else if err := json.Unmarshal(bytes, &body); err != nil {
		return body, err
	} else {
		return body, nil
	}
}

func ParseAud(accessToken string) (string, error) {
	if body, err := ParseBody(accessToken); err != nil {
		return "", err
	} else if aud, ok := body["aud"].(string); !ok {
		return "", fmt.Errorf("invalid 'aud' type: %T", body["aud"])
	} else {
		return strings.TrimSuffix(aud, "/"), nil
	}
}

func parseRSAPrivateKey(signingKey string, password string) (interface{}, error) {
	if decodedBlock, _ := pem.Decode([]byte(signingKey)); decodedBlock == nil {
		return nil, fmt.Errorf("Unable to decode private key")
	} else if key, _, err := pkcs8.ParsePrivateKey(decodedBlock.Bytes, []byte(password)); err != nil {
		return nil, err
	} else {
		return key, nil
	}
}

func x5t(certificate string) (string, error) {
	if decoded, _ := pem.Decode([]byte(certificate)); decoded == nil {
		return "", fmt.Errorf("Unable to decode certificate")
	} else if cert, err := x509.ParseCertificate(decoded.Bytes); err != nil {
		return "", fmt.Errorf("Unable to parse certificate: %w", err)
	} else {
		checksum := sha1.Sum(cert.Raw)
		return base64.StdEncoding.EncodeToString(checksum[:]), nil
	}
}
