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
  "errors"
  "github.com/Zauberstuhl/go-xml"
  "encoding/base64"
  "strings"
  federation "github.com/ganggo/federation"
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

func (m Message) Signature() string {
  return m.Sig.Sig
}

func (m Message) SignatureText(order string) []string {
  return []string{
    m.Data.Data,
    base64.StdEncoding.EncodeToString([]byte(m.Data.Type)),
    base64.StdEncoding.EncodeToString([]byte(m.Encoding)),
    base64.StdEncoding.EncodeToString([]byte(m.Alg)),
  }
}

func (m Message) Type() string { return "diaspora" }

func (message *Message) Parse() (entity Entity, err error) {
  if !strings.EqualFold(message.Encoding, BASE64_URL) {
    federation.Log.Error("Encoding doesn't match",
      "message", message.Encoding, "lib", BASE64_URL)
    return entity, errors.New("Encoding doesn't match")
  }

  if !strings.EqualFold(message.Alg, RSA_SHA256) {
    federation.Log.Error("Algorithm doesn't match",
      "message", message.Alg, "lib", RSA_SHA256)
    return entity, errors.New("Algorithm doesn't match")
  }

  keyId, err := base64.StdEncoding.DecodeString(message.Sig.KeyId)
  if err != nil {
    federation.Log.Error("Cannot decode signature key ID", "err", err)
    return entity, err
  }
  message.Sig.KeyId = string(keyId)
  federation.Log.Info("Entity sender", message.Sig.KeyId)

  data, err := base64.URLEncoding.DecodeString(message.Data.Data)
  if err != nil {
    federation.Log.Error("Cannot decode message data", "err", err)
    return entity, err
  }
  federation.Log.Info("Entity raw", string(data))

  entity.SignatureOrder, err = FetchEntityOrder(data)
  if err != nil {
    federation.Log.Error("Cannot fetch entity order", "err", err)
    return entity, err
  }
  federation.Log.Info("Entity order", entity.SignatureOrder)

  err = xml.Unmarshal(data, &entity)
  if err != nil {
    federation.Log.Error("Cannot unmarshal data", "err", err)
    return entity, err
  }
  return entity, nil
}

func (e *Entity) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
  // NOTE since the encoder ignores the interface type
  // (see https://golang.org/src/encoding/xml/read.go#L377)
  // we have to decode on every single step
  switch local := start.Name.Local; local {
  case federation.Retraction:
    content := EntityRetraction{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case federation.Profile:
    content := EntityProfile{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case federation.StatusMessage:
    content := EntityStatusMessage{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case federation.Reshare:
    content := EntityReshare{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case federation.Comment:
    content := EntityComment{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case federation.Like:
    content := EntityLike{}
    if err := d.DecodeElement(&content, &start); err != nil {
      return err
    }
    (*e).Type = local
    (*e).Data = content
  case federation.Contact:
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
