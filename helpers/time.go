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

import "time"

const TIME_FORMAT = "2006-01-02T15:04:05Z"

type Time string

func (t *Time) New(newTime time.Time) *Time {
  *t = Time(newTime.UTC().Format(TIME_FORMAT))
  return t
}

func (t Time) Time() (time.Time, error) {
  return time.Parse(TIME_FORMAT, string(t))
}

func (t Time) String() string {
  return string(t)
}
