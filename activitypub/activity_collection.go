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

type ActivityCollection struct {
  *ActivityContext
  TotalItems int `json:"totalItems"`
  First *ActivityCollectionPage `json:"first,omitempty"`
  Items *[]ActivityCreate `json:"items,omitempty"`
  OrderedItems *[]ActivityCreate `json:"orderedItems,omitempty"`
}

type ActivityCollectionPage struct {
  *ActivityBase
  TotalItems int `json:"totalItems"`
  PartOf string `json:"partOf"`
  Items *[]ActivityCreate `json:"items,omitempty"`
  OrderedItems *[]ActivityCreate `json:"orderedItems,omitempty"`
}

type ActivityCreate struct {
  *ActivityBase
  Actor string `json:"actor"`
  Published Time `json:"published"`
  To []string `json:"to"`
  Cc []string `json:"cc"`
  Object ActivityNote `json:"object"`
}

type ActivityNote struct {
  *ActivityBase
  Summary string `json:"summary"`
  Content string `json:"content"`
  InReplyTo string `json:"inReplyTo"`
  Published Time `json:"published"`
  Url string `json:"url"`
  To []string `json:"to"`
  Cc []string `json:"cc"`
  Sensitive bool `json:"sensitive"`
  Attachments []string `json:"attachment"`
}
