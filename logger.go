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
  "runtime"
  "regexp"
)

func log(label string, msgs... interface{}) {
  pc := make([]uintptr, 10)  // at least 1 entry needed
  runtime.Callers(3, pc)
  f := runtime.FuncForPC(pc[0])
  file, line := f.FileLine(pc[0])
  regex, _ := regexp.Compile(`\/([^\/]+?\.go)`)
  result := regex.FindAllStringSubmatch(file, -1)
  if len(result) == 1 {
    file = result[0][1]
  }

  fmt.Printf("%s:%d %s ", file, line, f.Name())

  for _, e := range msgs {
    switch msg := e.(type) {
      case error:
        fmt.Printf("[%s] ", label)
        fmt.Print(msg)
      case []error:
        fmt.Println(" \\")
        for _, err := range msg {
          fmt.Printf("\t[%s] ", label)
          fmt.Println(err)
        }
      case string:
        fmt.Printf("[%s] ", label)
        fmt.Print(msg)
      case byte, []byte:
        fmt.Printf("[%s] ", label)
        fmt.Print("%s", msg)
      default:
        fmt.Printf("[%s] ", label)
        fmt.Print(msg)
    }
    fmt.Println()
  }
}

func warn(msgs... interface{}) {
  log("W", msgs)
}

func info(msgs... interface{}) {
  log("I", msgs)
}

func fatal(msgs... interface{}) {
  log("F", msgs)
}
