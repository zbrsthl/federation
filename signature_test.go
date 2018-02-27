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
)

func TestSignatureInterface(t *testing.T) {
  var signature Signature
  var sig string

  priv, err := ParseRSAPrivateKey(TEST_PRIV_KEY)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  err = signature.New(EntityLike{}).Sign(priv, &sig)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if !signature.New(EntityLike{AuthorSignature: sig}).Verify(
      "positive guid parent_guid parent_type author", &priv.PublicKey) {
    t.Errorf("Expected to be a valid signature, got invalid")
  }

  err = signature.New(EntityComment{}).Sign(priv, &sig)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if !signature.New(EntityComment{AuthorSignature: sig}).Verify(
      "author created_at guid parent_guid text", &priv.PublicKey) {
    t.Errorf("Expected to be a valid signature, got invalid")
  }

  sigBytes, err := base64.URLEncoding.DecodeString(sig)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  sig = base64.StdEncoding.EncodeToString(sigBytes)
  if !signature.New(EntityComment{AuthorSignature: sig}).Verify(
      "author created_at guid parent_guid text", &priv.PublicKey) {
    t.Errorf("Expected to be a valid signature, got invalid")
  }
}
