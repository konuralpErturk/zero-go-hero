package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.example1.com/"
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	req.Header.Set("User-Agent", "")
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	fmt.Printf("response type %T\n", resp)

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)
	fmt.Println(string(bytes))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
