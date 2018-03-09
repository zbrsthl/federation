package federation
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
  "regexp"
)

func TestHcard(t *testing.T) {
  var hcard Hcard
  err := hcard.Fetch(TEST_HCARD_LINK)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  matched, err := regexp.MatchString(`^[\w\d]+?$`, hcard.Guid)
  if err != nil || !matched {
    t.Errorf("Expected nil and match, got %v and %v", err, matched)
  }

  nickname := "podmin"
  if hcard.Nickname != nickname {
    t.Errorf("Expected to be %s, got %s", nickname, hcard.Nickname)
  }

  matched, err = regexp.MatchString(`^-----BEGIN.+?KEY-----`, hcard.PublicKey)
  if err != nil || !matched {
    t.Errorf("Expected nil and match, got %v and '%s'", err, hcard.PublicKey)
  }

  url := "https://joindiaspora.com/"
  if hcard.Url != url {
    t.Errorf("Expected to be %s, got %s", url, hcard.Url)
  }

  photo := url + "uploads/images/thumb_"
  matched, err = regexp.MatchString(photo + `large_[\w\d]+?\.png`, hcard.Photo)
  if err != nil || !matched {
    t.Errorf("Expected nil and match, got %v and %s", err, hcard.Photo)
  }

  matched, err = regexp.MatchString(photo + `medium_[\w\d]+?\.png`, hcard.PhotoMedium)
  if err != nil || !matched {
    t.Errorf("Expected nil and match, got %v and %s", err, hcard.PhotoMedium)
  }

  matched, err = regexp.MatchString(photo + `small_[\w\d]+?\.png`, hcard.PhotoSmall)
  if err != nil || !matched {
    t.Errorf("Expected nil and match, got %v and %s", err, hcard.PhotoSmall)
  }
}
