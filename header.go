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

import(
  "encoding/base64"
  "encoding/json"
  "encoding/xml"
  "crypto/rsa"
  "crypto/rand"
)

func (request *PrivateMarshal) EncryptHeader(authorId string, serializedPubKey []byte) (err error) {
  // Generate the AES key before you start
  // encrypting the plain header
  var aesKeySet Aes
  err = aesKeySet.Generate()
  if err != nil {
    return
  }
  // The actual header
  //
  //  <decrypted_header>
  //     <iv>...</iv>
  //     <aes_key>...</aes_key>
  //     <author_id>one@two.tld</author_id>
  //   </decrypted_header>
  decryptedHeader := XmlDecryptedHeader{
    Iv: aesKeySet.Iv,
    AesKey: aesKeySet.Key,
    AuthorId: authorId,
  }

  decryptedHeaderXml, err := json.Marshal(decryptedHeader)
  if err != nil {
    return err
  }

  err = aesKeySet.Encrypt(decryptedHeaderXml)
  if err != nil {
    return
  }

  aesKeySetXml, err := json.Marshal(aesKeySet)
  if err != nil {
    return err
  }

  pubKey, err := ParseRSAPubKey(serializedPubKey)
  if err != nil {
    return err
  }

  // aes_key
  aesKey, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, aesKeySetXml)
  if err != nil {
    return err
  }

  aesKeyEncoded := base64.StdEncoding.EncodeToString(aesKey)

  header := JsonEnvHeader{
    AesKey: aesKeyEncoded,
    Ciphertext: aesKeySet.Data,
  }

  headerXml, err := json.Marshal(header)
  if err != nil {
    return err
  }

  (*request).EncryptedHeader = base64.StdEncoding.EncodeToString(headerXml)
  return
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

  var aesKeySet Aes
  err = json.Unmarshal(aesKeyJson, &aesKeySet)
  if err != nil {
    return err
  }
  aesKeySet.Data = header.Ciphertext

  ciphertext, err := aesKeySet.Decrypt()
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
