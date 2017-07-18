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
  "encoding/xml"
  "encoding/base64"
  "strings"
)

func ParseDecryptedRequest(entityXML []byte) (message Message, err error) {
  err = xml.Unmarshal(entityXML, &message)
  if err != nil {
    fatal(err)
    return
  }

  if !strings.EqualFold(message.Encoding, BASE64_URL) {
    fatal(err)
    return
  }

  if !strings.EqualFold(message.Alg, RSA_SHA256) {
    fatal(err)
    return
  }

  keyId, err := base64.StdEncoding.DecodeString(message.Sig.KeyId)
  if err != nil {
    fatal(err)
    return
  }
  message.Sig.KeyId = string(keyId)

  data, err := base64.URLEncoding.DecodeString(message.Data.Data)
  if err != nil {
    fatal(err)
    return
  }


  var entity = Entity{
    SignatureOrder: FetchEntityOrder(string(data)),
  }
  err = xml.Unmarshal(data, &entity)
  if err != nil {
    fatal(err)
    return
  }
  message.Entity = entity
  return
}

func ParseEncryptedRequest(wrapper AesWrapper, privkey []byte) (message Message, err error) {
  entityXML, err := wrapper.Decrypt(privkey)
  if err != nil {
    fatal(err)
    return
  }

  return ParseDecryptedRequest(entityXML)
}
