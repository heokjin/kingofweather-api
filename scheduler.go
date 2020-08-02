package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/parnurzeal/gorequest"
	"heokjin/kingofweather-api/model"
	"time"
)

func GoScheduleDeleteDB() {
	fmt.Println("Go MidLand Schedule!")
	s := gocron.NewScheduler()
	s.Every(3).Day().At("16:00").Do(taskDeleteDB)
	<-s.Start()
}

func GoScheduleMidLandFcst() {
	fmt.Println("Go MidLand Schedule!")
	s := gocron.NewScheduler()
	s.Every(1).Day().At("21:10").Do(taskWeatherMidLandFcst, "0600")
	s.Every(1).Day().At("09:10").Do(taskWeatherMidLandFcst, "1800")
	<-s.Start()
}

func GoScheduleMidTemp() {
	fmt.Println("Go MidTemp Schedule!")
	s := gocron.NewScheduler()
	s.Every(1).Day().At("21:11").Do(taskWeatherMidTemp, "0600")
	s.Every(1).Day().At("09:11").Do(taskWeatherMidTemp, "1800")
	<-s.Start()
}

func taskDeleteDB() {
	model.DeleteWeatherDB()
}

//중기강수(6/18)
func taskWeatherMidLandFcst(date string) {
	//날짜계산
	s := time.Now().Format("20060102")
	tmFc := s + date

	request := gorequest.New()
	url := "http://apis.data.go.kr/1360000/MidFcstInfoService/getMidLandFcst?"
	serviceKey := "a1lu6Dpas2bFHganEisR8Z0sJKcIpU%2BhXPLFzIsyUq4Pk2HTAhNrVfKeRj%2BJSAKw8gayDE0OgQ5nX5a2SCwLIw%3D%3D"
	pageNo := "2"
	numOfRows := "100"
	types := "json"

	//지역별로 loop
	for index, regId := range MidLandList {
		fmt.Printf("Get MidLandFcst index -> %d\n", index)

		fullUrl := fmt.Sprintf("%sServiceKey=%s&tmFc=%s&regId=%s&pageNo=%s&numOfRows=%s&dataType=%s", url, serviceKey, tmFc, regId, pageNo, numOfRows, types)
		_, body, _ := request.Get(fullUrl).Timeout(20 * time.Second).End()

		fmt.Println(fullUrl)
		fmt.Println(body)
		model.InsertMidLandFcst(tmFc, regId, body)
		fmt.Println("==============================================================================")
	}
}

//중기 기온
func taskWeatherMidTemp(date string) {
	//날짜계산
	s := time.Now().Format("20060102")
	tmFc := s + date

	request := gorequest.New()
	url := "http://apis.data.go.kr/1360000/MidFcstInfoService/getMidTa?"
	serviceKey := "a1lu6Dpas2bFHganEisR8Z0sJKcIpU%2BhXPLFzIsyUq4Pk2HTAhNrVfKeRj%2BJSAKw8gayDE0OgQ5nX5a2SCwLIw%3D%3D"
	pageNo := "1"
	numOfRows := "100"
	types := "json"

	//지역별로 loop
	for index, regId := range MidTempList {
		fmt.Printf("Get MidTemp index -> %d\n", index)

		fullUrl := fmt.Sprintf("%sServiceKey=%s&tmFc=%s&regId=%s&pageNo=%s&numOfRows=%s&dataType=%s", url, serviceKey, tmFc, regId, pageNo, numOfRows, types)
		_, body, _ := request.Get(fullUrl).Timeout(20 * time.Second).End()

		fmt.Println(fullUrl)
		fmt.Println(body)
		model.InsertMidTemp(tmFc, regId, body)
		fmt.Println("==============================================================================")
	}
}
