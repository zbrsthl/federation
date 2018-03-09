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
  "time"
  "net/http"
  "encoding/json"
  "github.com/Zauberstuhl/go-xml"
  "errors"
  "io"
  "strings"
)

const (
  PROTO_HTTP = "http://"
  PROTO_HTTPS = "https://"
  CONTENT_TYPE_ENVELOPE = "application/magic-envelope+xml"
  CONTENT_TYPE_JSON = "application/json"
  USER_AGENT = "GangGo/v0 (Federation library)"
)

var timeout = time.Duration(10 * time.Second)

func PushToPrivate(host, guid string, body io.Reader) error {
  return push(host, "/receive/users/" + guid, PROTO_HTTPS, CONTENT_TYPE_JSON, body)
}

func PushToPublic(host string, body io.Reader) error {
  return push(host, "/receive/public", PROTO_HTTPS, CONTENT_TYPE_ENVELOPE, body)
}

func push(host, endpoint, proto, contentType string, body io.Reader) error {
  req, err := http.NewRequest("POST", proto + host + endpoint, body)
  if err != nil {
    return err
  }
  req.Header.Set("User-Agent", USER_AGENT)
  req.Header.Set("Content-Type", contentType)

  client := &http.Client{Timeout: timeout}
  resp, err := client.Do(req)
  if err != nil {
    if proto == PROTO_HTTPS {
      Log.Info("Retry with", PROTO_HTTP, "on", host, err)
      return push(host, endpoint, PROTO_HTTP, contentType, body)
    }
    return err
  }
  defer resp.Body.Close()

  if !(resp.StatusCode == 200 || resp.StatusCode == 202) {
    return errors.New("push results in: " + resp.Status)
  }
  return nil
}

func FetchJson(method, url string, body io.Reader, result interface{}) error {
  resp, err := fetch(method, url, "application/json", body)
  if err != nil {
    return err
  }
  return json.NewDecoder(resp.Body).Decode(result)
}

func FetchXml(method, url string, body io.Reader, result interface{}) error {
  resp, err := fetch(method, url, "application/xrd+xml", body)
  if err != nil {
    return err
  }
  return xml.NewDecoder(resp.Body).Decode(result)
}

func FetchHtml(method, url string, body io.Reader) (resp *http.Response, err error) {
  return fetch(method, url, "text/html", body)
}

func fetch(method, url, contentType string, body io.Reader) (*http.Response, error) {
  var proto string
  if !strings.HasPrefix(url, "http") {
    proto = "https://"
  }
  req, err := http.NewRequest(method, proto + url, body)
  if err != nil {
    return nil, err
  }
  req.Header.Set("User-Agent", USER_AGENT)
  req.Header.Set("Content-Type", contentType)

  client := &http.Client{Timeout: timeout}
  resp, err := client.Do(req)
  if err != nil {
    if !strings.HasPrefix(url, "http") {
      return fetch(method, "http://" + url, contentType, body)
    }
    return nil, err
  }
  return resp, nil
}
