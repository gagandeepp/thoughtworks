package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fxtlabs/date"
	"io/ioutil"
	"net/http"
	"strconv"
)

type response3 struct {
	Name      string `json:"name"`
	Category  string `json:"category"`
	Price     int    `json:"price"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
func main() {



	client := &http.Client{}
	r, _ := http.NewRequest("GET","https://http-hunt.thoughtworks-labs.net/challenge/input",nil )
	r.Header.Add("userId", "ttDT1_hcK")
	res, _ := client.Do(r)
	body, _ := ioutil.ReadAll(res.Body)

	var data []response3
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	//file,_ := ioutil.ReadFile("test.json")
	//json.Unmarshal([]byte(file), &data)

	count := map[string]int{}
	for i := 0; i < len(data); i++ {
		fmt.Println("i=", i, "start", data[i].StartDate, "end ", data[i].EndDate)

		year_start, _ := strconv.Atoi(data[i].StartDate[:4])
		month_start, _ := strconv.Atoi(data[i].StartDate[5:7])
		day_start, _ := strconv.Atoi(data[i].StartDate[8:])

		if data[i].EndDate == "" {
			if year_start < 2018 {
				if _, ok := count[data[i].Category]; ok{
					count[data[i].Category]++
				} else {
					count[data[i].Category] = 1
				}
				continue
			} else if year_start == 2018 {
				if month_start == 12 && day_start >= date.Today().Day() {
					continue
				}
				if _, ok := count[data[i].Category]; ok{
					count[data[i].Category]++
				} else {
					count[data[i].Category] = 1
				}
			}
		}

		if year_start > date.Today().Year() {
			continue
		} else if year_start == date.Today().Year() {
			if month_start == 12 && day_start > date.Today().Day() {
				continue
			}
		}

		year_end, _ := strconv.Atoi(data[i].EndDate[:4])
		month_end, _ := strconv.Atoi(data[i].EndDate[5:7])
		day_end, _ := strconv.Atoi(data[i].EndDate[8:])

		if year_end > date.Today().Year() {
			if _, ok := count[data[i].Category]; ok{
				count[data[i].Category]++
			} else {
				count[data[i].Category] = 1
			}

		} else if year_end == date.Today().Year() {
			if month_end == 12 && day_end >= date.Today().Day() {
				if _, ok := count[data[i].Category]; ok{
					count[data[i].Category]++
				} else {
					count[data[i].Category] = 1
				}
			}
		}

	}

	//for i:=0; i< len(data); i++ {
	//	if _, ok := count[data[i].Category]; ok{
	//		count[data[i].Category]++
	//	} else {
	//		count[data[i].Category] = 1
	//	}
	//}
	//


	//maps := map[string]int{}
	//maps["count"] = count

	answer, _ := json.Marshal(count)


	req, err := http.NewRequest("POST", "https://http-hunt.thoughtworks-labs.net/challenge/output",  bytes.NewReader(answer))
	req.Header.Add("userid", "ttDT1_hcK")
	req.Header.Add("content-type","application/json")
	res, _ = client.Do(req)
	fmt.Println(count)

}

//func generateOutput(w http.ResponseWriter , req *http.Request,  ps httprouter.Params) {
//	w.Header().Set("userId", "ttDT1_hcK")
//	w.Header().Set("Content-Type", "application/json")
//	w.Write()
//}



//{"stage":"1/4","statement":"Start silly, acceleration for the count. Given list of products, return back the count and you get a furious start","instructions":"You can 'GET' the input from /challenge/input and output should be 'POST' json to /challenge/output. Important note: The time between request 'GET' input and 'POST' requests should not exceed 2 secs.","sampleInput":{"input":[{"name":"Apple iPhone 6s 128GB Space Grey Refurbished","price":1737},{"name":"Nokia 3220","price":999},{"name":"iBall Slide Brace-X1 Mini 16GB","price":1737},{"name":"Dekor World POLKA DOT DIWAN SET- PACK OF 8 PCS","price":786},{"name":"Swayam Yellow Colour Floral Printed Eyelet Door Curtain - Window Curtain","price":654}]},"sampleOutput":{"output":{"count":5}}}
