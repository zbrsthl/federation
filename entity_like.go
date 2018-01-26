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

type EntityLike struct {
  XMLName xml.Name `xml:"like"`
  Positive bool `xml:"positive"`
  Guid string `xml:"guid"`
  ParentGuid string `xml:"parent_guid"`
  ParentType string `xml:"parent_type"`
  Author string `xml:"author"`
  AuthorSignature string `xml:"author_signature"`

  // store relayable signature order
  SignatureOrder string `xml:"-"`
}

func (e EntityLike) Signature() string {
  return e.AuthorSignature
}

func (e EntityLike) SignatureText(order string) []string {
  if order != "" {
    return ExractSignatureText(order, e)
  }

  positive := "false"
  if e.Positive {
    positive = "true"
  }
  return []string{
    positive,
    e.Guid,
    e.ParentGuid,
    e.ParentType,
    e.Author,
  }
}
