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

import "github.com/Zauberstuhl/go-xml"

type EntityComment struct {
  XMLName xml.Name `xml:"comment"`
  Author string `xml:"author"`
  CreatedAt Time `xml:"created_at"`
  Guid string `xml:"guid"`
  ParentGuid string `xml:"parent_guid"`
  Text string `xml:"text"`
  AuthorSignature string `xml:"author_signature"`
}

func (e EntityComment) Signature() string {
  return e.AuthorSignature
}

func (e EntityComment) SignatureText(order string) (signatureOrder []string) {
  if order != "" {
    return ExractSignatureText(order, e)
  }
  return []string{
    e.Author,
    e.CreatedAt.String(),
    e.Guid,
    e.ParentGuid,
    e.Text,
  }
}
