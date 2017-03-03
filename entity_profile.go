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

type EntityProfile struct {
  DiasporaHandle string `xml:"diaspora_handle"`
  FirstName string `xml:"first_name"`
  LastName string `xml:"last_name"`
  ImageUrl string `xml:"image_url"`
  ImageUrlMedium string `xml:"image_url_medium"`
  ImageUrlSmall string `xml:"image_url_small"`
  Birthday string `xml:"birthday"`
  Gender string `xml:"gender"`
  Bio string `xml:"bio"`
  Location string `xml:"location"`
  Searchable bool `xml:"searchable"`
  Nsfw bool `xml:"nsfw"`
  TagString string `xml:"tag_string"`
}