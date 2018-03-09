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

import "errors"

const (
  ACTIVITY_STREAMS = "https://www.w3.org/ns/activitystreams"

  ActivityTypeCreate = "Create"
  ActivityTypeNote = "Note"
  ActivityTypePerson = "Person"
  ActivityTypeCollection = "Collection"
  ActivityTypeOrderedCollection = "OrderedCollection"
  ActivityTypeCollectionPage = "CollectionPage"
  ActivityTypeOrderedCollectionPage = "OrderedCollectionPage"
)

var (
  ERROR_MISSING_TYPE = errors.New("missing type in request")
)
