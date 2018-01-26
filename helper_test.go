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
  "crypto/rsa"
  "testing"
  "strings"
)

var TEST_AUTHOR = `diaspora_2nd@localhost:3001`

var TEST_HCARD_LINK = "https://joindiaspora.com/hcard/users/db3d5a70c1ee01334cdd5215e5076e52"

var TEST_PRIV_KEY = []byte(`-----BEGIN PRIVATE KEY-----
MIIEpAIBAAKCAQEA5qryP/x8Wtht+VLH/TvLNofTF9B4aj78sYlctQxIuzEmuLc3
nU+EhpeExe9DpJEt6/l5V3NIHymkkNMA7Mu4fWmxdyVMEweY7utTl+wkwdJFFlBV
TI1H9dF91sUB2p5o7irZUdgW8zyZv2d2aHW5cOwR1mysh16IwDSNSRd8quqIkqPC
xYKcVc2cga0jYb5RIpv4mq1MlKrmFwxXws3OK+/5ZtErOJZCyVkz2lf0gqzMIyBD
JT9iU0x0HEq4A9LcFMCtjv/AuQj5y7AV2ehY/J7kxmw/m7sLr3R7YZ/H+pwdiW7n
LyQilUNXRuwLJQY2HmQ1rZaVivwviBqaqeFkpwIDAQABAoIBAHAkr/33zKWGD4Fl
g6FUDqoGQtSTH9fXo5bUx2OmAz4u2Tp4qOssG6wrwftRJbu+cWsGML4ZZ/juj/lw
/EQjjyA54HOiiGfAC9QsSMnVntE0Xy5IBBBhp5iVLu7ZfNtCpJUV8+3cdtvunHj3
3hNPGMcTnmB3GTH+/dEkO4RLjOqyi0GAZEv5NgeKavP6TkkwYuFChh+QiOy0tQSc
P+5pFtDuVn/ytyxxwH3FuML00y6moOishxpWgV4Ik8tcGzpk4VKDKMxtrhecAp+O
Z4fnYTWKUMb9eQhI8K68IM7I6VcMAYIXidHtKrgGoVvAQQk4ih1EpVhF1WwEX9ra
ZHYR28ECgYEA7+x1v48ymXe1wz1d24x8MkYD3oxNHTMhJ+17dfpJjhPlO3q8oWu5
fgK5n5EzGfjlBKdCXK/STIDunz69WnNRs8b2HLyyoFaXiFJVSM7juY0vV8yIfmQN
0GXSfO7Evz0xct8QZfm61wD80LJ9YWuqFKXqGxhBRZRMg3vupwSLcZcCgYEA9h+2
9OxyKI9h0+me2oiibF3ke+78sjUPqqTGb8QrcQp0QenBqO7n4Y9PpXE5X3/yjyqk
GcsL5oaFFxZn+8t6XMZnT6nubQpyY7TFTK3TKyqhfdQPfsXbqgH5KFFisK40TL1a
yx2a0MzYOyGi5eArr+kRhRlt8vSpwnRQmGJz53ECgYEAoLlYPAZy0CpIok021f/r
p0YOC4UTl68L1BKcNXGA2uPrGYhkWwKuVYL/1KxRfmGlEhP2Od8y0ztAH3/JG5HL
NtLfRmsGgrDffFwjc83c8g1pnLiQ65KdSnEbq8PMG4yj1p8l/hpoluW7dxdLNPsK
CiEHjjUWbMUm6KIaQtqhi2sCgYAtK8zsTqj1ALu3pNzexszojqLsjAQcwNhLPUqe
IKbIbF7B6iD83Dv6jc7UUl9xQ45E8FKF2VopyO6MOjSDZejjNhan7Ewx/wTXf8nm
NNDYz04sRctCPRX/sbUEzUsLmi1HGEmdlaVgRPg6ggXforDh7CinAO/I81ZktexE
y2zyQQKBgQDVxzgzD4CcJtpkxwn2qiEIW94Dp/E91Rbr+PUqzV2vSESMLTZyfOh5
NxEAgN/1IVKPsH0Ct7G9TqEgkG7UaHOw4HBtybSUf+gDRw98tb0mk9gGK6m2A1vz
sK4ExAtclJCC8pc2MldVoeGbePUIYIsMMTGMZQ2O6jlINbYaOkFM/w==
-----END PRIVATE KEY-----`)

