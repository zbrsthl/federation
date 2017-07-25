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

import "testing"

func TestCommentSignatureOrder(t *testing.T) {
  var comment EntityComment

  expected := "author created_at guid parent_guid text"
  if expected != comment.SignatureOrder() {
    t.Errorf("Expected to be %s, got %s", expected, comment.SignatureOrder())
  }
}

func TestCommentAppendSignature(t *testing.T) {
  comment := EntityComment{
    Author: "author@localhost",
    CreatedAt: "01.01.1970",
    Guid: "1234",
    ParentGuid: "4321",
    Text: "hello world",
  }

  if comment.AuthorSignature != "" {
    t.Errorf("Expected to be empty, got %s", comment.AuthorSignature)
  }

  if comment.ParentAuthorSignature != "" {
    t.Errorf("Expected to be empty, got %s", comment.AuthorSignature)
  }

  err := comment.AppendSignature(TEST_PRIV_KEY,
    comment.SignatureOrder(), AuthorSignatureType)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if comment.AuthorSignature == "" {
    t.Errorf("Expected signature, was empty")
  }

  err = comment.AppendSignature(TEST_PRIV_KEY,
    comment.SignatureOrder(), ParentAuthorSignatureType)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if comment.ParentAuthorSignature == "" {
    t.Errorf("Expected signature, was empty")
  }
}
