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
  "net/http"
  "net/http/httptest"
  "fmt"
)

func TestWebFinger(t *testing.T) {
  tmplBody := `{"subject": "acct:podmin@joindiaspora.com","aliases":[],"links":[`
  body := tmplBody + `{"rel":"http://microformats.org/profile/hcard"}]}`
  failBody := tmplBody + `]}`

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, body)
  }))
  defer ts.Close()

  finger := WebFinger{
    Host: ts.URL[7:], // without protocol
    Handle: "podmin@joindiaspora.com",
  }
  err := finger.Discovery()
  if err != nil {
    t.Errorf("Some error occured while discovering: %v", err)
  }

  ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, failBody)
  }))
  defer ts.Close()

  finger.Host = ts.URL[7:]
  err = finger.Discovery()
  if err == nil {
    t.Errorf("Webfinger discovery should throw an error on invalid links")
  }

  finger.Host = ""
  err = finger.Discovery()
  if err == nil {
    t.Errorf("Webfinger discovery should throw an error on empty host")
  }
}
