package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/parnurzeal/gorequest"
	"heokjin/kingofweather-api/model"
	"time"
)

func GoScheduleMidLandFcst() {
	fmt.Println("Go Schedule!")

	s := gocron.NewScheduler()
	s.Every(1).Day().At("06:10").Do(taskWeatherMidLandFcst06hour)
	<-s.Start()

}

//중기강수(6/18)
func taskWeatherMidLandFcst06hour() {
	//날짜계산
	s := time.Now().Format("20060102")
	tmFc := s + "0600"

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
