package federation
//
// GangGo Diaspora Federation Library
// Copyright (C) 2017 Lukas Matt <lukas@zauberstuhl.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

import (
  "regexp"
  "crypto/rsa"
  "crypto/x509"
  "encoding/pem"
  "errors"
  "reflect"
  "strings"
)

func ExractSignatureText(order string, entity interface{}) (signatureOrder []string) {
  orderArr := strings.Split(order, " ")

  typeOf := reflect.TypeOf(entity)
  var mappingFields = map[string]int{}
  for i := 0; i < typeOf.NumField(); i++ {
    field := typeOf.Field(i)
    xmlTag := field.Tag.Get("xml")
    if xmlTag != "" {
      xmlOpts := strings.Split(xmlTag, ",")
      if len(xmlOpts) > 0 && strings.Contains(order, xmlOpts[0]) {
        mappingFields[xmlOpts[0]] = i
      }
    }
  }

  valueOf := reflect.ValueOf(entity)
  for _, orderElem := range orderArr {
    if i, ok := mappingFields[orderElem]; ok {
      if value, ok := valueOf.Field(i).Interface().(Time); ok {
        signatureOrder = append(signatureOrder, string(value))
      }
      if value, ok := valueOf.Field(i).Interface().(string); ok {
        signatureOrder = append(signatureOrder, value)
      }
      if value, ok := valueOf.Field(i).Interface().(bool); ok {
        valueBool := "false"
        if value {
          valueBool = "true"
        }
        signatureOrder = append(signatureOrder, valueBool)
      }
    }
  }
  return
}

func ParseRSAPublicKey(decodedKey []byte) (*rsa.PublicKey, error) {
  block, _ := pem.Decode(decodedKey)
  if block == nil {
    return nil, errors.New("Decode public key block is nil")
  }
  data, err := x509.ParsePKIXPublicKey(block.Bytes)
  if err != nil {
    return nil, err
  }
  if pubKey, ok := data.(*rsa.PublicKey); ok {
    return pubKey, nil
  }
  return nil, errors.New("Wasn't able to parse the public key!")
}

func ParseRSAPrivateKey(decodedKey []byte) (*rsa.PrivateKey, error) {
  block, _ := pem.Decode(decodedKey)
  if block == nil {
    return nil, errors.New("Decode private key block is nil!")
  }
  return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func ParseWebfingerHandle(handle string) (string, error) {
  parts, err := parseStringHelper(handle, `^acct:(.+?)@.+?$`, 1)
  if err != nil {
    return "", err
  }
  return parts[1], nil
}

func parseStringHelper(line, regex string, max int) (parts []string, err error) {
  r := regexp.MustCompile(regex)
  parts = r.FindStringSubmatch(line)

  if (len(parts) - 1) < max {
    err = errors.New("Cannot extract " + regex + " from " + line)
    return
  }
  return
}
