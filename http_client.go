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
  "encoding/xml"
  "errors"
  "io"
  "strings"
)

var timeout = time.Duration(10 * time.Second)

func PushXmlToPublic(host string, body io.Reader, ssl bool) error {
  proto :=  "https://"
  if !ssl {
    proto = "http://"
  }

  req, err := http.NewRequest("POST",
    proto + host + "/receive/public", body)
  if err != nil {
    return err
  }
  req.Header.Set("Content-Type", "application/magic-envelope+xml")

  client := &http.Client{
    Timeout: timeout,
  }
  resp, err := client.Do(req)
  if err != nil {
    if ssl {
      info("Retry with", proto, "on", host, err)
      return PushXmlToPublic(host, body, false)
    }
    return err
  }
  defer resp.Body.Close()

  if !(resp.StatusCode == 200 || resp.StatusCode == 202) {
    return errors.New("PushXmlToPublic results in: " + resp.Status)
  }
  return nil
}

func FetchJson(method, url string, body io.Reader, result interface{}) error {
  req, err := http.NewRequest(method, url, body)
  if err != nil {
    return err
  }
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{
    Timeout: timeout,
  }
  resp, err := client.Do(req)
  if err != nil {
    return err
  }

  err = json.NewDecoder(resp.Body).Decode(result)
  if err != nil {
    return err
  }
  info(result)
  return nil
}

func FetchXml(method, url string, body io.Reader, result interface{}) error {
  var proto string
  if !strings.HasPrefix(url, "http") {
    proto = "https://"
  }
  req, err := http.NewRequest(method, proto + url, body)
  if err != nil {
    return err
  }
  req.Header.Set("Content-Type", "application/xrd+xml")

  client := &http.Client{
    Timeout: timeout,
  }
  resp, err := client.Do(req)
  if err != nil {
    if !strings.HasPrefix(url, "http") {
      return FetchXml(method, "http://" + url, body, result)
    }
    return err
  }
  err = xml.NewDecoder(resp.Body).Decode(result)
  if err != nil {
    return err
  }
  info(result)
  return nil
}

func FetchHtml(method, url string, body io.Reader) (resp *http.Response, err error) {
  var proto string
  if !strings.HasPrefix(url, "http") {
    proto = "https://"
  }
  req, err := http.NewRequest(method, proto + url, body)
  if err != nil {
    return nil, err
  }
  req.Header.Set("Content-Type", "application/xrd+xml")

  client := &http.Client{
    Timeout: timeout,
  }
  resp, err = client.Do(req)
  if err != nil {
    if !strings.HasPrefix(url, "http") {
      return FetchHtml(method, "http://" + url, body)
    }
    return nil, err
  }
  return
}
