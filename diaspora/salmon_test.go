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
  "fmt"
)

func TestParseDecryptedRequest(t *testing.T) {
  var msgData = `PHN0YXR1c19tZXNzYWdlPgogIDxhdXRob3I-ZGlhc3BvcmFfMm5kQGxvY2FsaG9zdDozMDAxPC9hdXRob3I-CiAgPGd1aWQ-ZmUyZDJhODA1MzQ4MDEzNWQwOGY1Mjk2ZjJlNzQ0N2I8L2d1aWQ-CiAgPGNyZWF0ZWRfYXQ-MjAxNy0wNy0yNVQwOToyNDozM1o8L2NyZWF0ZWRfYXQ-CiAgPHByb3ZpZGVyX2Rpc3BsYXlfbmFtZS8-CiAgPHRleHQ-cGluZzwvdGV4dD4KICA8cHVibGljPmZhbHNlPC9wdWJsaWM-Cjwvc3RhdHVzX21lc3NhZ2U-`
  var tmpl = `<?xml version="1.0" encoding="UTF-8"?><me:env xmlns:me="http://salmon-protocol.org/ns/magic-env"><me:data type="application/xml">%s</me:data><me:encoding>%s</me:encoding><me:alg>%s</me:alg><me:sig key_id="%s">NbuD4kERZzXPFRORH4NOcr7EAij-dWKTCG0eBBGZObN3Aic0lMAZ_rLU7o6PLOH9Q6p6dyneYjUjSu07vtI5Jy_N2XQpKUni3fUWxfDNgfMo26XKmxdJ5S2Gp1ux1ToO3FY9RByTZw5HZRpOBAfRSgttTgiY5_Yu5D-BLcEm_94R6FMWRniQXrMAt8hU9qCNSuVQlUKtuuy8qJXu6Z21VhI9lAT7wIALlR9UwIgz0e6UG9S9sU95f_38co0ibD1KbQpBd8c_lu5vCVIqlEe_Fa_xYZupMLaU8De-wzoBpBgqR65mRtUQTu2jP-Qxa3aXrANHxweIbnYfpZ5QcNA50hfyVJJSolczDSlDljTunEmHmWNaS3J7waEQsIDFATPFy6H5leRPpSzebXYca4T-EiapPP-mn41Vs3VKIdUXOHus_HcTPWRVT-Vr-yt7byFYEanb5b5lQ_IHcI0oyqX7RrVJid6UsBtwxwkX0FSc1cZgLhBQUgxBsUh5MNte-WZJv_6c9rHyNsH3rn9YEZp431P9GCe8gNdLY9bFQ1pYS9BxOAS2enu3yVpWpWRechiR7D__HC4-Hw2MHfSSmBQTxq5oO01_efEHB8XxWF85XYNT6_icXf3ZsTxkURT9HlHapkFwL7TlO5gPUZZVJt9f6kn9HoGQ56MX2n46KdKKid8=</me:sig></me:env>`
  tests := [][]string{
    []string{msgData, "base32url", "RSA-SHA256", "ZGlhc3BvcmFfMm5kQGxvY2FsaG9zdDozMDAx"},
    []string{msgData, "base64url", "RSA-SHA128", "ZGlhc3BvcmFfMm5kQGxvY2FsaG9zdDozMDAx"},
    []string{msgData, "base64url", "RSA-SHA256", "not valid at all"},
    []string{"not valid", "base64url", "RSA-SHA256", "ZGlhc3BvcmFfMm5kQGxvY2FsaG9zdDozMDAx"},
  }

  for i, test := range tests {
    _, entity, err := ParseDecryptedRequest([]byte(fmt.Sprintf(
      tmpl, test[0], test[1], test[2], test[3],
    ))); if err == nil {
      t.Errorf("#%d: Expected to be an error, got %+v, with %+v", i, entity, test)
    }
  }

  _, _, err := ParseDecryptedRequest([]byte("<broken></broken"))
  if err == nil {
    t.Errorf("Expected to be an error, got nil")
  }

  message, entity, err := ParseDecryptedRequest([]byte(fmt.Sprintf(
    tmpl, msgData, "base64url", "RSA-SHA256", "ZGlhc3BvcmFfMm5kQGxvY2FsaG9zdDozMDAx",
  ))); if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  parseEntityRequest(t, entity)
  parseMessageRequest(t, message)
}

