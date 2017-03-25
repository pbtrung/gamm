// Copyright 2017 Trung Pham. All rights reserved.		t.Errorf("Open: Incorrect artist, %b", []byte("Nathan"))

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id3

import (
	"testing"

	"github.com/pbtrung/gamm/id3/id3v1"
)

const (
	testFile = "test.mp3"
)

func TestOpen(t *testing.T) {
	file, err := Open(testFile)
	if err != nil {
		t.Errorf("Open: Unable to open file")
	}

	tag, ok := file.Tagger.(*id3v1.Tag)
	if !ok {
		t.Errorf("Open: Incorrect tagger type")
	}

	if s := tag.Artist(); s != "Nathan" {
		t.Errorf("Open: Incorrect artist, %v", s)
	}

	if s := tag.Title(); s != "A Good Song" {
		t.Errorf("Open: Incorrect title, %v", s)
	}

	if s := tag.Album(); s != "Life" {
		t.Errorf("Open: Incorrect album, %v", s)
	}
}
