package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB(name string, delete bool) {
	db, _ = sql.Open("sqlite3", "./" + name + ".db")

	if delete {
		statement, _ := db.Prepare("DROP TABLE IF EXISTS playerscore")
		statement.Exec()
	}

	var err error
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS playerscore (name varchar(200) PRIMARY KEY, score INT)")
	if err != nil {
		log.Fatalf("CREATE TABLE playerscore Prepare: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("CREATE TABLE playerscore Exec: %v", err)
	}
}

type PlayerScore struct {
	Name string
	Score int
}

func SetScore(name string, score int) {
	statement, err := db.Prepare("INSERT INTO playerscore (name, score) VALUES (?, ?) ON CONFLICT(name) DO UPDATE SET score = ?")
	if err != nil {
		log.Fatalf("SetScore db.Prepare: %v", err)
	}
	_, err = statement.Exec(name, score, score)
	if err != nil {
		log.Fatalf("SetScore Exec: %v", err)
	}

	if err != nil {
		log.Fatalf("InsertThing statement.Exec: %v", err)
	}
}

func GetScores(start int,  count int) (ts []PlayerScore) {
 	rows, err := db.Query("SELECT name, score FROM playerscore ORDER BY score DESC, name ASC LIMIT ?, ?", start, count)
 	defer rows.Close()
	if err != nil {
		log.Fatalf("GetScores db.Query: %v", err)
	}
 	for rows.Next() {
 		var t PlayerScore
 		err = rows.Scan(&t.Name, &t.Score)
 		if err != nil {
 			log.Fatalf("GetScores rows.Scan: %v", err)
 		}
 		ts = append(ts, t)
 	}

 	return ts
}
