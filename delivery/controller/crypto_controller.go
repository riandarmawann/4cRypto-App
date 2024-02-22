package controller

import (
	cryptoEntity "4crypto/model/entity/crypto"
	"4crypto/usecase"
	"4crypto/utils/common"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	commonEth "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type CryptoController struct {
	cu usecase.CryptoUseCase
	rg *gin.RouterGroup
	// jwtService common.JwtToken
	cryptoEntity.Exchange
}

func NewCryptoController(cu usecase.CryptoUseCase, rg *gin.RouterGroup, client *ethclient.Client) *CryptoController {
	orderbooks := make(map[cryptoEntity.Market]*common.Orderbook)
	orderbooks[cryptoEntity.MarketETH] = common.NewOrderBook()

	return &CryptoController{cu: cu, rg: rg, Exchange: cryptoEntity.Exchange{
		Client:     client,
		Users:      make(map[int64]*cryptoEntity.User),
		Orders:     make(map[int64]int64),
		Orderbooks: orderbooks,
	}}
}

func (c *CryptoController) Route() {
	cryptoGroup := c.rg.Group("crypto")
	cryptoGroup.GET("/book/:market", c.handleGetBook)
	cryptoGroup.POST("/order", c.handlerPlaceOrder)
	cryptoGroup.DELETE("/order/:id", c.cancelOrder)

	// ex, err := NewExchange(PrivateKey, client)

	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func (c CryptoController) handleGetBook(ctx *gin.Context) {
	market := cryptoEntity.Market(ctx.Param("market"))
	ob, ok := c.Orderbooks[market]

	pk1 := "1b893acda70e2d37856e8a87b7454ab4a3d61b2ef0789cbf5531b25d4ce53836"
	pk2 := "bd1d974c9f35470bf5cef4925393e9f3098bfd16f22592877110093b8cf8ebb3"

	user1, err := cryptoEntity.NewUser(pk1, 1)

	if err != nil {
		log.Fatal(err)
	}

	user2, err := cryptoEntity.NewUser(pk2, 2)

	if err != nil {
		log.Fatal(err)
	}

	c.Users[user1.ID] = user1
	c.Users[user2.ID] = user2

	// address := "0x7a5f533A4ada3369F543Dd205B1a73bD990cb78e"
	// balance, _ := c.Client.BalanceAt(context.Background(), commonEth.HexToAddress(address), nil)

	// fmt.Println(balance)

	address := "0xE7Ba61fb64117f2999cC15fB03B53eeb4c866cd4"
	balance, _ := c.Client.BalanceAt(context.Background(), commonEth.HexToAddress(address), nil)
	fmt.Println("buyer :", balance)

	address = "0x67FeEB101eFaF6eA46a5f16FB0D53AFd9741Bb2F"
	balance, _ = c.Client.BalanceAt(context.Background(), commonEth.HexToAddress(address), nil)
	fmt.Println("seller :", balance)

	if !ok {
		ctx.JSON(http.StatusBadRequest, map[string]any{"msg": "market not found"})
		return
	}

	orderbookData := c.cu.Orderbooks(ob)

	ctx.JSON(http.StatusOK, orderbookData)
	// return
}

func (ex *CryptoController) handlerPlaceOrder(c *gin.Context) {

	var placeOrderData cryptoEntity.PlaceOrderReq

	if err := c.BindJSON(&placeOrderData); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"msg": "bad request",
		})
		return
	}

	market := cryptoEntity.Market(placeOrderData.Market)

	// ob := ex.orderbooks[market]
	order := common.NewOrder(placeOrderData.Bid, placeOrderData.Size, placeOrderData.UserID)

	if placeOrderData.Type == cryptoEntity.LimitOrder {
		// ob.PlaceLimitOrder(placeOrderData.Price, order)
		if err := ex.cu.HandlePlaceLimitOrder(ex.Orderbooks[market], placeOrderData.Price, order); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{
				"msg": "failed to order placed",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"msg": "limit order placed",
		})
		return
	}

	if placeOrderData.Type == cryptoEntity.MarketOrder {

		matches, matchedOrders := ex.cu.HandlePlaceMarketOrder(ex.Orderbooks[market], order)

		if err := ex.cu.HandleMatches(matches, ex.Users, ex.Client); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{
				"msg": "failed to match order",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]any{
			"matches": matchedOrders,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{})

}

func (ex *CryptoController) cancelOrder(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{"msg": "Gagal Parsing Id"})
		return
	}

	ob := ex.Orderbooks[cryptoEntity.MarketETH]
	order := ob.Orders[id]

	ob.CancelOrder(order)

	c.JSON(http.StatusOK, map[string]any{"msg": "order canceled"})
	// return
}
