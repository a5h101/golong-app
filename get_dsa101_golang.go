package get_dsa101_golang

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get_dsa101_golang() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://dsa101.onrender.com/get-chapters", nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println(string(body))
}
