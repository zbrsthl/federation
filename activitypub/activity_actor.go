package activitypub
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

type ActivityActor struct {
  *ActivityContext
  Inbox string `json:"inbox"`
  Outbox string `json:"outbox"`
  Following string `json:"following"`
  Followers string `json:"followers"`

  PreferredUsername *string `json:"preferredUsername,omitempty"`
  Name *string `json:"name,omitempty"`
  Summary *string `json:"summary,omitempty"`
  PublicKey *struct {
    PublicKeyPem string `json:"publicKeyPem"`
  } `json:"publicKey,omitempty"`
  Icon *struct {
    Url string `json:"url"`
  } `json:"icon,omitempty"`
}
