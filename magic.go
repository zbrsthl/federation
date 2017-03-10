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
)

func MagicEnvelope(privkey string, handle, plainXml []byte) (payload []byte, err error) {
  info("plain xml", string(plainXml))
  info("privkey length", len(privkey))
  info("handle", string(handle))

  data := base64.URLEncoding.EncodeToString(plainXml)
  keyId := base64.URLEncoding.EncodeToString(handle)

  xmlBody := DiasporaMarshal{}
  xmlBody.Data.Type = "application/xml"
  xmlBody.Data.Data = data
  xmlBody.Me = "http://salmon-protocol.org/ns/magic-env"
  xmlBody.Encoding = "base64url"
  xmlBody.Alg = "RSA-SHA256"
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
