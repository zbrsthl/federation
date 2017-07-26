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
  "testing"
  "bytes"
  "log"
  "os"
  "regexp"
)

type TestLogger struct {
  Log string
}
func (t *TestLogger) Println(v... interface{}) {
  t.Log = v[0].(string)
}

func TestSetLogger(t *testing.T) {
  var buf bytes.Buffer
  var expected = "Hello World"

  SetLogger(log.New(&buf, defaultPrefix, log.Lshortfile))
  logger.Info(expected)

  matched, err := regexp.MatchString(expected + "\n", buf.String())
  if err != nil || !matched {
    t.Errorf("Expected to be %s, got %s (%v)", expected, buf.String(), err)
  }

  var testLogger TestLogger
  SetLogger(&testLogger)
  logger.Info(expected)
  if expected != testLogger.Log {
    t.Errorf("Expected to be %s, got %s", expected, testLogger.Log)
  }

  // reset otherwise it will break test output
  SetLogger(log.New(os.Stdout, defaultPrefix, log.Lshortfile))
}
