package picsart

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UploadData struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type Response struct {
	Name  string `json:"name"`
	Mana  string `json:"mana_cost"`
	Type  string `json:"type_line"`
	Image string `json:"image_uris.normal"`
}

func ApiUpload() {
	fmt.Println("Calling PicsArt /Upload API...")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.scryfall.com/cards/random", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf(" %s\n", responseObject.Name)
	fmt.Printf(" %s\n", responseObject.Type)
	fmt.Printf(" %s\n", responseObject.Mana)
	fmt.Printf(" %s\n", responseObject.Image)
}