func TestParseEncryptedRequest(t *testing.T) {
  var wrapper = AesWrapper{
    AesKey: `SOcOINIGdXX5QslLmPKXRllOXgBb3HhcUS78BNsCn0hGff0WdobvmgITvd6v+FiQOqIx5RUU2EH7Woh1KEc1yOn0AamyaJrbgDt9wl4az3hxacJsBd+xpcUca8niIGzCZuoJdtnENlAAUU1mKpT7R0Pikdd0/3bjoS6FN2dX1frBx2hlMvHsnADazgfckmK+53ftWWBJ7cAWNJtYOj3MphnipInuOZ3JH43rjpc6QLYRxQ9cvA9cMV+zJr5PKy1QXafILp/55K5YMarwSQPUpN+fZeufMkRFvDPFxOFXrng2V6SvScKOAt95Q5zUf/RDItZaq3smRueCsrJQXoQPJA==`,
    MagicEnvelope: `YLoRPK39sEWfAESJ5IknDfAyZXms4c9Us9K5zoW+8Z4HvR2MUpZOfP5TDFqZDgORp/dY4T6AlSFmu6VXAQsW1ajV7YDAdzhagId/c63kkDgeC1kDu1Ny1xFX23W97hwNizIHIK5uGpw44KgASgy3tXCLIe/JTCB0ykaqSJf0lJ2RO4PnXnz4m3z52WxmsDacBl8Cg0NfHwOgJAX+NGy9so9ECJckCGjCHsJrmVVt/Hp3/MftDFFNFKQ1COaBMFIa2l3qop4TW4yVsBhG5nVBM81+uurr3UZmdN38pmWWgWyTOoHsSK91WuaUbDDRMVb9G+adAIbcb6LXz8qoX68DnQE+7jh8eH7tuyog3+n67JLW3x34iNEpqr/fVdg/DQuujCeXI/OyHH8b6dmVnNMUzOwd4WPeGy6SclDS7s5bjYxRdWoH63d4QQMqrHWmH0RQMBwm8Yj7cRsFKSvQzqQAND+Hfk9XO75QLuBVa2DW/18r3qBiNIaMw1+6VU6XeDoH2vQJU01vK6uR4vXuOXRF0ZawaWS2AHMvDmhMEtGBfQaZpBXDMWTTbyq99vS04tL1AZXs+U2UyxovUSRgLUgqK5qya+MNDn4VcdAdj5tcQ3kciFZyxXDfZ7p0ir7XeTPZHeI9PnjqJEwOuVADmdUTUCx4k5T49YG1op7W14lDuw1DNC3d0KjNYwEFhyMwrgMq9x4Wv5Tsp3avrggdXuHdfJjiCgXKXor9qIqZAG1lD6lyOUV/BBvmEEv6x/hYMmUOmL/rlTNlv11HeynLGyFq6T+6A94Ea9Th+XVU+eYKIKoXs5bJ154s3o1J94zItmMArVV52BUOK8Kpg0OaotEjaMHTT3fOmHzErZTLc63Tb7h21A+DpiLERC+zFEdYs9ifkPDvCyB+TOO1AqRISSGaltaxxIfvo/XQPjP6yjWFcJgkJVy4Lg+nAvfCzZeTAMI8otzDpG6fCIfpg3BJK/5MObAR6rLwT9EvvfWTHNKZR3SNKyIxYjNJr1dwhgB9jGWFRHA+bgDnV2yQPB+Vp+YE2E60jJlnTeSwEfTxPhz5ueQ/rsbPoI2BuqgwZKYkOQ6vjuLfo81EIhQeUfrid5oCCAGDOFjerRaaLVM6iqKeWaAyVujTzGdYOB2tCkFhB9rju6hAHP5ycHp/utOQssDO0LDThosH98fVyVlmC2L+ZIHZ2B+n3OiFz/E1hJ7EOJI2P+jQhwb2uwKVkHgywGIdRTGdzZ1DzXLDcwF5+lGS1wEikmnh1nMEovNNATCFp7qMLK01EywgrLFuF75T00jHld2eU/K/6KhUYi0SJSGDjCx5DxR48xVBRKrn8dCMBC+kbyQ/1pMtM5vg05uPZ963gzWw4uMr/StowT347H/WuuyKIeOSM4RPi+vB1QN2oINVecq3ZKU8U1xKWvuV2M1j/V2OYVcz5NP9Z9nf0fkNqpmek0D4epU4/6bwAuw1YAa6eEvIt1yGrrmh81wumBOHkk614bcyljVut1JAdreJsAj9n7FBIO4UvFO9zmPF2PIRZ+dxt6uNvIItlR659PpIBtoQiLW803SsLcwMqQ+Opg+eFgAB+qvUjZ4F9ZhFHHMRbMJlfu/ezatg7mL9VidKwLavHbCvgmz6ckRdU6m6aQsnRIdWKHU43sjCoPRnSwgMz/D4vNh5F8A43o4RICOUUUT2jADphblbpos61tZyuhR+1uTLEmVrpGfSigj1A11ByHIPtlFxpN0/D2iJkO3OuQ==`,
  }

  privKey, err := ParseRSAPrivateKey(TEST_PRIV_KEY)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  message, entity, err := ParseEncryptedRequest(wrapper, privKey)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  parseEntityRequest(t, entity)
  parseMessageRequest(t, message)

  w := wrapper
  w.AesKey = ""
  _, _, err = ParseEncryptedRequest(w, privKey)
  if err == nil {
    t.Errorf("Expected to be an error, got nil")
  }

  w = wrapper
  w.MagicEnvelope = ""
  _, _, err = ParseEncryptedRequest(w, privKey)
  if err == nil {
    t.Errorf("Expected to be an error, got nil")
  }
}

