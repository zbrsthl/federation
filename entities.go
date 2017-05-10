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
  RequestBody []byte `xml:"-"` // original body payload
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
