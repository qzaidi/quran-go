package quran

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./data/qurandb")
	if err != nil {
		log.Fatal("unable to open qurandb", err)
	}
}

func AvailableLangs() ([]string, error) {

	var result []string

	rows, err := db.Query("select name from sqlite_master where type='table'")
	if err != nil {
		return result, err
	}
	defer rows.Close()

	var tableName string

	for rows.Next() {
		rows.Scan(&tableName)
		if err != nil {
			log.Println(err)
			return result, err
		}

		log.Println(tableName)
		if tableName != "chapters" && tableName != "juz" {
			result = append(result, tableName)
		}
	}

	return result, nil
}

func GetVerse(chapter, verse int) (string, error) {

	var arabic string
	stmt, err := db.Prepare("select * from ar where chapter = ? and verse = ?")
	if err != nil {
		log.Println(err)
		return arabic, err
	}

	defer stmt.Close()
	err = stmt.QueryRow(chapter, verse).Scan(&chapter, &verse, &arabic)
	if err != nil {
		return arabic, err
	}

	log.Println(chapter, verse, arabic)
	return arabic, nil
}
