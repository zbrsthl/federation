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
  "regexp"
)

var TEST_MAGIC_DATA = []byte(`<x></x>`)

var TEST_MAGIC_PAYLOAD = []byte(`<me:env xmlns:me="http://salmon-protocol.org/ns/magic-env"><me:data type="application/xml">PHg-PC94Pg==</me:data><me:encoding>base64url</me:encoding><me:alg>RSA-SHA256</me:alg><me:sig key_id="ZGlhc3BvcmFfMm5kQGxvY2FsaG9zdDozMDAx">PIlS0XhUHGqSsoGKP2efeitDKv7uO0ekNoDQPm5lk844muzMPk7iK9t6T3ageqIsl14xmnInDGKlbrM49JiuYB4aFKEwqHAIEj2axCjdm6HRF5mv+2nhVjKISx+AcuKY1Rav9pKQoQqphRm8p9CQr632TK5mkFfBAeGpyJE8I3WNwguy9AozR+ym0b3MrbDhHxpzGxcRAvjyzbRMfvLhOlVKauaIEGDVN6nbBXVSY4hSBYu38+02PyGuyPjjlBJHNIPQXUL9dcSq/LXs/ElwQA2JBLwF6+opQvIBDbjUVkX4spKo/uRNEAlFuR4Ul+bi/Y7+ssoD3DrMHN4Fg2hx5w==</me:sig></me:env>`)

func TestMagicEnvelope(t *testing.T) {
  payload, err := MagicEnvelope(
    string(TEST_PRIV_KEY), string(TEST_AUTHOR), TEST_MAGIC_DATA)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if bytes.Compare(payload, TEST_MAGIC_PAYLOAD) != 0 {
    t.Errorf("Expected to be %s, got %s",
      string(TEST_MAGIC_PAYLOAD), string(payload))
  }
}

func TestEncryptedMagicEnvelope(t *testing.T) {
  payload, err := EncryptedMagicEnvelope(
    string(TEST_PRIV_KEY), string(TEST_PUB_KEY),
    TEST_AUTHOR, TEST_MAGIC_DATA)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  pattern := `aes_key.+?encrypted_magic_envelope`
  matched, err := regexp.Match(pattern, payload)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if !matched {
    t.Errorf("Expected match for pattern '%s', got nothing", pattern)
  }
}
