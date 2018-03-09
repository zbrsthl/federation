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
  "crypto/rsa"
  "github.com/Zauberstuhl/go-xml"
  federation "github.com/ganggo/federation"
)

func ParseDecryptedRequest(entityXML []byte) (message Message, entity Entity, err error) {
  err = xml.Unmarshal(entityXML, &message)
  if err != nil {
    federation.Log.Error(err)
    return
  }
  entity, err = message.Parse()
  return
}

func ParseEncryptedRequest(wrapper AesWrapper, privKey *rsa.PrivateKey) (message Message, entity Entity, err error) {
  entityXML, err := wrapper.Decrypt(privKey)
  if err != nil {
    federation.Log.Error(err)
    return
  }

  return ParseDecryptedRequest(entityXML)
}
