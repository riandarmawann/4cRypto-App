package main

import (
	"4crypto/model/dto/res"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {

	urlAPI := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map?start=1&limit=10&sort=cmc_rank"
	client := resty.New()

	client.SetHeader("X-CMC_PRO_API_KEY", "ba777d1f-caee-4be5-8314-335bb1c9ea35")

	resp, err := client.R().Get(urlAPI)
	if err != nil {
		log.Println("resp.err:", err.Error())
		return
	}

	// convert JSON -> struct
	err = json.Unmarshal(resp.Body(), &res.Response)
	if err != nil {
		log.Println("Unmarshal.err:", err.Error())
		return
	}

	fmt.Println("Status:", res.Response.Status)
	fmt.Println("Data:")
	for _, rank := range res.Response.Data {
		fmt.Println(rank)
	}
}
