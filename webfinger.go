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
  "github.com/Zauberstuhl/go-xml"
  "fmt"
  "errors"
)

type WebFinger struct {
  Host string
  Handle string
  Data WebfingerData
}

type WebfingerData struct {
  // xml
  XMLName xml.Name `xml:"XRD" json:"-"`
  Xmlns string `xml:"xmlns,attr" json:"-"`
  Alias string `xml:"Alias,omitempty" json:"-"`
  // json
  Aliases []string `json:"aliases" xml:"-"`

  Subject string `json:"subject" xml:"Subject,omitempty"`
  Links []WebfingerDataLink `json:"links" xml:"Link"`
}

type WebfingerDataLink struct {
  // xml
  XMLName xml.Name `xml:"Link" json:"-"`

  Rel string `json:"rel" xml:"rel,attr"`
  Type string `json:"type,omitempty" xml:"type,attr"`
  Href string `json:"href,omitempty" xml:"href,attr,omitempty"`
  Template string `json:"template,omitempty" xml:"template,attr,omitempty"`
}

func (w *WebFinger) Discovery() error {
  url := fmt.Sprintf("%s/.well-known/webfinger?resource=acct:%s", w.Host, w.Handle)
  err := FetchJson("GET", url, nil, &w.Data)
  if err != nil {
    url = fmt.Sprintf("%s/webfinger?q=acct:%s", w.Host, w.Handle)
    err = FetchXml("GET", url, nil, &w.Data)
    if err != nil {
      return err
    }
  }

  if len(w.Data.Links) < 1 {
    return errors.New("Webfinger Links missing")
  }
  return nil
}
