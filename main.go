package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func WritetoFile(f *os.File, str string) {
	_, err := f.WriteString(str + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		arr := strings.Split(line, "-")
		company_name := arr[0]
		address := arr[1]
		//search web for company
		client := &http.Client{
			Timeout: time.Second * 10,
		}

		request_get, err := http.NewRequest("GET", "https://www.google.com/search?", nil)

		if err != nil {
			log.Fatal(err)
		}

		query := request_get.URL.Query()
		query.Add("query", address)
		request_get.URL.RawQuery = query.Encode()

		response, _ := client.Do(request_get)
		response_val, _ := io.ReadAll(response.Body)

		string_val := string(response_val)
		//write to file
		f, _ := os.OpenFile(("out.html"), os.O_CREATE|os.O_WRONLY, 0644)

		_, err2 := f.WriteString(string_val)

		if err2 != nil {
			log.Fatal(err2)
		}

		//f.Close()
		company_name = strings.ReplaceAll(company_name, " ", "")
		company_name = strings.ToLower(company_name)
		readFile2, _ := os.Open("out.html")
		fileScanner2 := bufio.NewScanner(readFile2)
		fileScanner2.Split(bufio.ScanLines)
		f2, _ := os.OpenFile(("output.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		for fileScanner2.Scan() {
			line2 := fileScanner2.Text()
			fmt.Println(line2)
			if strings.Contains(line2, company_name) {
				WritetoFile(f2, line2)
			}
		}
		f.Close()

	}
}

/*
client := &http.Client{
		Timeout: time.Second * 10,
	}

	request_get, err := http.NewRequest("GET", "https://www.google.com/search?", nil)

	if err != nil {
		log.Fatal(err)
	}

	query := request_get.URL.Query()
	query.Add("query", "7735 W 59th St, Summit, IL 60501")
	request_get.URL.RawQuery = query.Encode()

	response, _ := client.Do(request_get)
	response_val, _ := io.ReadAll(response.Body)

	string_val := string(response_val)
	fmt.Println(string_val)
	//fmt.Println(string(response_val))

	f, _ := os.OpenFile(("out.html"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	_, err2 := f.WriteString(string_val)

	if err2 != nil {
		log.Fatal(err2)
	}

	f.Close()

	// Bork Transport, LLC
	company_name := "Bork Transport"
	company_name = strings.ReplaceAll(company_name, " ", "")
	company_name = strings.ToLower(company_name)
	//fmt.Println(company_name)
	readFile, _ := os.Open("out.html")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	f, _ = os.OpenFile(("output.txt"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, company_name) {
			WritetoFile(f, line)
		}
	}
	f.Close()

*/
