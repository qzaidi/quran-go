package quran

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var langs []string

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./data/qurandb")
	if err != nil {
		log.Fatal("unable to open qurandb", err)
	}

	langs, err = AvailableLangs()
	if err != nil {
		log.Fatal("failed to fetch available languages from db\n")
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

// Get Metadata for a given chapter
func Chapter(chapter int) (*ChapterMeta, error) {

	var c ChapterMeta
	q := "select * from chapters where id = ?"

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println("prepare:", err, q)
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(chapter).Scan(&c.Start, &c.Ayas, &c.Ord, &c.Rukus, &c.Arname, &c.Tname, &c.Enname, &c.Text, &c.Id)
	if err != nil {
		log.Println("scan:", err)
		return nil, err
	}

	return &c, nil

}

func Select(filters Filters, options Options) ([]Verse, error) {
	chapter := filters.Chapter
	verse := filters.Verse

	j := " "
	f := "ar"

	for _, lang := range options.Langs {
		f += "," + lang
		j += "join " + lang + " using (chapter,verse)"
	}

	q := "select " + f + " from ar a" + j + "  where chapter = ? and verse = ? order by chapter,verse"

	log.Println("query:", q)

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Println("prepare:", err, q)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(chapter, verse)
	if err != nil {
		log.Println("query:", err, q)
		return nil, err
	}

	defer rows.Close()

	cols, _ := rows.Columns()

	data := make([]interface{}, len(cols))
	for idx := range cols {
		data[idx] = new(string)
	}

	var verses []Verse

	for rows.Next() {
		verse := make(Verse)
		err = rows.Scan(data...)
		if err != nil {
			log.Println("scan:", err)
			return nil, err
		}
		for idx := range cols {
			str := data[idx].(*string)
			//log.Println(*str)
			verse[cols[idx]] = *str
		}
		log.Println(verse)
		verses = append(verses, verse)
	}
	return verses, nil
}
