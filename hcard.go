package federation
//
// GangGo Federation Library
// Copyright (C) 2017-2018 Lukas Matt <lukas@zauberstuhl.de>
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
  "strings"
  "github.com/PuerkitoBio/goquery"
)

type Hcard struct {
  Guid string
  Nickname string
  FullName string
  Searchable bool
  PublicKey string
  FirstName string
  LastName string
  Url string
  Photo string
  PhotoMedium string
  PhotoSmall string
}

func (h *Hcard) Fetch(endpoint string) error {
  resp, err := FetchHtml("GET", endpoint, nil)
  if err != nil {
    return err
  }

  doc, err := goquery.NewDocumentFromResponse(resp)
  if err != nil {
    return err
  }

  doc.Find(".entity_uid").Each(
  func(i int, s *goquery.Selection) {
    (*h).Guid = s.Find("span").Text()
  })
  doc.Find(".entity_nickname").Each(
  func(i int, s *goquery.Selection) {
    (*h).Nickname = s.Find("span").Text()
  })
  doc.Find(".entity_full_name").Each(
  func(i int, s *goquery.Selection) {
    (*h).FullName = s.Find("span").Text()
  })
  doc.Find(".entity_first_name").Each(
  func(i int, s *goquery.Selection) {
    (*h).FirstName = s.Find("span").Text()
  })
  doc.Find(".entity_family_name").Each(
  func(i int, s *goquery.Selection) {
    (*h).LastName = s.Find("span").Text()
  })
  doc.Find(".entity_key").Each(
  func(i int, s *goquery.Selection) {
    (*h).PublicKey = s.Find("pre").Text()
  })
  doc.Find(".entity_url").Each(
  func(i int, s *goquery.Selection) {
    (*h).Url = s.Find("a").Text()
  })

  doc.Find(".entity_searchable").Each(
  func(i int, s *goquery.Selection) {
    var searchable bool = false
    if s.Find("span").Text() == "true" {
      searchable = true
    }
    (*h).Searchable = searchable
  })

  doc.Find(".entity_photo").Each(
  func(i int, s *goquery.Selection) {
    nodes := s.Find("img")
    for _, node := range nodes.Nodes {
      for _, attr := range node.Attr {
        if attr.Key == "src" {
          value := attr.Val
          if !strings.HasPrefix(value, "http") {
            value = h.Url + value
          }
          (*h).Photo = value
        }
      }
    }
  })
  doc.Find(".entity_photo_medium").Each(
  func(i int, s *goquery.Selection) {
    nodes := s.Find("img")
    for _, node := range nodes.Nodes {
      for _, attr := range node.Attr {
        if attr.Key == "src" {
          value := attr.Val
          if !strings.HasPrefix(value, "http") {
            value = h.Url + value
          }
          (*h).PhotoMedium = value
        }
      }
    }
  })
  doc.Find(".entity_photo_small").Each(
  func(i int, s *goquery.Selection) {
    nodes := s.Find("img")
    for _, node := range nodes.Nodes {
      for _, attr := range node.Attr {
        if attr.Key == "src" {
          value := attr.Val
          if !strings.HasPrefix(value, "http") {
            value = h.Url + value
          }
          (*h).PhotoSmall = value
        }
      }
    }
  })
  return nil
}
