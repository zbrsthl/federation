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
)

func PreparePublicRequest(body string) (request DiasporaUnmarshal, err error) {
  err = xml.Unmarshal([]byte(body), &request)
  if err != nil {
    warn(err)
    return
  }
  info("public request", request.Env.Data)
  return
}

func (request *DiasporaUnmarshal) Parse(pubkey []byte) (entity Entity, err error) {
  err = request.VerifySignature(pubkey)
  if err != nil {
    warn(err)
    return
  }

  xmlPayload, err := base64.URLEncoding.DecodeString(request.Env.Data.Data)
  if err != nil {
    warn(err)
    return
  }
  return _parse(xmlPayload)
}

func PreparePrivateRequest(body string, privkey []byte) (request DiasporaUnmarshal, err error) {
  err = xml.Unmarshal([]byte(body), &request)
  if err != nil {
    warn(err)
    return
  }

  err = request.DecryptHeader(privkey)
  if err != nil {
    warn(err)
    return
  }
  info("private request to", request.Env.Data)
  return
}

func (request *DiasporaUnmarshal) ParsePrivate(pubkey []byte) (entity Entity, err error) {
  err = request.VerifySignature(pubkey)
  if err != nil {
    warn(err)
    return
  }

  var xmlPayload []byte
  err = request.DecryptedHeader.DecryptAES(&xmlPayload, request.Env.Data.Data)
  if err != nil {
    warn(err)
    return
  }
  return _parse(xmlPayload)
}

func _parse(payload []byte) (entity Entity, err error) {
  err = xml.Unmarshal(payload, &entity)
  if err != nil {
    warn(err)
    return
  }
  return
}
