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
  "testing"
  "reflect"
  "github.com/Zauberstuhl/go-xml"
)

func TestEntitiesUnmarshalXML(t *testing.T) {
  var entity Entity
  var tests = []struct{
    Type string
    Raw []byte
  }{
    {Type: "EntityRetraction", Raw: []byte(`<retraction></retraction>`)},
    {Type: "EntityProfile", Raw: []byte(`<profile></profile>`)},
    {Type: "EntityStatusMessage", Raw: []byte(`<status_message></status_message>`)},
    {Type: "EntityReshare", Raw: []byte(`<reshare></reshare>`)},
    {Type: "EntityComment", Raw: []byte(`<comment></comment>`)},
    {Type: "EntityLike", Raw: []byte(`<like></like>`)},
    {Type: "EntityContact", Raw: []byte(`<contact></contact>`)},
  }

  for i, test := range tests {
    err := xml.Unmarshal(test.Raw, &entity)
    if err != nil {
      t.Errorf("#%d: Some error occured while parsing: %v", i, err)
    }
    name := reflect.TypeOf(entity.Data).Name()
    if test.Type != name {
      t.Errorf("#%d: Expected to be '%s', got '%s'", i, test.Type, name)
    }
    err = xml.Unmarshal(test.Raw[:len(test.Raw)-1], &entity)
    if err == nil {
      t.Errorf("#%d: Expected an error, got nil", i)
    }
  }

  err := xml.Unmarshal([]byte(`<not-supported></not-supported>`), &entity)
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}

func TestEntitiesTimeMarshalAndUnmarshal(t *testing.T) {
  // federation time format
  // 2006-01-02T15:04:05Z
  var time = "2018-01-19T01:32:23Z"
  var rawXml = "<time><CreatedAt>"+time+"</CreatedAt></time>";
  var origTime = struct {
    XMLName xml.Name `xml:"time"`
    CreatedAt Time
  }{}

  err := xml.Unmarshal([]byte(rawXml), &origTime)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if origTime.CreatedAt.String() != time {
    t.Errorf("Expected to be '%s', got '%s'", origTime.CreatedAt.String())
  }

  result, err := xml.Marshal(origTime)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if string(result) != rawXml {
    t.Errorf("Expected to be '%s', got '%s'", result, rawXml)
  }

  timeTime, err := origTime.CreatedAt.Time()
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if timeTime.Format(TIME_FORMAT) != time {
    t.Errorf("Expected to be '%s', got '%s'",
      time, timeTime.Format(TIME_FORMAT))
  }

  // XXX the application server uses time.Now if this happens
  // we should change that and let the library decide what is best
  //err = xml.Unmarshal([]byte("<time><CreatedAt></CreatedAt></time>"), &origTime)
  //if err == nil {
  //  t.Errorf("Expected an error, got nil")
  //}
}
