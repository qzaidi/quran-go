package main

import (
	"fmt"

	quran "github.com/qzaidi/quran-go"
)

func main() {
	verses, err := quran.Select(quran.Filters{Chapter: 1, Verse: 7}, quran.Options{Langs: []string{"ur"}})
	if err != nil {
		fmt.Println("error", err)
		return
	}
	for _, verse := range verses {
		fmt.Println(verse)
	}
}
