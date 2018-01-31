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

import "github.com/Zauberstuhl/go-xml"

type EntityStatusMessage struct {
  XMLName xml.Name `xml:"status_message"`
  Author string `xml:"author"`
  Guid string `xml:"guid"`
  CreatedAt Time `xml:"created_at"`
  ProviderName string `xml:"provider_display_name"`
  Text string `xml:"text,omitempty"`
  Photos *EntityPhotos `xml:"photo,omitempty"`
  Location *EntityLocation `xml:"location,omitempty"`
  Poll *EntityPoll `xml:"poll,omitempty"`
  Public bool `xml:"public"`
  Event *EntityEvent `xml:"event,omitempty"`
}

type EntityReshare struct {
  XMLName xml.Name `xml:"reshare"`
  Author string `xml:"author"`
  Guid string `xml:"guid"`
  CreatedAt Time `xml:"created_at"`
  RootAuthor string `xml:"root_author"`
  RootGuid string `xml:"root_guid"`
}

type EntityLocation struct {
  Address string `xml:"address"`
  Lat string `xml:"lat"`
  Lng string `xml:"lng"`
}

type EntityPoll struct {
  Guid string `xml:"guid"`
  Question string `xml:"question"`
  PollAnswers []EntityPollAnswer `xml:"poll_answers"`
}

type EntityPollAnswer struct {
  Guid string `xml:"guid"`
  Answer string `xml:"answer"`
}

type EntityEvent struct {
  Author string `xml:"author"`
  Guid string `xml:"guid"`
  Summary string `xml:"summary"`
  Start Time `xml:"start"`
  End Time `xml:"end"`
  AllDay bool `xml:"all_day"`
  Timezone string `xml:"timezone"`
  Description string `xml:"description"`
  Location *EntityLocation `xml:"location,omitempty"`
}
