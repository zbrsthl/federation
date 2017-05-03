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
  "encoding/xml"

  //XXX
  "encoding/json"
  "crypto/rsa"
  "crypto/rand"
)

func MagicEnvelope(privkey, handle string, plainXml []byte) (payload []byte, err error) {
  info("plain xml", string(plainXml))
  info("privkey length", len(privkey))
  info("handle", handle)

  data := base64.URLEncoding.EncodeToString(plainXml)
  keyId := base64.URLEncoding.EncodeToString([]byte(handle))

  xmlBody := PrivateEnvMarshal{}
  xmlBody.Data.Type = APPLICATION_XML
  xmlBody.Data.Data = data
  xmlBody.Me = XMLNS_ME
  xmlBody.Encoding = BASE64_URL
  xmlBody.Alg = RSA_SHA256
  xmlBody.Sig.KeyId = keyId

  err = xmlBody.Sign(privkey)
  if err != nil {
    warn(err)
    return
  }

  payload, err = xml.Marshal(xmlBody)
  if err != nil {
    warn(err)
    return
  }
  info("payload", string(payload))
  return
}

func EncryptedMagicEnvelope(privkey, pubkey, handle string, serializedXml []byte) (payload []byte, err error) {
  info("serialized xml", string(serializedXml))
  info("privkey length", len(privkey))
  info("pubkey length", len(pubkey))
  info("handle", handle)

  data := base64.URLEncoding.EncodeToString(serializedXml)
  keyId := base64.URLEncoding.EncodeToString([]byte(handle))

  // encrypted header
  //xmlBody := PrivateMarshal{
  //  Xmlns: XMLNS,
  //  XmlnsMe: XMLNS_ME,
  xmlBodyEnv := PrivateEnvMarshal{
    Me: XMLNS_ME,
    Encoding: BASE64_URL,
    Alg: RSA_SHA256,
  }
  //}
  xmlBodyEnv.Data.Type = APPLICATION_XML
  xmlBodyEnv.Data.Data = data
  xmlBodyEnv.Sig.KeyId = keyId

  err = xmlBodyEnv.Sign(privkey)
  if err != nil {
    warn(err)
    return
  }

  // XXX NOTE move below to header file


  // Generate a new AES key pair
  var (
    aesKeySet Aes
    aesWrapper AesWrapper
  )
  err = aesKeySet.Generate()
  if err != nil {
    warn(err)
    return
  }

  // payload mit aes versch und base64
  payload, err = xml.Marshal(xmlBodyEnv)
  if err != nil {
    warn(err)
    return
  }

  err = aesKeySet.Encrypt(payload)
  if err != nil {
    warn(err)
    return
  }

  //aesWrapper.MagicEnvelope = base64.StdEncoding.EncodeToString([]byte(aesKeySet.Data))
  aesWrapper.MagicEnvelope = aesKeySet.Data

  // aes mit pub key versch

  aesKeySetXml, err := json.Marshal(aesKeySet)
  if err != nil {
    warn(err)
    return
  }

  pubKey, err := ParseRSAPubKey([]byte(pubkey))
  if err != nil {
    warn(err)
    return
  }

  info("aesKeySetXml", string(aesKeySetXml))

  // aes_key
  aesKey, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, aesKeySetXml)
  if err != nil {
    warn(err)
    return
  }

  aesWrapper.AesKey = base64.StdEncoding.EncodeToString(aesKey)


  payload, err = json.Marshal(aesWrapper)
  if err != nil {
    warn(err)
    return
  }


  info("payload", string(payload))
  return
}
