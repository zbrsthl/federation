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
  "crypto"
  "crypto/rsa"
  "crypto/x509"
  "crypto/rand"
  "crypto/sha256"
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
  "encoding/pem"
  "encoding/xml"
  "encoding/json"
  "errors"
  "github.com/revel/revel"
)

func ParseBase64RSAPubKey(encodedKey string) (pubkey *rsa.PublicKey, err error) {
  decodedKey, err := base64.StdEncoding.DecodeString(encodedKey)
  if err != nil {
    return
  }
  return ParseRSAPubKey(decodedKey)
}

func ParseRSAPubKey(decodedKey []byte) (pubkey *rsa.PublicKey, err error) {
  block, _ := pem.Decode(decodedKey)
  if block == nil {
    err = errors.New("Decode public key block is nil!")
    return
  }
  data, err := x509.ParsePKIXPublicKey(block.Bytes)
  if err != nil {
    return
  }
  switch data := data.(type) {
  case *rsa.PublicKey:
    pubkey = data
  default:
    err = errors.New("Wasn't able to parse the public key!")
  }
  return
}

func ParseRSAPrivKey(decodedKey []byte) (privkey *rsa.PrivateKey, err error) {
  block, _ := pem.Decode(decodedKey)
  if block == nil {
    err = errors.New("Decode private key block is nil!")
    return
  }
  privkey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
  if err != nil {
    return
  }
  return
}

func (aes *XmlDecryptedHeader) DecryptAES(ciphertext *[]byte, data string) error {
  key, err := base64.StdEncoding.DecodeString(aes.AesKey)
  if err != nil {
    return err
  }

  iv, err := base64.StdEncoding.DecodeString(aes.Iv)
  if err != nil {
    return err
  }

  *ciphertext, err = base64.URLEncoding.DecodeString(data)
  if err != nil {
    return err
  }
  // diaspora magic do it twice
  *ciphertext, err = base64.StdEncoding.DecodeString(string(*ciphertext))
  if err != nil {
    return err
  }

  *ciphertext, err = DecryptAES(key, iv, *ciphertext)
  if err != nil {
    return err
  }
  return nil
}

func (aes *JsonAesKey) DecryptAES(ciphertext *[]byte, data string) error {
  key, err := base64.StdEncoding.DecodeString(aes.Key)
  if err != nil {
    return err
  }

  iv, err := base64.StdEncoding.DecodeString(aes.Iv)
  if err != nil {
    return err
  }

  *ciphertext, err = base64.StdEncoding.DecodeString(data)
  if err != nil {
    return err
  }

  *ciphertext, err = DecryptAES(key, iv, *ciphertext)
  if err != nil {
    return err
  }
  return nil
}

func DecryptAES(key, iv, ciphertext []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return ciphertext, err
  }

  mode := cipher.NewCBCDecrypter(block, iv)
  mode.CryptBlocks(ciphertext, ciphertext)

  return ciphertext, nil
}

func (request *DiasporaUnmarshal) VerifySignature(serialized []byte) error {
  pubkey, err := ParseRSAPubKey(serialized)
  if err != nil {
    warn(err)
    return err
  }

  type64 := base64.StdEncoding.EncodeToString(
    []byte(request.Env.Data.Type),
  )
  encoding64 := base64.StdEncoding.EncodeToString(
    []byte(request.Env.Encoding),
  )
  alg64 := base64.StdEncoding.EncodeToString(
    []byte(request.Env.Alg),
  )

  text := request.Env.Data.Data + "." + type64 + "." + encoding64 + "." + alg64

  h := sha256.New()
  h.Write([]byte(text))
  digest := h.Sum(nil)

  ds, err := base64.URLEncoding.DecodeString(request.Env.Sig.Sig)
  if err != nil {
    return err
  }

  return rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, digest, ds)
}

