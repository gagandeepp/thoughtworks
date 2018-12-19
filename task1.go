package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response1 struct {
	Name string
	Price   int

}

func main() {

	client := &http.Client{}
	r, _ := http.NewRequest("GET","https://http-hunt.thoughtworks-labs.net/challenge/input",nil )
	r.Header.Add("userId", "ttDT1_hcK")
	res, _ := client.Do(r)
	body, _ := ioutil.ReadAll(res.Body)

	var data []response1
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(len(data))
	maps := map[string]int{}
	maps["count"] = len(data)
	answer, _ := json.Marshal(maps)

	req, err := http.NewRequest("POST", "https://http-hunt.thoughtworks-labs.net/challenge/output",  bytes.NewReader(answer))
	req.Header.Add("userid", "ttDT1_hcK")
	req.Header.Add("content-type","application/json")
	res, _ = client.Do(req)
	fmt.Println(res)

}