var TEST_PUB_KEY = []byte(`-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAxREQ/8yKpfxNTET2MIIY
1sIOktYl+6u1O3Ad9GJctzrRxWa/PbxWKX8QmGoWNb9S9tt1oLeRZ4ZGTV51t/ln
4ZOUUhR5Xpx5+jxP6DM5Jf2wqGjzXmQaSgKWsrJ2R0EyP9Ga+j5zk+uMpRXBBt0C
bh2kcPAF7mRZ7EBK+p1a+YKaKhebTJNF3YXSjxNlN6LXvTq+JUg6zTzUWg/6w1K0
jXGpXMVMOSjmRsdfbMjIYg8AXSH4xK2AvqUjWvIrQRPiXa6EjB9468sT0IABzcaY
4iNs3kzMB2aRCcdzMdZplAUH9kA+BOC3sU6VmOCfKCsQ3RyudeWqkbKxJzk/duoG
HaO+3nFgxhubRhO1VNj0EhpGPkVuQH6hhDpNh6sYH0Tq4Fs6gTxQynlwiERqoHW/
bp4TZc8mmSgrnDNuWklM6IeyoSse1lO4ivaSLDuvm8UbZTT1P09QaaCPLI5iLHVT
dP7gLEAYalirOXeZSyzwMAWzWF7NXkVbHjFn4JvJFYKyN+xoBpAwlIXdI1DMOK1H
kJaT7PEyjvar0M9oIqVqEV5hdGqFrlDnW4MvP+wQkmuK+9CygggE/0oefhKYIYc3
zQKne2ejiu9e5cDD5WyVusjeRstj/+9bDlOrQ8X4eh6vmjvd+98B/ZWFCCEkTH5m
DX3H15lu+GelrDpYLThXjnkCAwEAAQ==
-----END PUBLIC KEY-----`)

func TestParseRSAPublicKey(t *testing.T) {
  var (
    err error
    pub interface{}
  )

  pub, err = ParseRSAPublicKey(TEST_PUB_KEY)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if data, ok := pub.(*rsa.PublicKey); !ok {
    t.Errorf("Expected to be '*rsa.PublicKey', got %v", data)
  }

  _, err = ParseRSAPublicKey([]byte(""))
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }

  _, err = ParseRSAPublicKey(TEST_PRIV_KEY)
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}

func TestParseRSAPrivateKey(t *testing.T) {
  var (
    err error
    priv interface{}
  )

  priv, err = ParseRSAPrivateKey(TEST_PRIV_KEY)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if data, ok := priv.(*rsa.PrivateKey); !ok {
    t.Errorf("Expected to be '*rsa.PrivateKey', got %v", data)
  }

  _, err = ParseRSAPrivateKey([]byte(""))
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }

  _, err = ParseRSAPrivateKey(TEST_PUB_KEY)
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}

func TestParseWebfingerHandle(t *testing.T) {
  expectedHandle := "diaspora_2nd"
  handle, err := ParseWebfingerHandle("acct:" + TEST_AUTHOR)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if handle != expectedHandle {
    t.Errorf("Expected to be %s, got %s", expectedHandle, handle)
  }
}

func TestParseStringHelper(t *testing.T) {
  parts, err := parseStringHelper("abc", `^(\w{2})`, 1)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  expected := "ab"
  if parts[1] != expected {
    t.Errorf("Expected to be %s, got %s", expected, parts[1])
  }

  parts, err = parseStringHelper("abc", `^\w{2}`, 1)
  if err == nil {
    t.Errorf("Expected an error, got nil")
  }
}

func TestExractSignatureText(t *testing.T) {
  var entity = EntityComment {
    Author: "1",
    CreatedAt: "2",
    Guid: "3",
    ParentGuid: "4",
    Text: "5",
    AuthorSignature: "6",
  }

  var tests = []struct {
    Order string
    Expected string
  }{
    {
      Order: "author parent_guid",
      Expected: "1;4",
    },
    {
      Order: "author parent_guid text guid",
      Expected: "1;4;5;3",
    },
    {
      Order: "",
      Expected: "",
    },
  }

  for i, test := range tests {
    result := ExractSignatureText(test.Order, entity)
    if strings.Join(result, ";") != test.Expected {
      t.Errorf("#%d: Expected to be '%s', got '%s'", i, test.Expected, strings.Join(result, ";"))
    }
  }
}
