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
  "time"
  "encoding/xml"
)

type EntityStatusMessage struct {
  XMLName xml.Name `xml:"status_message"`
  Author string `xml:"author"`
  Guid string `xml:"guid"`
  CreatedAt time.Time `xml:"created_at"`
  ProviderName string `xml:"provider_display_name"`
  Text string `xml:"text,omitempty"`
  Photo *EntityPhotos `xml:"photo,omitempty"`
  Location *EntityLocation `xml:"location,omitempty"`
  Poll *EntityPoll `xml:"poll,omitempty"`
  Public bool `xml:"public"`
  // on reshare
  RootHandle string `xml:"root_diaspora_id,omitempty"`
  RootGuid string `xml:"root_guid,omitempty"`
}

type EntityPhoto struct {
  Guid string `xml:"guid"`
  Author string `xml:"author"`
  Public bool `xml:"public"`
  CreatedAt time.Time `xml:"created_at"`
  RemotePhotoPath string `xml:"remote_photo_path"`
  RemotePhotoName string `xml:"remote_photo_name"`
  Text string `xml:"text"`
  StatusMessageGuid string `xml:"status_message_guid"`
  Height int `xml:"height"`
  Width int `xml:"width"`
}

type EntityPhotos []EntityPhoto

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

type PollParticipation struct {
  PollAnswerGuid string `xml:"poll_answer_guid"`
}
