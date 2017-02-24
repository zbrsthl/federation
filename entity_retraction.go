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

type EntityRelayableSignedRetraction struct {
  TargetGuid string `xml:"target_guid"`
  TargetType string `xml:"target_type"`
  SenderHandle string `xml:"sender_handle"`
  TargetAuthorSignature string `xml:"target_author_signature"`
  ParentAuthorSignature string `xml:"parent_author_signature"`
}

type EntityRetraction struct {
  DiasporaHandle string `xml:"diaspora_handle"`
  PostGuid string `xml:"post_guid"`
  Type string `xml:"type"`
}
