package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var (
	db *sql.DB
)

func InitDB() {
	var err error
	//db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		"ec2-54-146-91-153.compute-1.amazonaws.com", "phzmidhronstsz", "3bf5a66ae8169808307a5a471bb8e4ac66642688d32f7daa51849e2beedafa4e", "d3v41lr940afms")
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		log.Println("Got error when connect database", err)
		log.Fatal(err)
	}

	log.Println("DB OK2:", db)

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

func GetMidLand(tmFc, regId string) (string, error) {
	var response string
	err := db.QueryRow("select response from weather_mid_water where tm_fc=$1 and reg_id=$2 order by id desc limit 1", tmFc, regId).Scan(&response)
	checkError(err)
	return response, err
}

func InsertMidLandFcst(tmFc, regId, response string) {
	sql_statement := "INSERT INTO weather_mid_water (tm_fc, reg_id, response) VALUES ($1, $2, $3);"
	_, err := db.Exec(sql_statement, tmFc, regId, response)
	checkError(err)
}

func InsertMidTemp(tmFc, regId, response string) {
	sql_statement := "INSERT INTO weather_mid_temp (tm_fc, reg_id, response) VALUES ($1, $2, $3);"
	_, err := db.Exec(sql_statement, tmFc, regId, response)
	checkError(err)
}

func InsertShortTemp(tmFc, xy, response string) {
	sql_statement := "INSERT INTO weather_short_temp (tm_fc, reg_id, response) VALUES ($1, $2, $3);"
	_, err := db.Exec(sql_statement, tmFc, xy, response)
	checkError(err)
}

func DeleteWeatherDB() {
	t := time.Now()
	date := t.AddDate(0, 0, -2).Format("2006-01-02 15:04:05")

	sql_statement := "DELETE FROM weather_mid_water WHERE create_dt < $1;"
	_, err := db.Exec(sql_statement, date)
	checkError(err)

	sql_statement = "DELETE FROM weather_mid_temp WHERE create_dt < $1;"
	_, err = db.Exec(sql_statement, date)
	checkError(err)

	sql_statement = "DELETE FROM weather_short_temp WHERE create_dt < $1;"
	_, err = db.Exec(sql_statement, date)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
