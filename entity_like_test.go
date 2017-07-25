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


func TestLikeSignatureOrder(t *testing.T) {
  var like EntityLike

  expected := "positive guid parent_guid target_type author"
  if expected != like.SignatureOrder() {
    t.Errorf("Expected to be %s, got %s", expected, like.SignatureOrder())
  }
}

func TestLikeAppendSignature(t *testing.T) {
  like := EntityLike{
    Positive: true,
    Guid: "1234",
    ParentGuid: "4321",
    TargetType: "Post",
    Author: "author@localhost",
  }

  if like.AuthorSignature != "" {
    t.Errorf("Expected to be empty, got %s", like.AuthorSignature)
  }

  if like.ParentAuthorSignature != "" {
    t.Errorf("Expected to be empty, got %s", like.AuthorSignature)
  }

  err := like.AppendSignature(TEST_PRIV_KEY,
    like.SignatureOrder(), AuthorSignatureType)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if like.AuthorSignature == "" {
    t.Errorf("Expected signature, was empty")
  }

  err = like.AppendSignature(TEST_PRIV_KEY,
    like.SignatureOrder(), ParentAuthorSignatureType)
  if err != nil {
    t.Errorf("Some error occured while parsing: %v", err)
  }

  if like.ParentAuthorSignature == "" {
    t.Errorf("Expected signature, was empty")
  }
}
