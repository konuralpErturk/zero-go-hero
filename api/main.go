package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	content := callApi()

	convertResponseToJson(content)
}

func convertResponseToJson(content string) {
	panic("unimplemented")
}

func callApi() string {
	url := "https://dummyjson.com/products"
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
	content := string(bytes)
	return content
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
