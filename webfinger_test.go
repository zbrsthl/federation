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

import "testing"

func TestWebFinger(t *testing.T) {
  finger := WebFinger{
    Host: "joindiaspora.com",
    Handle: "podmin@joindiaspora.com",
  }
  err := finger.Discovery()
  if err != nil {
    t.Errorf("Some error occured while discovering: %v", err)
  }

  for _, link := range finger.Xrd.Links {
    if link.Rel == WebFingerHcard {
      if link.Href != TEST_HCARD_LINK {
        t.Errorf("Expected to be %s, got %s", TEST_HCARD_LINK, link.Href)
      }
      return
    }
  }
  t.Errorf("Expected hcard link, got nothing")
}
