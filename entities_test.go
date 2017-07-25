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
  "testing"
  "encoding/xml"
)

func TestEntitiesUnmarshalXML(t *testing.T) {
  var entity Entity

  var retractionRaw = []byte(`<retraction></retraction>`)
  var profileRaw = []byte(`<profile></profile>`)
  var statusMessageRaw = []byte(`<status_message></status_message>`)
  //var reshareRaw = []byte(`<reshare></reshare>`)
  var commentRaw = []byte(`<comment></comment>`)
  var likeRaw = []byte(`<like></like>`)
  var contactRaw = []byte(`<contact></contact>`)
  var notSupportedRaw = []byte(`<notsupported></notsupported>`)

  err := xml.Unmarshal(retractionRaw, &entity)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if data, ok := entity.Data.(EntityRetraction); !ok {
    t.Errorf("Expected to be 'like', got %v", data)
  }

  err = xml.Unmarshal(profileRaw, &entity)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if data, ok := entity.Data.(EntityProfile); !ok {
    t.Errorf("Expected to be 'profile', got %v", data)
  }

  err = xml.Unmarshal(statusMessageRaw, &entity)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if data, ok := entity.Data.(EntityStatusMessage); !ok {
    t.Errorf("Expected to be 'status_message', got %v", data)
  }

  //err = xml.Unmarshal(reshareRaw, &entity)
  //if err != nil {
  //  t.Errorf("Some error occured while parsing: %v", err)
  //}
  //if data, ok := entity.Data.(EntityStatusMessage); !ok {
  //  t.Errorf("Expected to be 'reshare', got %v", data)
  //}

  err = xml.Unmarshal(commentRaw, &entity)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if data, ok := entity.Data.(EntityComment); !ok {
    t.Errorf("Expected to be 'comment', got %v", data)
  }

  err = xml.Unmarshal(likeRaw, &entity)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if data, ok := entity.Data.(EntityLike); !ok {
    t.Errorf("Expected to be 'like', got %v", data)
  }

  err = xml.Unmarshal(contactRaw, &entity)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }
  if data, ok := entity.Data.(EntityContact); !ok {
    t.Errorf("Expected to be 'contact', got %v", data)
  }

  err = xml.Unmarshal(notSupportedRaw, &entity)
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}
