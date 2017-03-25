// Copyright 2017 Trung Pham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id3

import (
	"bytes"
	"io/ioutil"
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
		t.Errorf("Open: Incorrect artist, %s", s)
	}

	if s := tag.Title(); s != "A Good Song" {
		t.Errorf("Open: Incorrect title, %v", s)
	}

	if s := tag.Album(); s != "Life" {
		t.Errorf("Open: Incorrect album, %v", s)
	}
}

func TestClose(t *testing.T) {
	before, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("Test file error")
	}

	file, err := Open(testFile)
	if err != nil {
		t.Errorf("Close: Unable to open file")
	}
	beforeCutoff := file.originalSize

	file.SetArtist("Kim")
	file.SetTitle("A Test Song")

	afterCutoff := file.Size()

	if err := file.Close(); err != nil {
		t.Errorf("Close: Unable to close file")
	}

	after, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("Close: Unable to reopen file")
	}

	if !bytes.Equal(before[beforeCutoff:], after[afterCutoff:]) {
		t.Errorf("Close: Lose nontag data on close")
	}

	if err := ioutil.WriteFile(testFile, before, 0666); err != nil {
		t.Errorf("Close: Unable to write original contents to test file")
	}
}
