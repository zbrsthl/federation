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
  "github.com/Zauberstuhl/go-xml"
  "encoding/base64"
  "time"
)

type Message struct {
  XMLName xml.Name `xml:"me:env"`
  Me string `xml:"me,attr"`
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

type Entity struct {
  XMLName xml.Name
  // Use custom unmarshaler for xml fetch XMLName
  // and decide which entity to use
  Type string `xml:"-"`
  SignatureOrder string `xml:"-"`
  Data interface{} `xml:"-"`
}

type Time struct {
  time.Time
}

func (m Message) SignatureText() []string {
  return []string{
    m.Data.Data,
    base64.StdEncoding.EncodeToString([]byte(m.Data.Type)),
    base64.StdEncoding.EncodeToString([]byte(m.Encoding)),
    base64.StdEncoding.EncodeToString([]byte(m.Alg)),
  }
}

func (t *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
  e.EncodeElement(t.Format(TIME_FORMAT), start)
  return nil
}

func (t *Time) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
  var value string
  decoder.DecodeElement(&value, &start)
  parse, err := time.Parse(TIME_FORMAT, value)
  if err != nil {
    return err
  }
  *t = Time{parse}
  return nil
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
    content := EntityStatusMessage{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case Reshare:
    content := EntityReshare{}
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
