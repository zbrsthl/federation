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
  "crypto/rsa"
  "crypto/x509"
  "crypto/rand"
  "crypto/sha256"
  "encoding/base64"
  "encoding/pem"
  "errors"
  "reflect"
  "strings"
)

func ParseBase64RSAPubKey(encodedKey string) (pubkey *rsa.PublicKey, err error) {
  decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)
  if err != nil {
    return
  }
  return ParseRSAPubKey(decodedKey)
}

func ParseRSAPubKey(decodedKey []byte) (pubkey *rsa.PublicKey, err error) {
  block, _ := pem.Decode(decodedKey)
  if block == nil {
    err = errors.New("Decode public key block is nil!")
    return
  }
  data, err := x509.ParsePKIXPublicKey(block.Bytes)
  if err != nil {
    return
  }
  switch data := data.(type) {
  case *rsa.PublicKey:
    pubkey = data
  default:
    err = errors.New("Wasn't able to parse the public key!")
  }
  return
}

func ParseRSAPrivKey(decodedKey []byte) (privkey *rsa.PrivateKey, err error) {
  block, _ := pem.Decode(decodedKey)
  if block == nil {
    err = errors.New("Decode private key block is nil!")
    return
  }
  privkey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
  if err != nil {
    return
  }
  return
}

func (request *DiasporaUnmarshal) VerifySignature(serialized []byte) error {
  pubkey, err := ParseRSAPubKey(serialized)
  if err != nil {
    warn(err)
    return err
  }

  type64 := base64.StdEncoding.EncodeToString(
    []byte(request.Env.Data.Type),
  )
  encoding64 := base64.StdEncoding.EncodeToString(
    []byte(request.Env.Encoding),
  )
  alg64 := base64.StdEncoding.EncodeToString(
    []byte(request.Env.Alg),
  )

  text := request.Env.Data.Data + "." + type64 + "." + encoding64 + "." + alg64

  h := sha256.New()
  h.Write([]byte(text))
  digest := h.Sum(nil)

  ds, err := base64.URLEncoding.DecodeString(request.Env.Sig.Sig)
  if err != nil {
    return err
  }

  return rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, digest, ds)
}

func AuthorSignature(data interface{}, order, privKey string) (string, error) {
  var text string
  var r = reflect.TypeOf(data)
  var v = reflect.ValueOf(data)

  for _, o := range strings.Split(order, " ") {
    for i := 0; i < r.NumField(); i++ {
      tagList := strings.Split(r.Field(i).Tag.Get("xml"), ",")
      if len(tagList) <= 0 {
        panic("xml struct always requires an xml tag for signatures")
      }
      tag := tagList[0] // the first element is always the xml name

      if tag == o {
        value := v.Field(i).Interface()
        switch v := value.(type) {
        case string:
          text += v + ";"
        case bool:
          positive := "false"
          if v {
            positive = "true"
          }
          text += positive + ";"
        default:
          fatal("Unknown type in AuthorSignature that will break federation!")
        }
      }
    }
  }
  // trim last semicolon
  text = text[:len(text)-1]

  return Sign(text, privKey)
}

func (envelope *MagicEnvelopeMarshal) Sign(privKey string) (err error) {
  type64 := base64.StdEncoding.EncodeToString(
    []byte(envelope.Data.Type),
  )
  encoding64 := base64.StdEncoding.EncodeToString(
    []byte(envelope.Encoding),
  )
  alg64 := base64.StdEncoding.EncodeToString(
    []byte(envelope.Alg),
  )

  text := envelope.Data.Data + "." + type64 +
    "." + encoding64 + "." + alg64
  (*envelope).Sig.Sig, err = Sign(text, privKey)
  return
}


func Sign(text, privKey string) (sig string, err error) {
  privkey, err := ParseRSAPrivKey([]byte(privKey))
  if err != nil {
    return "", err
  }

  h := sha256.New()
  h.Write([]byte(text))
  digest := h.Sum(nil)

  rng := rand.Reader
  sigInByte, err := rsa.SignPKCS1v15(rng, privkey, crypto.SHA256, digest[:])
  if err != nil {
    return "", err
  }

  return base64.StdEncoding.EncodeToString(sigInByte), nil
}
