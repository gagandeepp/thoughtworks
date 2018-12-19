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

type response4 struct {
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

	var data []response4
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	//file,_ := ioutil.ReadFile("test.json")
	//json.Unmarshal([]byte(file), &data)

	count := 0
	for i := 0; i < len(data); i++ {
		fmt.Println("i=", i, "start", data[i].StartDate, "end ", data[i].EndDate, "price ", data[i].Price)

		year_start, _ := strconv.Atoi(data[i].StartDate[:4])
		month_start, _ := strconv.Atoi(data[i].StartDate[5:7])
		day_start, _ := strconv.Atoi(data[i].StartDate[8:])

		if data[i].EndDate == "" {
			if year_start < 2018 {
				count += data[i].Price
				continue
			} else if year_start == 2018 {
				if month_start == 12 && day_start >= date.Today().Day() {
					continue
				}
				count += data[i].Price
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
			count += data[i].Price
		} else if year_end == date.Today().Year() {
			if month_end == 12 && day_end >= date.Today().Day() {
				count += data[i].Price
			}
		}

	}
	maps := map[string]int{}
	maps["totalValue"] = count
	answer, _ := json.Marshal(maps)
	req, err := http.NewRequest("POST", "https://http-hunt.thoughtworks-labs.net/challenge/output",  bytes.NewReader(answer))
	req.Header.Add("userid", "ttDT1_hcK")
	req.Header.Add("content-type","application/json")
	res, _ = client.Do(req)
	fmt.Println(count)

}
