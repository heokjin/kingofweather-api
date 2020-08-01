package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	db *sql.DB
)

func InitDB() {
	var err  error
	//db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		"ec2-54-146-91-153.compute-1.amazonaws.com", "phzmidhronstsz", "3bf5a66ae8169808307a5a471bb8e4ac66642688d32f7daa51849e2beedafa4e", "d3v41lr940afms")
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Println("Got error when connect database", err)
		log.Fatal(err)
	}

	log.Println("DB OK1:", db)

	var rows *sql.Rows
	rows, err = db.Query("SELECT cluster_id, user_id FROM users")
	if err != nil {
		panic(err)
	}
	var cluster_id int8
	var user_id string

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cluster_id, &user_id)
		if err != nil {
			panic(err)
		}

		fmt.Println(cluster_id, user_id)
	}


}