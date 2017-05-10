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
  "encoding/json"
  "crypto/rsa"
  "crypto/rand"
)

type MagicEnvelopeMarshal struct {
  XMLName xml.Name `xml:"me:env"`
  Me string `xml:"xmlns:me,attr"`
  Data struct {
    XMLName xml.Name `xml:"me:data"`
    Type string `xml:"type,attr"`
    Data string `xml:",chardata"`
  }
  Encoding string `xml:"me:encoding"`
  Alg string `xml:"me:alg"`
  Sig struct {
    XMLName xml.Name `xml:"me:sig"`
    Sig string `xml:",chardata"`
    KeyId string `xml:"key_id,attr,omitempty"`
  }
}

func MagicEnvelope(privkey, handle string, plainXml []byte) (payload []byte, err error) {
  info("plain xml", string(plainXml))
  info("privkey length", len(privkey))
  info("handle", handle)

  data := base64.URLEncoding.EncodeToString(plainXml)
  keyId := base64.URLEncoding.EncodeToString([]byte(handle))

  xmlBody := MagicEnvelopeMarshal{}
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
  var aesKeySet Aes
  var aesWrapper AesWrapper

  info("serialized xml", string(serializedXml))
  info("privkey length", len(privkey))
  info("pubkey length", len(pubkey))
  info("handle", handle)

  data := base64.URLEncoding.EncodeToString(serializedXml)
  keyId := base64.URLEncoding.EncodeToString([]byte(handle))

  envelope := MagicEnvelopeMarshal{
    Me: XMLNS_ME,
    Encoding: BASE64_URL,
    Alg: RSA_SHA256,
  }
  envelope.Data.Type = APPLICATION_XML
  envelope.Data.Data = data
  envelope.Sig.KeyId = keyId

  err = envelope.Sign(privkey)
  if err != nil {
    warn(err)
    return
  }

  // Generate a new AES key pair
  err = aesKeySet.Generate()
  if err != nil {
    warn(err)
    return
  }

  // payload with aes encryption
  payload, err = xml.Marshal(envelope)
  if err != nil {
    warn(err)
    return
  }

  info("payload, err = xml.Marshal(envelope) ", string(payload))

  err = aesKeySet.Encrypt(payload)
  if err != nil {
    warn(err)
    return
  }
  //aesWrapper.MagicEnvelope = base64.StdEncoding.EncodeToString([]byte(aesKeySet.Data))
  aesWrapper.MagicEnvelope = aesKeySet.Data

  // aes with rsa encryption
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
