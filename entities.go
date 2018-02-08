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
  "strings"
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

type Time string

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

func (message *Message) Parse() (entity Entity, err error) {
  if !strings.EqualFold(message.Encoding, BASE64_URL) {
    logger.Error("Encoding doesn't match",
      "message", message.Encoding, "lib", BASE64_URL)
    return entity, errors.New("Encoding doesn't match")
  }

  if !strings.EqualFold(message.Alg, RSA_SHA256) {
    logger.Error("Algorithm doesn't match",
      "message", message.Alg, "lib", RSA_SHA256)
    return entity, errors.New("Algorithm doesn't match")
  }

  keyId, err := base64.StdEncoding.DecodeString(message.Sig.KeyId)
  if err != nil {
    logger.Error("Cannot decode signature key ID", "err", err)
    return entity, err
  }
  message.Sig.KeyId = string(keyId)
  logger.Info("Entity sender", message.Sig.KeyId)

  data, err := base64.URLEncoding.DecodeString(message.Data.Data)
  if err != nil {
    logger.Error("Cannot decode message data", "err", err)
    return entity, err
  }
  logger.Info("Entity raw", string(data))

  entity.SignatureOrder, err = FetchEntityOrder(data)
  if err != nil {
    logger.Error("Cannot fetch entity order", "err", err)
    return entity, err
  }
  logger.Info("Entity order", entity.SignatureOrder)

  err = xml.Unmarshal(data, &entity)
  if err != nil {
    logger.Error("Cannot unmarshal data", "err", err)
    return entity, err
  }
  return entity, nil
}

func (t *Time) New(newTime time.Time) *Time {
  *t = Time(newTime.UTC().Format(TIME_FORMAT))
  return t
}

func (t Time) Time() (time.Time, error) {
  return time.Parse(TIME_FORMAT, string(t))
}

func (t Time) String() string {
  return string(t)
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
