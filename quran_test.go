package quran

import (
	"testing"
)

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
