package diaspora
//
// GangGo Federation Library
// Copyright (C) 2017-2018 Lukas Matt <lukas@zauberstuhl.de>
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
  "bytes"
  "io"
  "encoding/base64"
  "encoding/json"
  "crypto/aes"
  "crypto/rsa"
  "crypto/cipher"
  "crypto/rand"
)

type Aes struct {
  Key string `json:"key,omitempty"`
  Iv string `json:"iv,omitempty"`
  Data string `json:"-"`
}

type AesWrapper struct {
  AesKey string `json:"aes_key"`
  MagicEnvelope string `json:"encrypted_magic_envelope"`
}

func (a *Aes) Generate() error {
  // The key argument should be the AES key,
  // either 16, 24, or 32 bytes to select
  // AES-128, AES-192, or AES-256.
  key := make([]byte, 32)

  _, err := io.ReadFull(rand.Reader, key)
  if err != nil {
    return err
  }
  a.Key = base64.StdEncoding.EncodeToString(key)

  // The IV needs to be unique, but not secure. Therefore it's common to
  // include it at the beginning of the ciphertext.
  iv := make([]byte, aes.BlockSize)
  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    return err
  }
  a.Iv = base64.StdEncoding.EncodeToString(iv)
  return nil
}

func (a *Aes) Encrypt(data []byte) error {
  // CBC mode works on blocks so plaintexts may need to be padded to the
  // next whole block. For an example of such padding, see
  // https://tools.ietf.org/html/rfc5246#section-6.2.3.2.
  padding := aes.BlockSize - len(data)%aes.BlockSize
  padtext := bytes.Repeat([]byte{byte(padding)}, padding)
  data = append(data, padtext...)

  key, err := base64.StdEncoding.DecodeString(a.Key)
  if err != nil {
    return err
  }

  block, err := aes.NewCipher(key)
  if err != nil {
    return err
  }

  ciphertext := make([]byte, len(data))

  iv, err := base64.StdEncoding.DecodeString(a.Iv)
  if err != nil {
    return err
  }

  mode := cipher.NewCBCEncrypter(block, iv)
  mode.CryptBlocks(ciphertext, data)

  a.Data = base64.StdEncoding.EncodeToString(ciphertext)
  return nil
}

func (a Aes) Decrypt() (ciphertext []byte, err error) {
  key, err := base64.StdEncoding.DecodeString(a.Key)
  if err != nil {
    return ciphertext, err
  }

  iv, err := base64.StdEncoding.DecodeString(a.Iv)
  if err != nil {
    return ciphertext, err
  }

  ciphertext, err = base64.StdEncoding.DecodeString(a.Data)
  if err != nil {
    return ciphertext, err
  }

  block, err := aes.NewCipher(key)
  if err != nil {
    return ciphertext, err
  }

  mode := cipher.NewCBCDecrypter(block, iv)
  mode.CryptBlocks(ciphertext, ciphertext)

  return ciphertext, nil
}

func (w AesWrapper) Decrypt(privKey *rsa.PrivateKey) (entityXML []byte, err error) {
  encryptedAesKey, err := base64.StdEncoding.DecodeString(w.AesKey)
  if err != nil {
    return
  }

  decryptedAesKey, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, encryptedAesKey)
  if err != nil {
    return
  }

  var aes Aes
  err = json.Unmarshal(decryptedAesKey, &aes)
  if err != nil {
    return
  }

  aes.Data = w.MagicEnvelope
  entityXML, err = aes.Decrypt()
  if err != nil {
    return
  }
  return
}