func parseEntityRequest(t *testing.T, entity Entity) {
  // {XMLName:{Space: Local:} Type:status_message SignatureOrder:author guid created_at text public Data:{XMLName:{Space: Local:status_message} Author:diaspora_2nd@localhost:3001 Guid:fe2d2a8053480135d08f5296f2e7447b CreatedAt:2017-2018-07-25 09:24:33 +0000 UTC ProviderName: Text:ping Photo:<nil> Location:<nil> Poll:<nil> Public:false Event:<nil>}}
  if entity.Type != "status_message" {
    t.Errorf("Expected type string 'status_message', got '%s'", entity.Type)
  }

  if entity.SignatureOrder != "author guid created_at provider_display_name text public" {
    t.Errorf("Expected an order of 'author guid created_at text public', got %s", entity.SignatureOrder)
  }

  if _, ok := entity.Data.(EntityStatusMessage); !ok {
    t.Errorf("Expected the struct type EntityStatusMessage, got %+v", entity.Data)
  }
}

func parseMessageRequest(t *testing.T, message Message) {
  var data = `PHN0YXR1c19tZXNzYWdlPgogIDxhdXRob3I-ZGlhc3BvcmFfMm5kQGxvY2FsaG9zdDozMDAxPC9hdXRob3I-CiAgPGd1aWQ-ZmUyZDJhODA1MzQ4MDEzNWQwOGY1Mjk2ZjJlNzQ0N2I8L2d1aWQ-CiAgPGNyZWF0ZWRfYXQ-MjAxNy0wNy0yNVQwOToyNDozM1o8L2NyZWF0ZWRfYXQ-CiAgPHByb3ZpZGVyX2Rpc3BsYXlfbmFtZS8-CiAgPHRleHQ-cGluZzwvdGV4dD4KICA8cHVibGljPmZhbHNlPC9wdWJsaWM-Cjwvc3RhdHVzX21lc3NhZ2U-`

  var sig = `NbuD4kERZzXPFRORH4NOcr7EAij-dWKTCG0eBBGZObN3Aic0lMAZ_rLU7o6PLOH9Q6p6dyneYjUjSu07vtI5Jy_N2XQpKUni3fUWxfDNgfMo26XKmxdJ5S2Gp1ux1ToO3FY9RByTZw5HZRpOBAfRSgttTgiY5_Yu5D-BLcEm_94R6FMWRniQXrMAt8hU9qCNSuVQlUKtuuy8qJXu6Z21VhI9lAT7wIALlR9UwIgz0e6UG9S9sU95f_38co0ibD1KbQpBd8c_lu5vCVIqlEe_Fa_xYZupMLaU8De-wzoBpBgqR65mRtUQTu2jP-Qxa3aXrANHxweIbnYfpZ5QcNA50hfyVJJSolczDSlDljTunEmHmWNaS3J7waEQsIDFATPFy6H5leRPpSzebXYca4T-EiapPP-mn41Vs3VKIdUXOHus_HcTPWRVT-Vr-yt7byFYEanb5b5lQ_IHcI0oyqX7RrVJid6UsBtwxwkX0FSc1cZgLhBQUgxBsUh5MNte-WZJv_6c9rHyNsH3rn9YEZp431P9GCe8gNdLY9bFQ1pYS9BxOAS2enu3yVpWpWRechiR7D__HC4-Hw2MHfSSmBQTxq5oO01_efEHB8XxWF85XYNT6_icXf3ZsTxkURT9HlHapkFwL7TlO5gPUZZVJt9f6kn9HoGQ56MX2n46KdKKid8=`

  if message.Me != XMLNS_ME {
    t.Errorf("Expected to be %s, got %s", XMLNS_ME, message.Me)
  }

  if message.Encoding != BASE64_URL {
    t.Errorf("Expected to be %s, got %s", BASE64_URL, message.Encoding)
  }

  if message.Alg != RSA_SHA256 {
    t.Errorf("Expected to be %s, got %s", RSA_SHA256, message.Alg)
  }

  if message.Data.Type != APPLICATION_XML {
    t.Errorf("Expected to be %s, got %s", APPLICATION_XML, message.Data.Type)
  }

  if message.Data.Data != data {
    t.Errorf("Expected to be %s, got %s", data, message.Data.Data)
  }

  if message.Sig.Sig != sig {
    t.Errorf("Expected to be %s, got %s", sig, message.Sig.Sig)
  }

  if message.Sig.KeyId != TEST_AUTHOR {
    t.Errorf("Expected to be %s, got %s", TEST_AUTHOR, message.Sig.KeyId)
  }

  if message.Signature() != `NbuD4kERZzXPFRORH4NOcr7EAij-dWKTCG0eBBGZObN3Aic0lMAZ_rLU7o6PLOH9Q6p6dyneYjUjSu07vtI5Jy_N2XQpKUni3fUWxfDNgfMo26XKmxdJ5S2Gp1ux1ToO3FY9RByTZw5HZRpOBAfRSgttTgiY5_Yu5D-BLcEm_94R6FMWRniQXrMAt8hU9qCNSuVQlUKtuuy8qJXu6Z21VhI9lAT7wIALlR9UwIgz0e6UG9S9sU95f_38co0ibD1KbQpBd8c_lu5vCVIqlEe_Fa_xYZupMLaU8De-wzoBpBgqR65mRtUQTu2jP-Qxa3aXrANHxweIbnYfpZ5QcNA50hfyVJJSolczDSlDljTunEmHmWNaS3J7waEQsIDFATPFy6H5leRPpSzebXYca4T-EiapPP-mn41Vs3VKIdUXOHus_HcTPWRVT-Vr-yt7byFYEanb5b5lQ_IHcI0oyqX7RrVJid6UsBtwxwkX0FSc1cZgLhBQUgxBsUh5MNte-WZJv_6c9rHyNsH3rn9YEZp431P9GCe8gNdLY9bFQ1pYS9BxOAS2enu3yVpWpWRechiR7D__HC4-Hw2MHfSSmBQTxq5oO01_efEHB8XxWF85XYNT6_icXf3ZsTxkURT9HlHapkFwL7TlO5gPUZZVJt9f6kn9HoGQ56MX2n46KdKKid8=` {
    t.Errorf("Expected different signature, got '%s'", message.Signature())
  }
}
