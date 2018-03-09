package activitypub
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
  "strings"
)

type ActivityBase struct {
  Id string `json:"id"`
  Type string `json:"type"`
}

type ActivityContext struct {
  Context []string `json:"@context"`
  *ActivityBase
}

type Message struct {
  Header map[string]interface{}
  Body map[string]interface{}
}

//func (m Message) Parse() (entity Entity, err error) {
//  if mType, ok := m.Body["type"]; ok {
//    switch strings.ToLower(mType) {
//    case "follow":
//    case "undo":
//    }
//  } else {
//    err = ERROR_MISSING_TYPE
//  }
//  return
//}

// Signature and SignatureText implements HTTP signature
// concatination check https://tools.ietf.org/html/draft-cavage-http-signatures
// for a more detailed explanation

// Signature returns the http signature stored in the request header
func (m Message) Signature() (signature string) {
  if shInt, ok := m.Header["Signature"]; ok {
    sigHeader := shInt.(string)
    sigHeaderSlice := strings.Split(sigHeader, ",")
    for _, elem := range sigHeaderSlice {
      vk := strings.Split(elem, "=")
      if len(vk) > 1 && len(vk[1]) > 0 && vk[0] == "signature" {
        quotedSig := vk[1]
        // signature quotes should be removed before returning
        signature = quotedSig[1:len(quotedSig)-1]
        break
      }
    }
  }
  return
}

// SignatureText returns the ordered signature text required
// for generating or validating valid http signatures
func (m Message) SignatureText() (text []string) {
  if shInt, ok := m.Header["Signature"]; ok {
    sigHeader := shInt.(string)
    sigHeaderSlice := strings.Split(sigHeader, ",")
    for _, elem := range sigHeaderSlice {
      vk := strings.Split(elem, "=")
      if len(vk) > 1 && len(vk[1]) > 0 && vk[0] == "headers" {
        quotedOrder := vk[1]
        // same as above the value is quoted
        order := quotedOrder[1:len(quotedOrder)-1]
        for _, orderElem := range strings.Split(order, " ") {
          // the ordered keys can be stored in
          // lowercases therefore we need strings.Title
          orderElemTitle := strings.Title(orderElem)
          if valInt, ok := m.Header[orderElemTitle]; ok {
            // according to section 2.3 of the RFC
            // we have to omit leading and trailing
            // whitespaces and set all keys to lowercase
            // since this is not guaranteed in quotedOrder
            value := strings.TrimSpace(valInt.(string))
            key := strings.ToLower(orderElem)
            text = append(text, key + ": " + value)
          }
        }
        break
      }
    }
  }
  // according to section 2.1.3 of the RFC
  // the above headers-field is optional
  // in this case we have to fallback to
  // the Date header field
  if len(text) == 0 {
    if dhInt, ok := m.Header["Date"]; ok {
      date := dhInt.(string)
      text = append(text, "date: " + date)
    }
  }
  return
}
