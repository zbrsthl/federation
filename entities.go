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

import "encoding/xml"

/* Header

  <?xml version="1.0" encoding="UTF-8"?>
  <diaspora xmlns="https://joindiaspora.com/protocol" xmlns:me="http://salmon-protocol.org/ns/magic-env">
    <header>
      <author_id>dia@192.168.0.173:3000</author_id>
    </header>
*/
type Header struct {
  XMLName xml.Name `xml:"header"`
  AuthorId string `xml:"author_id"`
}

/* Encrypted Header

  <?xml version="1.0" encoding="UTF-8"?>
  <decrypted_header>
    <iv>...</iv>
    <aes_key>...</aes_key>
    <author_id>one@two.tld</author_id>
  </decrypted_header>
*/
// XXX legacy?
type XmlDecryptedHeader struct {
  XMLName xml.Name `xml:"decrypted_header"`
  Iv string `xml:"iv"`
  AesKey string `xml:"aes_key"`
  AuthorId string `xml:"author_id"`
}

type JsonEnvHeader struct {
  AesKey string `json:"aes_key"`
  Ciphertext string `json:"ciphertext"`
}


// Private Request
//
//  <?xml version="1.0" encoding="UTF-8"?>
//  <diaspora xmlns="https://joindiaspora.com/protocol" xmlns:me="http://salmon-protocol.org/ns/magic-env">
//    <encrypted_header>eyJhZXNfa2V5IjoiVTFGbVZ5TE5CT0pyOWViZVpiMUxnSGQzMlJwMzNMa1ViZVRnY3BEaDluZHAyT2cyMUZNeHh2Mm9RUXp0eWxLVjZsOEkzR0wvQlEzcEQxbGNYbjFGQWlEVTBmTHJwRUZwUEJNc1AwT3MvdmhEZ3I3MDJvaWNkMUxFN0ZMZ3ZVc1VKRGxGOHdvTVZUaGtnTmNVWGlCNUZpajcvb1J4ZFZ1QlJxbVpCQ0VXLys4PSIsImNpcGhlcnRleHQiOiJveWhWVjR5bXJoTGxYRVU1WVcxQWdpK29sTkxRSDlvR2NPQnVBOG00a25FZ09GZnJoNzEwUi9XQloyQ1I3WVlMeEhuOEcvUWFyU1UraDFka0VHMGFhcS9FQ1RWQjFHOXNkT2lQMnZSRTRkY1JOclNqb3RINXBsV3F4QVBvNlBpRkdwNUJZcURoYnB2RUNYNFNpM3BTN2hyOTdyWUs3NTFyYnYvRjNwRERJNG13d2NoUFA3S21nN21yVXBCZTBEeXRqQk5xK0k3S2lxSVQrRmVTUUMxc0pQSmd4TVAxck1VcDB6cURpT2ZlR3U2RGVnVG5MRnd5WHM3WEhGQW1FN29iSlZJY2NxS3czNjNWcnBrWXE0N1BDRWc3R3p1bFVlOXA5eFhDN3dVRmNOak91RTVQV1NXU1h3cWM0K2EybHFuOCJ9</encrypted_header>
//    <me:env>
//      <me:data type="application/xml">WXRPc09JakFleHh0dW5rQytRallnSEg0a08zOGd3ZVQyVVg4NkppVVZzVlJlMFFHbExhSjBtbzBnRUVGSnQyc2pqRG5OOUFETk1nakJ6MThGVnJsRWNEcjBqUDJ2emZPRjlCVHVLNXlvRjZqQ3NIRlh5USt4L3RlMnUxN0xCSUlxSDk4OGNhL1BVLzE2NEhBL0plQU1sMEZCOFBnTTFzYW9qUUJaV3V1RW95Y25ONWVZTm9wRi9adElXU3AxOFVzL05RWG13cktiWVVYYXJpNitOSWo1RWVOdTJrYnJnN2FoWUdhRFJKMVFtcEdhL0pMNmVvMEM1ZHFkS2s1amgrNFU0dnA0ZVY5bkRPK0ZGSkhpd3pDWm84NWJkdUR4T0NRRURpa2QxeGE1OWl4NkZWMTYzQ0MvMXBqdmhTa3cvTjNOTGdHOUZjWExQZDRwdTYzQ2pkS1E4QlE2MTYrMEFGREV0bjF0RldqYU9DV0VXOW5ORzQrTmNVZ0ZKamM0dzZxU3Ruc00rSFdzS0hOTnpRZElRc1pXT3c2bzJUK3ZCOUZmeUY3T3NTNHhwYnF0OW5MZFM1Z0U1M1lFb3FXVmRzK3E1a2xYV3dXdjlBeHRGdlhIVWxyNEE9PQ==</me:data>
//      <me:encoding>base64url</me:encoding>
//      <me:alg>RSA-SHA256</me:alg>
//      <me:sig>sJPN--TJ9IqVwhT_j9mNGrLF4yQvwXUQKo24cPLi5FVXl-tVpyEOxrUI1gwRJ5j5UkkqNJO8mLph2ravlxqt7PNhS9YAOTuo46nXWXyOJjP_ESxq3DaMrYqQt57PDnM29x5yQ0QATbSAs6XneHtxmVKzwKgi4ZpdQ3THFj_iWwac0BI3Or1okt9wLxxl3LTLO9vwfIZaeo-XDNT7JlIfSMZDv1xipjtbl-P0z0q4u2wYOLquvvDjRdI_9vStZK3EOmYARhDXhH0vcJNjVXYTuq16BtXsyfEW3WLBPH67t9Ef3c6cWqU3qPSS3-ddZY5VVq6pPpmtnHuBNzB5hZvZ8asMexc7S0V075ZG-7axUcwXkWKTwCZuxwZNm3VinQze4meWY6vWITtD6zHCguMIWZgxW5z7LGZ04j8_26NbBmZXV52-TFRJExi6H6kUmDb3GrYTlLTOziEUB9pl4NkCX_ghi-Pixbzg7zc_LD_cKXoUyj7iHRNfdfXLck9SOxXkmWAI7hNAswBgx1ngnH6AVyNSsFYVNF1jzc-uTwANfMQqjqq5_XCdE2Z2GFOvFJQ6JK4S-gAEpLygzeltbvYxuK_qqONA9cCTqoJlOxSgQY_lj7h6s__1EuyP9_-Fh4J9MWi9i118ndkmzaswOcxU_VfnHLDgbQbXD5B7zaS1c90=</me:sig>
//    </me:env>
//  </diaspora>

