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
  "crypto"
  "crypto/rand"
  "crypto/sha256"
  "crypto/rsa"
  "encoding/base64"
  "strings"
)

type SignatureInterface interface {
  Signature() string
  SignatureText(string) []string
  Type() string
}

type Signature struct {
  entity SignatureInterface
  delim string

  Err error
}

func (signature *Signature) New(entity SignatureInterface) *Signature {
  signature.entity = entity
  signature.delim = SignatureAuthorDelimiter
  if entity.Type() == "diaspora" {
    signature.delim = SignatureDelimiter
  } else if entity.Type() == "activitypub" {
    signature.delim = SignatureHTTPDelimiter
  }
  return signature
}

func (signature *Signature) Sign(privKey *rsa.PrivateKey, sig *string) error {
  h := sha256.New()
  h.Write([]byte(strings.Join(
    signature.entity.SignatureText(""), signature.delim)))
  digest := h.Sum(nil)

  rng := rand.Reader
  bytes, err := rsa.SignPKCS1v15(rng, privKey, crypto.SHA256, digest[:])
  if err != nil {
    signature.Err = err
    return err
  }
  *sig = base64.URLEncoding.EncodeToString(bytes)
  return nil
}

func (signature *Signature) Verify(order string, pubKey *rsa.PublicKey) bool {
  sig, err := base64.URLEncoding.DecodeString(signature.entity.Signature())
  if err != nil {
    sig, err = base64.StdEncoding.DecodeString(signature.entity.Signature())
    if err != nil {
      signature.Err = err
      return false
    }
  }
  orderArr := signature.entity.SignatureText(order)
  message := []byte(strings.Join(orderArr, signature.delim))
  hashed := sha256.Sum256(message)

  err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], sig)
  signature.Err = err
  return err == nil
}
