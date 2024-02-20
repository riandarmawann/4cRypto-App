// package main

// import "4crypto/delivery"

//	func main() {
//		delivery.NewServer().Run()
//	}
package main

import (
	"4crypto/model/entity"
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

	// Kita masukkan response JSON kedalam struct TODO
	var cmcRank entity.CmcRank
	//convert JSON -> struct = Unmurshal
	err = json.Unmarshal(resp.Body(), &cmcRank)

	if err != nil {
		log.Println("Unmarshal.err:", err.Error())
		return
	}

	// // Kita masukkan status response JSON kedalam struct CmcRankStatus
	// var cmcRankStatus entity.CmcRankStatus
	// // convert JSON -> struct = Unmarshal
	// err = json.Unmarshal(resp.Body(), &cmcRankStatus)
	// if err != nil {
	// 	log.Println("Unmarshal.err:", err.Error())
	// 	return
	// }

	// fmt.Println("status: ")
	// fmt.Println("CmcRankStatus:", cmcRankStatus)
	fmt.Println("Data: ")
	fmt.Println(string (resp.Body()))
	// fmt.Println("ID: ", cmcRank.ID)
	// fmt.Println("Rank: ", cmcRank.Rank)
	// fmt.Println("Name: ", cmcRank.Name)
	// fmt.Println("Symbol: ", cmcRank.Symbol)
	// fmt.Println("Slug: ", cmcRank.Slug)
	// fmt.Println("Is_Active: ", cmcRank.IsActive)
	// fmt.Println("First_Historical_Data: ", cmcRank.FirstHistoricalData)
	// fmt.Println("Last_Historical_Data: ", cmcRank.LastHistoricalData)
	// fmt.Println("Platform: ", cmcRank.Platform)
}
