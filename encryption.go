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
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
  "encoding/pem"
  "encoding/xml"
  "encoding/json"
  "errors"
  "github.com/revel/revel"
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

func (aes *XmlDecryptedHeader) DecryptAES(ciphertext *[]byte, data string) error {
  key, err := base64.StdEncoding.DecodeString(aes.AesKey)
  if err != nil {
    return err
  }

  iv, err := base64.StdEncoding.DecodeString(aes.Iv)
  if err != nil {
    return err
  }

  *ciphertext, err = base64.URLEncoding.DecodeString(data)
  if err != nil {
    return err
  }
  // diaspora magic do it twice
  *ciphertext, err = base64.StdEncoding.DecodeString(string(*ciphertext))
  if err != nil {
    return err
  }

  *ciphertext, err = DecryptAES(key, iv, *ciphertext)
  if err != nil {
    return err
  }
  return nil
}

func (aes *JsonAesKey) DecryptAES(ciphertext *[]byte, data string) error {
  key, err := base64.StdEncoding.DecodeString(aes.Key)
  if err != nil {
    return err
  }

  iv, err := base64.StdEncoding.DecodeString(aes.Iv)
  if err != nil {
    return err
  }

  *ciphertext, err = base64.StdEncoding.DecodeString(data)
  if err != nil {
    return err
  }

  *ciphertext, err = DecryptAES(key, iv, *ciphertext)
  if err != nil {
    return err
  }
  return nil
}

func DecryptAES(key, iv, ciphertext []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return ciphertext, err
  }

  mode := cipher.NewCBCDecrypter(block, iv)
  mode.CryptBlocks(ciphertext, ciphertext)

  return ciphertext, nil
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

func AuthorSignature(data interface{}, privKey string) (string, error) {
  var text string
  switch entity := data.(type) {
  case *EntityComment:
    text = entity.DiasporaHandle+";"+entity.Guid+";"+
      entity.ParentGuid+";"+entity.Text
  case *EntityLike:
    positive := "false"
    if entity.Positive {
      positive = "true"
    }
    // positive guid parent_guid parent_type author
    text = positive+";"+entity.Guid+";"+entity.ParentGuid+
      ";"+entity.TargetType+";"+entity.DiasporaHandle
  }
  return Sign(text, privKey)
}

func (d *DiasporaMarshal) Sign(privKey string) (err error) {
  type64 := base64.StdEncoding.EncodeToString(
    []byte(d.Data.Type),
  )
  encoding64 := base64.StdEncoding.EncodeToString(
    []byte(d.Encoding),
  )
  alg64 := base64.StdEncoding.EncodeToString(
    []byte(d.Alg),
  )

  text := d.Data.Data + "." + type64 +
    "." + encoding64 + "." + alg64
  (*d).Sig.Sig, err = Sign(text, privKey)
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


func (request *DiasporaUnmarshal) DecryptHeader(serialized []byte) error {
  header := JsonEnvHeader{}
  decryptedHeader := XmlDecryptedHeader{}

  decoded, err := base64.StdEncoding.DecodeString(request.EncryptedHeader)
  if err != nil {
    return err
  }

  err = json.Unmarshal(decoded, &header)
  if err != nil {
    return err
  }

  privkey, err := ParseRSAPrivKey(serialized)
  if err != nil {
    return err
  }

  decoded, err = base64.StdEncoding.DecodeString(header.AesKey)
  if err != nil {
    return err
  }

  aesKeyJson, err := rsa.DecryptPKCS1v15(rand.Reader, privkey, decoded)
  if err != nil {
    return err
  }

  var aesKey JsonAesKey
  err = json.Unmarshal(aesKeyJson, &aesKey)
  if err != nil {
    return err
  }

  var ciphertext []byte
  err = aesKey.DecryptAES(&ciphertext, header.Ciphertext)
  if err != nil {
    return err
  }

  err = xml.Unmarshal(ciphertext, &decryptedHeader)
  if err != nil {
    return err
  }

  (*request).DecryptedHeader = &decryptedHeader
  return nil
}
