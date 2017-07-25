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
  "crypto/rsa"
  "testing"
)

func TestParseRSAPubKey(t *testing.T) {
  var (
    err error
    pub interface{}
  )

  pub, err = ParseRSAPubKey(TEST_PUB_KEY)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if data, ok := pub.(*rsa.PublicKey); !ok {
    t.Errorf("Expected to be '*rsa.PublicKey', got %v", data)
  }

  pub, err = ParseRSAPubKey([]byte("INVALID"))
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}

func TestParseRSAPrivKey(t *testing.T) {
  var (
    err error
    priv interface{}
  )

  priv, err = ParseRSAPrivKey(TEST_PRIV_KEY)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if data, ok := priv.(*rsa.PrivateKey); !ok {
    t.Errorf("Expected to be '*rsa.PrivateKey', got %v", data)
  }

  priv, err = ParseRSAPrivKey([]byte("INVALID"))
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}
