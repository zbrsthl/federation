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
  "encoding/base64"
  "github.com/Zauberstuhl/go-xml"
  "encoding/json"
  "crypto/rsa"
  "crypto/rand"
  federation "github.com/ganggo/federation"
)

func MagicEnvelope(privKey *rsa.PrivateKey, handle string, plainXml []byte) (payload []byte, err error) {
  federation.Log.Info("MagicEnvelope with", string(plainXml), "for", handle)

  data := base64.URLEncoding.EncodeToString(plainXml)
  keyId := base64.URLEncoding.EncodeToString([]byte(handle))

  xmlBody := Message{}
  xmlBody.Data.Type = APPLICATION_XML
  xmlBody.Data.Data = data
  xmlBody.Me = XMLNS_ME
  xmlBody.Encoding = BASE64_URL
  xmlBody.Alg = RSA_SHA256
  xmlBody.Sig.KeyId = keyId

  var signature federation.Signature
  err = signature.New(xmlBody).Sign(privKey,
    &(xmlBody.Sig.Sig))
  if err != nil {
    federation.Log.Warn(err)
    return
  }

  payload, err = xml.MarshalIndent(xmlBody, "", "  ")
  if err != nil {
    federation.Log.Warn(err)
    return
  }

  federation.Log.Info("MagicEnvelope payload", string(payload))
  return
}

func EncryptedMagicEnvelope(privKey *rsa.PrivateKey, pubKey *rsa.PublicKey, handle string, serializedXml []byte) (payload []byte, err error) {
  federation.Log.Info("EncryptedMagicEnvelope with", string(serializedXml), "for", handle)

  var aesKeySet Aes
  var aesWrapper AesWrapper
  data := base64.URLEncoding.EncodeToString(serializedXml)
  keyId := base64.URLEncoding.EncodeToString([]byte(handle))

  envelope := Message{
    Me: XMLNS_ME,
    Encoding: BASE64_URL,
    Alg: RSA_SHA256,
  }
  envelope.Data.Type = APPLICATION_XML
  envelope.Data.Data = data
  envelope.Sig.KeyId = keyId

  var signature federation.Signature
  err = signature.New(envelope).Sign(privKey,
    &(envelope.Sig.Sig))
  if err != nil {
    federation.Log.Warn(err)
    return
  }

  // Generate a new AES key pair
  err = aesKeySet.Generate()
  if err != nil {
    federation.Log.Warn(err)
    return
  }

  // payload with aes encryption
  payload, err = xml.MarshalIndent(envelope, "", "  ")
  if err != nil {
    federation.Log.Warn(err)
    return
  }

  federation.Log.Info(
    "EncryptedMagicEnvelope payload with aes encryption",
    string(payload),
  )

  err = aesKeySet.Encrypt(payload)
  if err != nil {
    federation.Log.Warn(err)
    return
  }
  aesWrapper.MagicEnvelope = aesKeySet.Data

  // aes with rsa encryption
  aesKeySetXml, err := json.MarshalIndent(aesKeySet, "", "  ")
  if err != nil {
    federation.Log.Warn(err)
    return
  }

  federation.Log.Info("AES key-set XML", string(aesKeySetXml))

  aesKey, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, aesKeySetXml)
  if err != nil {
    federation.Log.Warn(err)
    return
  }
  aesWrapper.AesKey = base64.StdEncoding.EncodeToString(aesKey)

  payload, err = json.MarshalIndent(aesWrapper, "", "  ")
  if err != nil {
    federation.Log.Warn(err)
    return
  }

  federation.Log.Info("EncryptedMagicEnvelope payload", string(payload))
  return
}
