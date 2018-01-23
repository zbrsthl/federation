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

type signature interface {
  SignatureText() []string
}

type Signature struct {
  delim string
  signatureText []string

  Err error
}

func (signature *Signature) New(sig signature) *Signature {
  signature.signatureText = sig.SignatureText()
  signature.delim = SignatureAuthorDelimiter
  if _, ok := sig.(Message); ok {
    signature.delim = SignatureDelimiter
  }
  return signature
}

func (signature *Signature) Sign(privKey *rsa.PrivateKey, sig *string) error {
  h := sha256.New()
  h.Write([]byte(strings.Join(signature.signatureText, signature.delim)))
  digest := h.Sum(nil)

  rng := rand.Reader
  bytes, err := rsa.SignPKCS1v15(rng, privKey, crypto.SHA256, digest[:])
  if err != nil {
    signature.Err = err
    return err
  }
  *sig = base64.StdEncoding.EncodeToString(bytes)
  return nil
}

func (signature *Signature) Verify(pubKey *rsa.PublicKey, sig []byte) bool {
  message := []byte(strings.Join(signature.signatureText, signature.delim))
  err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, message[:], sig)
  signature.Err = err
  return err == nil
}
