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
  "os"
  "log"
)

const (
  LOG_C_RED = "\033[31m"
  LOG_C_YELLOW = "\033[33m"
  LOG_C_RESET = "\033[0m"
)

var (
  logger Log
  defaultLogger Logger
  defaultPrefix string
)

func init() {
  pc := make([]uintptr, 10)  // at least 1 entry needed
  runtime.Callers(3, pc)
  f := runtime.FuncForPC(pc[0])
  file, line := f.FileLine(pc[0])
  regex, _ := regexp.Compile(`\/([^\/]+?\.go)`)
  result := regex.FindAllStringSubmatch(file, -1)
  if len(result) == 1 {
    file = result[0][1]
  }

  defaultPrefix = fmt.Sprintf("%s:%d %s ", file, line, f.Name())
  defaultLogger = log.New(os.Stdout, defaultPrefix, log.Lshortfile)
}

type Logger interface {
  Println(v... interface{})
}

type Log struct{
  Logger
}

func SetLogger(logger Logger) {
  defaultLogger = logger
}

func (l Log) Info(values... interface{}) {
  defaultLogger.Println(values...)
}

func (l Log) Error(values... interface{}) {
  l.Info(LOG_C_RED, values, LOG_C_RESET)
}

func (l Log) Warn(values... interface{}) {
  l.Info(LOG_C_YELLOW, values, LOG_C_RESET)
}
