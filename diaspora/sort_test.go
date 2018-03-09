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
  "testing"
  "bytes"
)

type TestContent struct{
  Entity string
  Order string
  Expected string
}

var tests = []TestContent {
  {
    Entity: `<poll_participation>
  <guid>f1eb866029f7013487753131731751e9</guid>
  <parent_guid>2a22d6c029e9013487753131731751e9</parent_guid>
  <author>alice@example.org</author>
  <poll_answer_guid>2a22db2029e9013487753131731751e9</poll_answer_guid>
  <author_signature>dT6KbT7kp0bE+s3//ZErxO1wvVIqtD0lY67i81+dO43B4D2m5kjCdzW240eWt/jZmcHIsdxXf4WHNdrb6ZDnamA8I1FUVnLjHA9xexBITQsSLXrcV88UdammSmmOxl1Ac4VUXqFpdavm6a7/MwOJ7+JHP8TbUO9siN+hMfgUbtY=</author_signature>
</poll_participation>`,
    Order: ``,
    Expected: `<poll_participation>
  <guid>f1eb866029f7013487753131731751e9</guid>
  <parent_guid>2a22d6c029e9013487753131731751e9</parent_guid>
  <author>alice@example.org</author>
  <poll_answer_guid>2a22db2029e9013487753131731751e9</poll_answer_guid>
  <author_signature>dT6KbT7kp0bE+s3//ZErxO1wvVIqtD0lY67i81+dO43B4D2m5kjCdzW240eWt/jZmcHIsdxXf4WHNdrb6ZDnamA8I1FUVnLjHA9xexBITQsSLXrcV88UdammSmmOxl1Ac4VUXqFpdavm6a7/MwOJ7+JHP8TbUO9siN+hMfgUbtY=</author_signature>
</poll_participation>`,
  },
  {
    Entity: `<like>
  <positive>true</positive>
  <guid>947a88f029f7013487753131731751e9</guid>
  <parent_type>Post</parent_type>
  <parent_guid>947a854029f7013487753131731751e9</parent_guid>
  <author>alice@example.org</author>
  <author_signature>gk8e+K7XRjVRblv8B8PVOf7BpURbf5HrXO5rmq8D/AkPO7lA0+Akwouu5JGKAHIhPR3dfXVp0o6bIDD+e8gtMYRdDd5IHRfBGNk3WsQecnbhmesHy40Qca/dCQcdcXd5aeWHJKeyUrSAvS55U6VUpk/DK/4IIEZfnr0T9+jM8I0=</author_signature>
</like>`,
    Order: `guid author`,
    Expected: `<like>
  <guid>947a88f029f7013487753131731751e9</guid>
  <author>alice@example.org</author>
  <positive>true</positive>
  <parent_type>Post</parent_type>
  <parent_guid>947a854029f7013487753131731751e9</parent_guid>
  <author_signature>gk8e+K7XRjVRblv8B8PVOf7BpURbf5HrXO5rmq8D/AkPO7lA0+Akwouu5JGKAHIhPR3dfXVp0o6bIDD+e8gtMYRdDd5IHRfBGNk3WsQecnbhmesHy40Qca/dCQcdcXd5aeWHJKeyUrSAvS55U6VUpk/DK/4IIEZfnr0T9+jM8I0=</author_signature>
</like>`,
  },
  {
    Entity: `<comment>
  <author>alice@example.org</author>
  <guid>5c241a3029f8013487763131731751e9</guid>
  <created_at>2016-07-12T00:49:06Z</created_at>
  <parent_guid>c3893bf029e7013487753131731751e9</parent_guid>
  <text>this is a very informative comment</text>
  <author_signature>cGIsxB5hU/94+rmgIg/Z+OUvXVYcY/kMOvc267ybpk1pT44P1JiWfnI26F1Mta62UjzIW/SjeAO0RIsJRguaISLpXX/d5DJCMpePAZaZiagUbdgH/w4L++fXiPxBKkSm+PB4txxmHGN8FHjwEUJFHJ1m3VfU4w2JC8+IBU93eag=</author_signature>
</comment>`,
    Order: `guid text author created_at parent_guid`,
    Expected: `<comment>
  <guid>5c241a3029f8013487763131731751e9</guid>
  <text>this is a very informative comment</text>
  <author>alice@example.org</author>
  <created_at>2016-07-12T00:49:06Z</created_at>
  <parent_guid>c3893bf029e7013487753131731751e9</parent_guid>
  <author_signature>cGIsxB5hU/94+rmgIg/Z+OUvXVYcY/kMOvc267ybpk1pT44P1JiWfnI26F1Mta62UjzIW/SjeAO0RIsJRguaISLpXX/d5DJCMpePAZaZiagUbdgH/w4L++fXiPxBKkSm+PB4txxmHGN8FHjwEUJFHJ1m3VfU4w2JC8+IBU93eag=</author_signature>
</comment>`,
  },
}

func TestFetchEntityOrder(t *testing.T) {
  var testsCopy []TestContent
  testsCopy = make([]TestContent, len(tests))
  copy(testsCopy, tests)

  testsCopy[0].Expected = "guid parent_guid author poll_answer_guid"
  testsCopy[1].Expected = "positive guid parent_type parent_guid author"
  testsCopy[2].Expected = "author guid created_at parent_guid text"

  for i, test := range testsCopy {
    result, err := FetchEntityOrder([]byte(test.Entity))
    if err != nil {
      t.Errorf("#%d: Some error occured while parsing: %v", i, err)
    }

    if result != test.Expected {
      t.Errorf("#%d: Expected to be '%s', got '%s'", i, test.Expected, result)
    }
  }
}

func TestSortByEntityOrder(t *testing.T) {
  for i, test := range tests {
    result, err := SortByEntityOrder(test.Order, []byte(test.Entity))
    if err != nil {
      t.Errorf("#%d: Some error occured while parsing: %v", i, err)
    }

    if bytes.Compare(result, []byte(test.Expected)) != 0 {
      t.Errorf("#%d: Expected to be '%s', got '%s'", i, test.Expected, result)
    }
  }
}
