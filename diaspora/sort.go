package diaspora
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
  "github.com/Zauberstuhl/go-xml"
  "strings"
  "sort"
)

type XmlOrder struct {
  XMLName xml.Name
  Order []string `xml:"-"`
  Lines []XmlOrderLine `xml:",any"`
}

type XmlOrderLine struct {
  XMLName xml.Name
  Value string `xml:",chardata"`
}

func (order XmlOrder) Len() int { return len(order.Lines) }

func (order XmlOrder) Swap(i, j int) {
  order.Lines[i], order.Lines[j] = order.Lines[j], order.Lines[i]
}

func (order XmlOrder) Less(i, j int) bool {
  for _, o := range order.Order {
    if order.Lines[i].XMLName.Local == o {
      return true
    }
    if order.Lines[j].XMLName.Local == o {
      return false
    }
  }
  // if element order is not
  // avialable ignore it
  return false
}

func (o *XmlOrderLine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
  o.XMLName = start.Name
  return d.DecodeElement(&(o.Value), &start)
}

func FetchEntityOrder(entity []byte) (string, error) {
  var order []string
  var xmlOrder XmlOrder
  err := xml.Unmarshal(entity, &xmlOrder)
  if err != nil {
    return "", err
  }
  for _, line := range xmlOrder.Lines {
    switch line.XMLName.Local {
    case "author_signature":
    case "parent_author_signature":
    default:
      order = append(order, line.XMLName.Local)
    }
  }
  return strings.Join(order, " "), nil
}

func SortByEntityOrder(order string, entity []byte) (sorted []byte, err error) {
  // if we do not require sorting skip it
  if order == "" {
    return entity, err
  }

  var xmlOrder XmlOrder
  xmlOrder.Order = strings.Split(order, " ")
  err = xml.Unmarshal(entity, &xmlOrder)
  if err != nil {
    return sorted, err
  }

  sort.Sort(xmlOrder)

  sorted, err = xml.MarshalIndent(xmlOrder, "", "  ")
  if err != nil {
    return sorted, err
  }
  return
}
