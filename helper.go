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
  "regexp"
  "strings"
  "errors"
)

func FetchEntityOrder(entityXML string) (order string) {
  re := regexp.MustCompile(`<([^/<>]+?)>.+?</[^/<>]+?>`)
  elements := re.FindAllStringSubmatch(entityXML, -1)
  for _, element := range elements {
    if len(element) == 2 {
      switch element[1] {
      case "author_signature":
      case "parent_author_signature":
      default:
        order += element[1] + " "
      }
    }
  }
  if len(order) <= 0 {
    logger.Warn("Entity order is empty")
    return
  }
  return order[:len(order)-1] // trim space
}

// This is a workaround for sorting xml elements. Diaspora requires
// a specific order otherwise the signature check will fail and
// ignore the entity. This should be a TODO since we could implement
// this kind of logic in a custom xml marshaller
func SortByEntityOrder(order string, entity []byte) (sorted []byte) {
  // if we do not require sorting skip it
  if order == "" {
    return entity
  }

  // remove all newline character
  entity = []byte(strings.Replace(string(entity), "\n", "", -1))
  entity = []byte(strings.Replace(string(entity), "\r", "", -1))

  var lines []string
  var linesOffset int
  var closingTag bool
  var entityLen = len(entity)

  for index, c := range entity {
    offset := index + 1
    // abort on last character
    if offset >= entityLen {
      lines = append(lines, string(entity[linesOffset:]))
      break
    }
    // check on "><" open xml tags
    if c == 0x003e && entity[offset] == 0x003c {
      lines = append(lines, string(entity[linesOffset:offset]))
      linesOffset = offset
    }
    // set the closing tag to true if "/" occurs
    if c == 0x002f {
      closingTag = true
    }
    // append the whole xml element after ">" if "/" is true
    if c == 0x003e && closingTag {
      lines = append(lines, string(entity[linesOffset:offset]))
      linesOffset = offset
      closingTag = false
    }
  }

  var start bool = true
  var orderArr = strings.Split(order, " ")
  var sortedXmlElements string

  // sort the elements in the prefered order
  for _, o := range orderArr {
    re := regexp.MustCompile("<"+o+">(.+?)</"+o+">")
    elements := re.FindAllStringSubmatch(string(entity), 1)
    if len(elements) > 0 && len(elements[0]) > 0 {
      sortedXmlElements += elements[0][0]
    }
  }

  // replace only the elements we have to sort
  // with the new sortedXmlElements
  for _, line := range lines {
    var orderMatch bool = false
    for _, o := range orderArr {
      re := regexp.MustCompile("<"+o+">(.+?)</"+o+">")
      if re.Find([]byte(line)) != nil {
        orderMatch = true
        break
      }
    }
    if !orderMatch {
      sorted = append(sorted, []byte(line)...)
    } else {
      if start {
        sorted = append(sorted, []byte(sortedXmlElements)...)
      }
      start = false
    }
  }
  return
}

func ParseWebfingerHandle(handle string) (string, error) {
  parts, err := parseStringHelper(handle, `^acct:(.+?)@.+?$`, 1)
  if err != nil {
    return "", err
  }
  return parts[1], nil
}

func parseStringHelper(line, regex string, max int) (parts []string, err error) {
  r := regexp.MustCompile(regex)
  parts = r.FindStringSubmatch(line)

  if (len(parts) - 1) < max {
    err = errors.New("Cannot extract " + regex + " from " + line)
    return
  }
  return
}
