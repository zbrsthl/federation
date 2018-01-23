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
  "os"
  "log"
)

const (
  LOG_C_TUR = "\033[0;36m"
  LOG_C_RED = "\033[31m"
  LOG_C_YELLOW = "\033[33m"
  LOG_C_RESET = "\033[0m"
)

var (
  logger Logger
  defaultPrefix string
)

func init() {
  pc := make([]uintptr, 1)
  runtime.Callers(3, pc)
  f := runtime.FuncForPC(pc[0])
  file, line := f.FileLine(pc[0])

  defaultPrefix = fmt.Sprintf("%s:%d %s ", file, line, f.Name())
  logger = Logger{
    log.New(os.Stdout, defaultPrefix, log.Lshortfile),
    LOG_C_TUR,
  }
}

type LogWriter interface {
  Println(v... interface{})
}

type Logger struct{
  LogWriter

  Prefix string
}

func SetLogger(writer LogWriter) {
  logger = Logger{writer, LOG_C_TUR}
}

func (l Logger) Info(values... interface{}) {
  values = append(values, []interface{}{""})
  copy(values[1:], values[0:])
  values[0] = l.Prefix
  values = append(values, LOG_C_RESET)
  l.Println(values)
}

func (l Logger) Error(values... interface{}) {
  l.Prefix = LOG_C_RED
  l.Info(values)
}

func (l Logger) Warn(values... interface{}) {
  l.Prefix = LOG_C_YELLOW
  l.Info(values)
}
