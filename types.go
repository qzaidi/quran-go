package quran

import (
	"errors"
)

// TODO: Implement validation for invalid chapter and verse numbers
type Filters struct {
	Chapter int // chapter selector
	Verse   int // verse selector
}

type Options struct {
	Langs []string
}

type ChapterMeta struct {
	Start  int
	Ayas   int
	Rukus  int
	Ord    int
	Id     int
	Arname string // chapter name in arabic
	Tname  string // translated name
	Enname string // transliterated name in english
	Text   string
}

type Verse map[string]string

var (
	ErrNotFound = errors.New("Not Found")
)
