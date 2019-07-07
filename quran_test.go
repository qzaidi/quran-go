package quran

import (
	"testing"
)

func TestChapter(t *testing.T) {
	c, err := Chapter(1)
	if err != nil {
		t.Error(err)
	}

	t.Log(c)
}

func TestChapterInvalid(t *testing.T) {
	c, err := Chapter(115)
	if err == nil {
		t.Error("expected error, got none")
	}

	t.Log(c, err)
}

func TestAvailableLangs(t *testing.T) {
	langs, err := AvailableLangs()
	if err != nil {
		t.Error(err)
	}

	t.Log(langs)
}

func TestGetVerse(t *testing.T) {
	ar, err := GetVerse(1, 3)
	if err != nil {
		t.Error(err)
	}
	t.Log(ar)
}

func TestSelect(t *testing.T) {
	verses, err := Select(Filters{Chapter: 3, Verse: 55}, Options{Langs: []string{"hi", "ur", "id"}})
	if err != nil {
		t.Error(err)
	}
	t.Log(verses)
}