func AuthorSignature(data interface{}, privKey string) (string, error) {
  var text string
  switch entity := data.(type) {
  case *EntityComment:
    text = entity.DiasporaHandle+";"+entity.Guid+";"+
      entity.ParentGuid+";"+entity.Text
  case *EntityLike:
    positive := "false"
    if entity.Positive {
      positive = "true"
    }
    // positive guid parent_guid parent_type author
    text = positive+";"+entity.Guid+";"+entity.ParentGuid+
      ";"+entity.TargetType+";"+entity.DiasporaHandle
  case *EntityStatusMessage:
    // parent author signature
    //timeLayout := "2006-01-02T15:04:05-07:00"
    //timeLayout := "2006-01-02T15:04:05.000000000Z"
    //       e.g. 2017-02-20T00:44:37.548811913+01:00
    public := "false"
    if entity.Public {
      public = "true"
    }





// <XML><post><status_message><diaspora_handle>lukas@192.168.0.173:9000</diaspora_handle><guid>de7f4b0a17ad43b74627628ee5f956fc</guid><created_at>2017-02-20T14:41:44.545004637Z</created_at><provider_display_name>GangGo</provider_display_name><raw_message>app/models/entities/post</raw_message><public>true</public></status_message></post></XML>


//<comment><diaspora_handle>lukas@192.168.0.173:9000</diaspora_handle><guid>3cadba7a21004e2ddebb0c3d18cb1951</guid><parent_guid>de7f4b0a17ad43b74627628ee5f956fc</parent_guid><text>app/models/entities/post2</text><author_signature>nTOIXwLuxeyR3Df9mEmIt56CnEFrnEWTCBfilioKydE6oHcjGLZTZL2nONzgBWmkhBCrM7TMd2k6LJFWDBZE4gTDirbBneAEamWtXoNAsQRjUD4NNJpfVjqZXC3D9269nmZQ1eRojIjNfNWHB13LDBIWQC1yfKG0Hokg23745nlqMeKGWB3ntrNm5rOHPfpcRt/VoxqB80nUXYdSePsbagAPh5KxvIaf+rNQdNVa5r8d1bHXyk41Doh2a4JdyVnC1D+vPJkX5R9vtoVbbzpSRSFQ9zJnbkVqIjVtS7oi3zI3zU+liM/n0iXA424/HOZW+nmP6LRGLsj3y/Wn6HW32A==</author_signature><parent_author_signature>jxxB0YtDkxfy+aAuKXliefGdtfcEAV1BmisdzYTNiay+ua7jtuSHq3AAIECEAQVwMWpxZ36uB69d6ji4KBqdJFMaD8momPqd7WjZYrTKUSObFl8mq2REAL1pBBFmECgmjyiOixwBNFE6r2dqD6Uk85GsY+IhNCuhvLsM7sM5ZYVZdoqPZv3lC2j6m/fvStG3EtYh7ruu/XxR1+mucRU5M4UaS48MD86F55+pDS7qj7QDMk/bronBz9SkDPJkpYGerrRTP4njKkNMhWjPxy8wpYURW8DRFqZLjOcYLLnmhppej3wJXrrx4PwufVW/zU6eZVdll8jsiaK7cmFWfjGKsQ==</parent_author_signature></comment>

//lukas@192.168.0.173:9000;de7f4b0a17ad43b74627628ee5f956fc;2017-02-20T15:41:44+00:00;GangGo;app/models/entities/post;true

    text = entity.DiasporaHandle+";"+entity.Guid+";"+
      //entity.CreatedAt.UTC().Format(timeLayout)+";"+
      "2017-02-20T14:41:44.545004637Z;"+
      entity.ProviderName+";"+entity.RawMessage+";"+public

      // author guid parent_guid text

//lukas@192.168.0.173:9000;de7f4b0a17ad43b74627628ee5f956fc;2017-02-20T14:41:44.545004637Z;GangGo;app/models/entities/post;true



    revel.WARN.Println(text)

  }
  return Sign(text, privKey)
}

func (d *DiasporaMarshal) Sign(privKey string) (err error) {
  type64 := base64.StdEncoding.EncodeToString(
    []byte(d.Data.Type),
  )
  encoding64 := base64.StdEncoding.EncodeToString(
    []byte(d.Encoding),
  )
  alg64 := base64.StdEncoding.EncodeToString(
    []byte(d.Alg),
  )

  text := d.Data.Data + "." + type64 +
    "." + encoding64 + "." + alg64
  (*d).Sig.Sig, err = Sign(text, privKey)
  return
}


func Sign(text, privKey string) (sig string, err error) {
  privkey, err := ParseRSAPrivKey([]byte(privKey))
  if err != nil {
    return "", err
  }

  h := sha256.New()
  h.Write([]byte(text))
  digest := h.Sum(nil)

  rng := rand.Reader
  sigInByte, err := rsa.SignPKCS1v15(rng, privkey, crypto.SHA256, digest[:])
  if err != nil {
    return "", err
  }
  return base64.StdEncoding.EncodeToString(sigInByte), nil
}


func (request *DiasporaUnmarshal) DecryptHeader(serialized []byte) error {
  header := JsonEnvHeader{}
  decryptedHeader := XmlDecryptedHeader{}

  decoded, err := base64.StdEncoding.DecodeString(request.EncryptedHeader)
  if err != nil {
    return err
  }

  err = json.Unmarshal(decoded, &header)
  if err != nil {
    return err
  }

  privkey, err := ParseRSAPrivKey(serialized)
  if err != nil {
    return err
  }

  decoded, err = base64.StdEncoding.DecodeString(header.AesKey)
  if err != nil {
    return err
  }

  aesKeyJson, err := rsa.DecryptPKCS1v15(rand.Reader, privkey, decoded)
  if err != nil {
    return err
  }

  var aesKey JsonAesKey
  err = json.Unmarshal(aesKeyJson, &aesKey)
  if err != nil {
    return err
  }

  var ciphertext []byte
  err = aesKey.DecryptAES(&ciphertext, header.Ciphertext)
  if err != nil {
    return err
  }

  err = xml.Unmarshal(ciphertext, &decryptedHeader)
  if err != nil {
    return err
  }

  (*request).DecryptedHeader = &decryptedHeader
  return nil
}
