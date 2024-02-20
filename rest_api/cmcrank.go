package main

import (
	"4crypto/model/dto/res"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	//exchange asset
	//exchage metadata

	client := resty.New()

	client.SetHeader("X-CMC_PRO_API_KEY", "ba777d1f-caee-4be5-8314-335bb1c9ea35")

	// Ranking Crypto CMC
	urlAPICoinRank := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map?start=1&limit=10&sort=cmc_rank"

	respCoinRank, err := client.R().Get(urlAPICoinRank)
	if err != nil {
		log.Println("resp.err:", err.Error())
		return
	}

	// convert JSON -> struct
	err = json.Unmarshal(respCoinRank.Body(), &res.ResponseCoinRank)
	if err != nil {
		log.Println("Unmarshal.err:", err.Error())
		return
	}

	fmt.Println("Status:", res.ResponseCoinRank.Status)
	fmt.Println("Data:")
	for _, rank := range res.ResponseCoinRank.Data {
		fmt.Println(rank)
	}

	// Cryptocurrency Map
	urlAPICoinMap := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/map"

	respCoinMap, err := client.R().Get(urlAPICoinMap)
	if err != nil {
		log.Println("resp.err:", err.Error())
		return
	}

	// convert JSON -> struct
	err = json.Unmarshal(respCoinMap.Body(), &res.ResponseCoinRank)
	if err != nil {
		log.Println("Unmarshal.err:", err.Error())
		return
	}

	fmt.Println("Status:", res.ResponseCoinRank.Status)
	fmt.Println("Data:")
	for _, rank := range res.ResponseCoinRank.Data {
		fmt.Println(rank)
	}

	// Fiat Map
	urlAPIFiatMap := "https://pro-api.coinmarketcap.com/v1/fiat/map"

	respFiatMap, err := client.R().Get(urlAPIFiatMap)
	if err != nil {
		log.Println("resp.err:", err.Error())
		return
	}

	// convert JSON -> struct
	err = json.Unmarshal(respFiatMap.Body(), &res.ResponseFiat)
	if err != nil {
		log.Println("Unmarshal.err:", err.Error())
		return
	}

	fmt.Println("Status:", res.ResponseFiat.Status)
	fmt.Println("Data:")
	for _, rank := range res.ResponseFiat.Data {
		fmt.Println(rank)
	}

	//Exchange Map
	urlAPIExchangeMap := "https://pro-api.coinmarketcap.com/v1/exchange/map"

	respExchangeMap, err := client.R().Get(urlAPIExchangeMap)
	if err != nil {
		log.Println("resp.err:", err.Error())
		return
	}

	// convert JSON -> struct
	err = json.Unmarshal(respExchangeMap.Body(), &res.ResponseExhange)
	if err != nil {
		log.Println("Unmarshal.err:", err.Error())
		return
	}

	fmt.Println("Status:", res.ResponseExhange.Status)
	fmt.Println("Data:")
	for _, rank := range res.ResponseExhange.Data {
		fmt.Println(rank)
	}
}
