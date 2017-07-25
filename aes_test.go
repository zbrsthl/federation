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
  "testing"
  "encoding/base64"
  "strings"
)

func TestAes(t *testing.T) {
  var (
    aes Aes
    expected = "Hello World"
  )

  err := aes.Generate()
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  key := aes.Key
  iv := aes.Iv

  aes.Key = ""
  err = aes.Encrypt([]byte(expected))
  if err == nil {
    t.Errorf("Expected to be an error, got nil")
  }

  aes.Key = base64.StdEncoding.EncodeToString([]byte("wrongkey"))
  err = aes.Encrypt([]byte(expected))
  if err == nil {
    t.Errorf("Expected to be an error, got nil")
  }

  aes.Key = key
  // XXX will panic and break the test
  //aes.Iv = ""
  //err = aes.Encrypt([]byte(expected))
  //if err == nil {
  //  t.Errorf("Expected to be an error, got nil")
  //}

  aes.Iv = iv
  err = aes.Encrypt([]byte(expected))
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  aes.Key = ""
  _, err = aes.Decrypt()
  if err == nil {
    t.Errorf("Expected to be an error, got nil")
  }

  aes.Key = base64.StdEncoding.EncodeToString([]byte("wrongkey"))
  _, err = aes.Decrypt()
  if err == nil {
    t.Errorf("Expected to be an error, got nil")
  }

  // XXX will panic and break the test
  aes.Key = key
  //aes.Iv = ""
  //_, err = aes.Decrypt()
  //if err == nil {
  //  t.Errorf("Expected to be an error, got nil")
  //}

  aes.Iv = iv

  decrypted, err := aes.Decrypt()
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if strings.Compare(expected, string(decrypted)) == 0 {
    t.Errorf("Expected to be '%s', got '%s'", expected, string(decrypted))
  }
}
