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
  "log"
  "os"
  "strings"
  "bytes"
)

func TestLoggerOutput(t *testing.T) {
  var buf bytes.Buffer
  var msg = "Hello World"

  SetLogger(log.New(&buf, "", log.Lshortfile))
  logger.Info(msg)

  expected := LOG_C_TUR + msg + LOG_C_RESET
  if strings.Contains(expected, buf.String()) {
    t.Errorf("Expected to contain '%s', got '%s'", expected, buf.String())
  }

  buf.Reset()
  logger.Error(msg)
  expected = LOG_C_RED + msg + LOG_C_RESET
  if strings.Contains(expected, buf.String()) {
    t.Errorf("Expected to contain '%s', got '%s'", expected, buf.String())
  }

  buf.Reset()
  logger.Warn(expected)
  expected = LOG_C_YELLOW + msg + LOG_C_RESET
  if strings.Contains(expected, buf.String()) {
    t.Errorf("Expected to contain '%s', got '%s'", expected, buf.String())
  }

  // reset otherwise it will break test output
  SetLogger(log.New(os.Stdout, defaultPrefix, log.Lshortfile))
}
