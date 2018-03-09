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

const (
  // entity names
  Retraction = "retraction"
  Profile = "profile"
  StatusMessage = "status_message"
  Reshare = "reshare"
  Comment = "comment"
  Like = "like"
  Contact = "contact"

  // webfinger
  WebFingerOstatus = "http://ostatus.org/schema/1.0/subscribe"
  WebFingerHcard = "http://microformats.org/profile/hcard"

  // signatures
  SignatureDelimiter = "."
  SignatureAuthorDelimiter = ";"
  SignatureHTTPDelimiter = "\n"
)
