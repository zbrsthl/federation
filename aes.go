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
  "encoding/base64"
  "crypto/aes"
  "crypto/cipher"
)

type Aes struct {
  Key string `json:"key,omitempty"`
  Iv string `json:"iv,omitempty"`
  Data string `json:"-"`
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

  ciphertext, err = base64.URLEncoding.DecodeString(a.Data)
  if err != nil {
    return ciphertext, err
  }
  headerText, err := base64.StdEncoding.DecodeString(string(ciphertext))
  if err == nil {
    // depending on the request
    // we have to do it twice
    ciphertext = headerText
  }

  return decryptAES(key, iv, ciphertext)
}

func decryptAES(key, iv, ciphertext []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return ciphertext, err
  }

  mode := cipher.NewCBCDecrypter(block, iv)
  mode.CryptBlocks(ciphertext, ciphertext)

  return ciphertext, nil
}
