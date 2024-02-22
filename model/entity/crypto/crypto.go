package cryptoEntity

import (
	"4crypto/utils/common"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type (
	Market    string
	OrderType string

	Exchange struct {
		Client     *ethclient.Client
		Users      map[int64]*User
		Orders     map[int64]int64
		PrivateKey *ecdsa.PrivateKey
		Orderbooks map[Market]*common.Orderbook
	}

	PlaceOrderReq struct {
		UserID int64
		Type   OrderType
		Bid    bool
		Size   float64
		Price  float64
		Market Market
	}

	Order struct {
		UserID    int64
		ID        int64
		Price     float64
		Size      float64
		Bid       bool
		Timestamp int64
	}

	OrderbookData struct {
		TotalBidVolume float64
		TotalAskVolume float64
		Asks           []*Order
		Bids           []*Order
	}

	MatchedOrders struct {
		Price float64
		Size  float64
		ID    int64
	}

	CancelOrderReq struct {
		Bid bool
		ID  int64
	}

	User struct {
		ID         int64
		PrivateKey *ecdsa.PrivateKey
	}
)

const (
	MarketOrder OrderType = "MARKET"
	LimitOrder  OrderType = "LIMIT"

	MarketETH = "ETH"

	PrivateKey = "be9b39aca0be732959111784e07ed8270eba44dc078a9918ab984f1af7f75422"
)

func NewUser(pk string, id int64) (*User, error) {
	privateKey, err := crypto.HexToECDSA(pk)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:         id,
		PrivateKey: privateKey,
	}, nil
}

// func NewExchange(PrivateKey string, client *ethclient.Client) (*Exchange, error) {
// 	orderbooks := make(map[Market]*orderbook.Orderbook)
// 	orderbooks[MarketETH] = orderbook.NewOrderBook()

// 	privateKey, err := crypto.HexToECDSA(PrivateKey)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Exchange{
// 		client:     client,
// 		users:      make(map[int64]*User),
// 		orders:     make(map[int64]int64),
// 		PrivateKey: privateKey,
// 		orderbooks: orderbooks,
// 	}, nil
// }
