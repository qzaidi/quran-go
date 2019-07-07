package quran

import (
	"testing"
)

func TestLoadTrans(t *testing.T) {
	err := LoadTrans("non-existent.txt")
	if err == nil {
		t.Error("download should have failed for non-existent translation")
	}
}