type PrivateEnvMarshal struct {
  XMLName xml.Name `xml:"me:env"`
  Me string `xml:"xmlns:me,attr"`
  Data struct {
    XMLName xml.Name `xml:"me:data"`
    Type string `xml:"type,attr"`
    Data string `xml:",chardata"`
  }
  Encoding string `xml:"me:encoding"`
  Alg string `xml:"me:alg"`
  Sig struct {
    XMLName xml.Name `xml:"me:sig"`
    Sig string `xml:",chardata"`
    KeyId string `xml:"key_id,attr,omitempty"`
  }
}

type PrivateMarshal struct {
  XMLName xml.Name `xml:"diaspora"`
  Xmlns string `xml:"xmlns,attr"`
  XmlnsMe string `xml:"me,attr"`
  EncryptedHeader string `xml:"encrypted_header,omitempty"`
  Env PrivateEnvMarshal
}

type PublicMarshal struct {
  XMLName xml.Name `xml:"diaspora"`
  Xmlns string `xml:"xmlns,attr"`
  XmlnsMe string `xml:"me,attr"`
  Header Header
  Env PrivateEnvMarshal
}

// NOTE I had huge problems with marshal
// and unmarshal with the same structs
// apperently namespaces in tags are a problem
// for the go xml implementation

type DiasporaUnmarshal struct {
  XMLName xml.Name `xml:"diaspora"`
  Xmlns string `xml:"xmlns,attr"`
  XmlnsMe string `xml:"me,attr"`
  Header struct {
    XMLName xml.Name `xml:"header"`
    AuthorId string `xml:"author_id"`
  }
  EncryptedHeader string `xml:"encrypted_header,omitempty"`
  DecryptedHeader *XmlDecryptedHeader `xml:",omitempty"`
  Env struct {
    XMLName xml.Name `xml:"env"`
    Me string `xml:"me,attr"`
    Data struct {
      XMLName xml.Name `xml:"data"`
      Type string `xml:"type,attr"`
      Data string `xml:",chardata"`
    }
    Encoding string `xml:"encoding"`
    Alg string `xml:"alg"`
    Sig struct {
      XMLName xml.Name `xml:"sig"`
      Sig string `xml:",chardata"`
      KeyId string `xml:"key_id,attr,omitempty"`
    }
  }
}

type Entity struct {
  XMLName xml.Name `xml:"XML"`
  Post EntityPost `xml:"post"`
}

type EntityPost struct {
  XMLName xml.Name `xml:"post,omitempty"`
  Request *EntityRequest `xml:"request,omitempty"`
  Retraction *EntityRetraction `xml:"retraction,omitempty"`
  Profile *EntityProfile `xml:"profile,omitempty"`
  Reshare *EntityStatusMessage `xml:"reshare,omitempty"`
  StatusMessage *EntityStatusMessage `xml:"status_message,omitempty"`
  Comment *EntityComment `xml:"comment,omitempty"`
  Like *EntityLike `xml:"like,omitempty"`
  SignedRetraction *EntityRelayableSignedRetraction `xml:"signed_retraction,omitempty"`
  RelayableRetraction *EntityRelayableSignedRetraction `xml:"relayable_retraction,omitempty"`
}
