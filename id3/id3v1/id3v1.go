// Copyright 2017 Trung Pham. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package id3v1

import (
	"io"
	"os"
)

const (
	TagSize = 128
)

var (
	Genres = []string{
		"Classic Rock",
		"Country",
		"Dance",
		"Disco",
		"Funk",
		"Grunge",
		"Hip-Hop",
		"Jazz",
		"Metal",
		"New Age",
		"Oldies",
		"Other",
		"Pop",
		"R&B",
		"Rap",
		"Reggae",
		"Rock",
		"Techno",
		"Industrial",
		"Alternative",
		"Ska",
		"Death Metal",
		"Pranks",
		"Soundtrack",
		"Euro-Techno",
		"Ambient",
		"Trip-Hop",
		"Vocal",
		"Jazz+Funk",
		"Fusion",
		"Trance",
		"Classical",
		"Instrumental",
		"Acid",
		"House",
		"Game",
		"Sound Clip",
		"Gospel",
		"Noise",
		"Alternative Rock",
		"Bass",
		"Punk",
		"Space",
		"Meditative",
		"Instrumental Pop",
		"Instrumental Rock",
		"Ethnic",
		"Gothic",
		"Darkwave",
		"Techno-Industrial",
		"Electronic",
		"Pop-Folk",
		"Eurodance",
		"Dream",
		"Southern Rock",
		"Comedy",
		"Cult",
		"Gangsta",
		"Top 40",
		"Christian Rap",
		"Pop/Funk",
		"Jungle",
		"Native US",
		"Cabaret",
		"New Wave",
		"Psychadelic",
		"Rave",
		"Showtunes",
		"Trailer",
		"Lo-Fi",
		"Tribal",
		"Acid Punk",
		"Acid Jazz",
		"Polka",
		"Retro",
		"Musical",
		"Rock & Roll",
		"Hard Rock",
		"Folk",
		"Folk-Rock",
		"National Folk",
		"Swing",
		"Fast Fusion",
		"Bebob",
		"Latin",
		"Revival",
		"Celtic",
		"Bluegrass",
		"Avantgarde",
		"Gothic Rock",
		"Progressive Rock",
		"Psychedelic Rock",
		"Symphonic Rock",
		"Slow Rock",
		"Big Band",
		"Chorus",
		"Easy Listening",
		"Acoustic	",
		"Humour",
		"Speech",
		"Chanson",
		"Opera",
		"Chamber Music",
		"Sonata",
		"Symphony",
		"Booty Bass",
		"Primus",
		"Porn Groove",
		"Satire",
		"Slow Jam",
		"Club",
		"Tango",
		"Samba",
		"Folklore",
		"Ballad",
		"Power Ballad",
		"Rhytmic Soul",
		"Freestyle",
		"Duet",
		"Punk Rock",
		"Drum Solo",
		"Acapella",
		"Euro-House",
		"Dance Hall",
		"Goa",
		"Drum & Bass",
		"Club-House",
		"Hardcore",
		"Terror",
		"Indie",
		"BritPop",
		"Negerpunk",
		"Polsk Punk",
		"Beat",
		"Christian Gangsta",
		"Heavy Metal",
		"Black Metal",
		"Crossover",
		"Contemporary C",
		"Christian Rock",
		"Merengue",
		"Salsa",
		"Thrash Metal",
		"Anime",
		"JPop",
		"SynthPop",
	}
)

// Tag represents an ID3v1 tag
type Tag struct {
	title, artist, album, year, comment string
	genre                               byte
	dirty                               bool
}

func ParseTag(readSeeker io.ReadSeeker) *Tag {
	readSeeker.Seek(-TagSize, os.SEEK_END)

	data := make([]byte, TagSize)
	n, err := io.ReadFull(readSeeker, data)
	if n < TagSize || err != nil || string(data[:3]) != "TAG" {
		return nil
	}

	return &Tag{
		title:   string(data[3:33]),
		artist:  string(data[33:63]),
		album:   string(data[63:93]),
		year:    string(data[93:97]),
		comment: string(data[97:127]),
		genre:   data[127],
		dirty:   false,
	}
}

func (t Tag) Dirty() bool {
	return t.dirty
}

func (t Tag) Title() string  { return t.title }
func (t Tag) Artist() string { return t.artist }
func (t Tag) Album() string  { return t.album }
func (t Tag) Year() string   { return t.year }

func (t Tag) Genre() string {
	if int(t.genre) < len(Genres) {
		return Genres[t.genre]
	}

	return ""
}

func (t Tag) Comment() string { return t.comment }

func (t *Tag) SetTitle(text string) {
	t.title = text
	t.dirty = true
}

func (t *Tag) SetArtist(text string) {
	t.artist = text
	t.dirty = true
}

func (t *Tag) SetAlbum(text string) {
	t.album = text
	t.dirty = true
}

func (t *Tag) SetYear(text string) {
	t.year = text
	t.dirty = true
}

func (t *Tag) SetGenre(text string) {
	t.genre = 255
	for i, genre := range Genres {
		if text == genre {
			t.genre = byte(i)
			break
		}
	}
	t.dirty = true
}

func (t Tag) Bytes() []byte {
	data := make([]byte, TagSize)

	copy(data[:3], []byte("TAG"))
	copy(data[3:33], []byte(t.title))
	copy(data[33:63], []byte(t.artist))
	copy(data[63:93], []byte(t.album))
	copy(data[93:97], []byte(t.year))
	copy(data[97:127], []byte(t.comment))
	data[127] = t.genre

	return data
}

func (t Tag) Size() int {
	return TagSize
}

func (t Tag) Version() string {
	return "1.0"
}