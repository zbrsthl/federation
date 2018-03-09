package helpers
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
  "reflect"
  "strings"
)

func ExractSignatureText(order string, entity interface{}) (signatureOrder []string) {
  orderArr := strings.Split(order, " ")

  typeOf := reflect.TypeOf(entity)
  var mappingFields = map[string]int{}
  for i := 0; i < typeOf.NumField(); i++ {
    field := typeOf.Field(i)
    xmlTag := field.Tag.Get("xml")
    if xmlTag != "" {
      xmlOpts := strings.Split(xmlTag, ",")
      if len(xmlOpts) > 0 && strings.Contains(order, xmlOpts[0]) {
        mappingFields[xmlOpts[0]] = i
      }
    }
  }

  valueOf := reflect.ValueOf(entity)
  for _, orderElem := range orderArr {
    if i, ok := mappingFields[orderElem]; ok {
      if value, ok := valueOf.Field(i).Interface().(Time); ok {
        signatureOrder = append(signatureOrder, string(value))
      }
      if value, ok := valueOf.Field(i).Interface().(string); ok {
        signatureOrder = append(signatureOrder, value)
      }
      if value, ok := valueOf.Field(i).Interface().(bool); ok {
        valueBool := "false"
        if value {
          valueBool = "true"
        }
        signatureOrder = append(signatureOrder, valueBool)
      }
    }
  }
  return
}
