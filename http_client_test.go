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
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/xml"
)

type Test struct {
  XMLName xml.Name `xml:"AB";json:"-"`
  A string `xml:"A";json:"A"`
  B string `xml:"B";json:"B"`
}

func TestPushToPrivate(t *testing.T) {
  var guid = "1234"

  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    if r.URL.Path != "/receive/users/" + guid {
      t.Errorf("%s", r.URL.Path)
    }
  }))
  defer ts.Close()

  err := PushToPrivate(ts.URL[7:], guid, nil)
  if err != nil {
    t.Errorf("Some error occured while sending: %v", err)
  }

  err = PushToPrivate("", guid, nil)
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}

func TestPushToPublic(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    if r.URL.Path != "/receive/public" {
      t.Errorf("%s", r.URL.Path)
    }
  }))
  defer ts.Close()

  err := PushToPublic(ts.URL[7:], nil)
  if err != nil {
    t.Errorf("Some error occured while sending: %v", err)
  }

  err = PushToPublic("", nil)
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}

func TestFetchJson(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, `{"A":"a","B":"b"}`)
  }))
  defer ts.Close()

  var res Test
  err := FetchJson("GET", ts.URL, nil, &res)
  if err != nil {
    t.Errorf("Some error occured while sending: %v", err)
  }

  if res.A != "a" || res.B != "b" {
    t.Errorf("Expected to be a and b, got %s and %s", res.A, res.B)
  }
}

func TestFetchXml(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, `<AB><A>a</A><B>b</B></AB>`)
  }))
  defer ts.Close()

  var res Test
  err := FetchXml("GET", ts.URL, nil, &res)
  if err != nil {
    t.Errorf("Some error occured while sending: %v", err)
  }

  if res.A != "a" || res.B != "b" {
    t.Errorf("Expected to be a and b, got %s and %s", res.A, res.B)
  }
}
