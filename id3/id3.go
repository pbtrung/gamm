// Copyright 2017 Trung Pham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id3

import (
	"os"

	"github.com/pbtrung/gamm/id3/id3v1"
)

// Tagger represents the metadata of a tag
type Tagger interface {
	Title() string
	Artist() string
	Album() string
	Year() string
	Genre() string
	Comments() []string
	SetTitle(string)
	SetArtist(string)
	SetAlbum(string)
	SetYear(string)
	SetGenre(string)
	Bytes() []byte
	Dirty() bool
	Padding() uint
	Size() int
	Version() string
}

// File represents the tagged file
type File struct {
	originalSize int
	file         *os.File
}

// Opens a new tagged file
func Open(name string) (*File, error) {
	fi, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	file := &File{file: fi}
	if id3v1Tag := id3v1.ParseTag(fi); id3v1Tag != nil {
		file.Tagger = id3v1Tag
	} else {
		// Add a new tag if none exists
		file.Tagger = v2.NewTag(LatestVersion)
	}

	return file, nil
}
