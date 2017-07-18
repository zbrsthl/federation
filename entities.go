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
  "errors"
  "encoding/xml"
  "reflect"
  "strings"
)

type Message struct {
  XMLName xml.Name `xml:"env"`
  Me string `xml:"me,attr"`
  Data struct {
    XMLName xml.Name `xml:"data"`
    Type string `xml:"type,attr"`
    Data string `xml:",chardata"`
  }
  Encoding string `xml:"encoding"`
  Alg string `xml:"alg"`
  Sig struct {
    XMLName xml.Name `xml:"sig"`
    Sig string `xml:",chardata"`
    KeyId string `xml:"key_id,attr,omitempty"`
  }
  Entity Entity `xml:"-"`
}

type Entity struct {
  XMLName xml.Name
  // Use custom unmarshaler for xml fetch XMLName
  // and decide which entity to use
  Type string `xml:"-"`
  SignatureOrder string `xml:"-"`
  Data interface{} `xml:"-"`
}

func (e *Entity) LocalSignatureOrder() (order string) {
  val := reflect.TypeOf(e.Data)
  for i := 0; i < val.NumField(); i++ {
    params := strings.Split(val.Field(i).Tag.Get("xml"), ",")
    if len(params) > 0 {
      switch tagName := params[0]; tagName {
      case "":
      case "-":
      case "author_signature":
      case "parent_author_signature":
      default:
        order += params[0] + " "
      }
    }
  }
  return order[:len(order)-1] // trim space
}

func (e *Entity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
  // NOTE since the encoder ignores the interface type
  // (see https://golang.org/src/encoding/xml/read.go#L377)
  // we have to decode on every single step
  switch local := start.Name.Local; local {
  case Retraction:
    content := EntityRetraction{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case Profile:
    content := EntityProfile{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case StatusMessage:
    fallthrough
  case Reshare:
    content := EntityStatusMessage{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case Comment:
    content := EntityComment{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case Like:
    content := EntityLike{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case Contact:
    content := EntityContact{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  default:
    return errors.New("Entity '" + local + "' not known!")
  }
  return nil
}
