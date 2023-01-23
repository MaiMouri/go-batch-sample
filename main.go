package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const IPV4_URL = "https://ipv4.fetus.jp/"

func main() {
	countryCode := "ph"
	res := readFile(countryCode)
	fmt.Println(res)
	// filePath := fmt.Sprintf("%s%s.txt", IPV4_URL, countryCode)
	// fmt.Println(filePath)

	// resp, err := http.Get(filePath)
	// if err != nil {
	// 	fmt.Printf("[ERROR] unread %s", filePath)
	// 	// return nil
	// }
	// fmt.Println(resp)

	// r := csv.NewReader(f)

	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(record)
	// }
}

func readFile(countryCode string) []string {
	filePath := fmt.Sprintf("%s%s.txt", IPV4_URL, countryCode)
	// fmt.Println(filePath)
	resp, err := http.Get(filePath)
	if err != nil {
		fmt.Printf("[ERROR] unread %s", filePath)
		return nil
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	response := strings.Split(string(content), "\n")
	fmt.Println(response)
	var list []string
	for _, line := range response {
		matched, _ := regexp.MatchString("^[0-9]", line)
		if matched {
			list = append(list, line)
		}
	}
	return list
}
