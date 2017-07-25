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
  "errors"
)

type EntityComment struct {
  XMLName xml.Name `xml:"comment"`
  Author string `xml:"author"`
  CreatedAt string `xml:"created_at"`
  Guid string `xml:"guid"`
  ParentGuid string `xml:"parent_guid"`
  Text string `xml:"text"`
  AuthorSignature string `xml:"author_signature"`
  ParentAuthorSignature string `xml:"parent_author_signature"`
}

func (e *EntityComment) SignatureOrder() string {
  return "author created_at guid parent_guid text"
}

func (e *EntityComment) AppendSignature(privKey []byte, order string, typ int) error {
  signature, err := AuthorSignature(*e, order, privKey)
  if err != nil {
    return err
  }

  if AuthorSignatureType == typ {
    (*e).AuthorSignature = signature
  } else if ParentAuthorSignatureType == typ {
    (*e).ParentAuthorSignature = signature
  } else {
    return errors.New("Unsupported author signature type!")
  }
  return nil
}
