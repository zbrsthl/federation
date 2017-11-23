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

// TODO XML webfinger is deprecated
type WebfingerXml struct {
  XMLName xml.Name `xml:"XRD"`
  Xmlns string `xml:"xmlns,attr"`
  Subject string `xml:"Subject,omitempty"`
  Alias string `xml:"Alias,omitempty"`
  Links []WebfingerXmlLink `xml:"Link"`
}

type WebfingerXmlLink struct {
  XMLName xml.Name `xml:"Link"`
  Rel string `xml:"rel,attr"`
  Type string `xml:"type,attr"`
  Template string `xml:"template,attr,omitempty"`
  Href string `xml:"href,attr,omitempty"`
}

type WebfingerJson struct {
  Subject string `json:"subject"`
  Aliases []string `json:"aliases"`
  Links []WebfingerJsonLink `json:"links"`
}

type WebfingerJsonLink struct {
  Rel string `json:"rel"`
  Type string `json:"type,omitempty"`
  Href string `json:"href,omitempty"`
  Template string `json:"template,omitempty"`
}

func (w *WebFinger) Discovery() error {
  err := FetchXml("GET", w.Host + "/.well-known/host-meta", nil, &w.Xrd)
  if err != nil {
    return err
  }

  if len(w.Xrd.Links) < 1 {
    return errors.New("XRD Link missing")
  }

  for _, link := range w.Xrd.Links {
    if link.Rel == "lrdd" && link.Template != "" {
      err = FetchXml("GET", strings.Replace(
        link.Template, "{uri}", "acct:" + w.Handle, 1), nil, &w.Xrd)
      if err != nil {
        return err
      }
      return nil
    }
  }
  return errors.New("No lrdd rel found in webfinger document!")
}
