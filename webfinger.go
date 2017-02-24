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
  "encoding/xml"
  "strings"
  "errors"
)

type WebFinger struct {
  Host string
  Handle string
  Xrd WebfingerXml
}

type WebfingerXml struct {
  XMLName xml.Name `xml:"XRD"`
  Xmlns string `xml:"xmlns,attr"`
  Subject string `xml:"Subject,omitempty"`
  Links []WebfingerXmlLink `xml:"Link"`
}

type WebfingerXmlLink struct {
  XMLName xml.Name `xml:"Link"`
  Rel string `xml:"rel,attr"`
  Type string `xml:"type,attr"`
  Template string `xml:"template,attr,omitempty"`
  Href string `xml:"href,attr,omitempty"`
}

func (w *WebFinger) Discovery() error {
  err := FetchXml("GET", w.Host +
    "/.well-known/host-meta", nil, &w.Xrd)
  if err != nil {
    return err
  }
  if len(w.Xrd.Links) < 1 {
    return errors.New("XRD Link missing")
  }
  discoveryUrl := strings.Replace(
    w.Xrd.Links[0].Template,
    "{uri}", "acct:" + w.Handle, 1,
  )
  err = FetchXml("GET", discoveryUrl, nil, &w.Xrd)
  if err != nil {
    return err
  }
  return nil
}
