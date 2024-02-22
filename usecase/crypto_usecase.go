package usecase

import (
	cryptoEntity "4crypto/model/entity/crypto"
	"4crypto/repository"
	"4crypto/utils/common"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CryptoUseCase interface {
	// FindById(id string) (entity.User, error)
	// FindByUsernamePassword(username string, password string) (entity.User, error)
	// DeleteById(id string) error
	Orderbooks(ob *common.Orderbook) cryptoEntity.OrderbookData
	HandlePlaceLimitOrder(ob *common.Orderbook, price float64, o *common.Order) error
	HandlePlaceMarketOrder(ob *common.Orderbook, order *common.Order) ([]common.Match, []*cryptoEntity.MatchedOrders)
	HandleMatches(matches []common.Match, users map[int64]*cryptoEntity.User, client *ethclient.Client) error
}

type cryptoUseCase struct {
	cryptoRepo repository.CryptoRepository
}

func NewCryptoUseCase(cryptoRepo repository.CryptoRepository) CryptoUseCase {
	return &cryptoUseCase{cryptoRepo: cryptoRepo}
}

func (cu cryptoUseCase) Orderbooks(ob *common.Orderbook) cryptoEntity.OrderbookData {

	orderbookData := cryptoEntity.OrderbookData{
		TotalBidVolume: ob.BidTotalVolume(),
		TotalAskVolume: ob.AskTotalVolume(),
		Asks:           []*cryptoEntity.Order{},
		Bids:           []*cryptoEntity.Order{},
	}

	for _, limit := range ob.Asks() {
		for _, order := range limit.Orders {
			o := cryptoEntity.Order{
				UserID:    order.UserID,
				ID:        order.ID,
				Price:     limit.Price,
				Size:      order.Size,
				Bid:       order.Bid,
				Timestamp: order.Timestamp,
			}
			orderbookData.Asks = append(orderbookData.Asks, &o)
		}
	}

	for _, limit := range ob.Bids() {
		for _, order := range limit.Orders {
			o := cryptoEntity.Order{
				UserID:    order.UserID,
				ID:        order.ID,
				Price:     limit.Price,
				Size:      order.Size,
				Bid:       order.Bid,
				Timestamp: order.Timestamp,
			}
			orderbookData.Bids = append(orderbookData.Bids, &o)
		}
	}

	return orderbookData
}

func (cu *cryptoUseCase) HandlePlaceLimitOrder(ob *common.Orderbook, price float64, o *common.Order) error {

	ob.PlaceLimitOrder(price, o)

	return nil
}

func (ex *cryptoUseCase) HandlePlaceMarketOrder(ob *common.Orderbook, order *common.Order) ([]common.Match, []*cryptoEntity.MatchedOrders) {
	// ob := ex.orderbooks[market]
	matches := ob.PlaceMarketOrder(order)
	matchedOrders := make([]*cryptoEntity.MatchedOrders, len(matches))

	for i := 0; i < len(matchedOrders); i++ {
		var id int64

		if !order.Bid {
			id = matches[i].Bid.ID
		} else {
			id = matches[i].Ask.ID
		}

		matchedOrders[i] = &cryptoEntity.MatchedOrders{
			ID:    id,
			Size:  matches[i].SizeFilled,
			Price: matches[i].Price,
		}
	}

	return matches, matchedOrders

}

func (ex *cryptoUseCase) HandleMatches(matches []common.Match, users map[int64]*cryptoEntity.User, client *ethclient.Client) error {

	for _, match := range matches {
		fromUser, ok := users[match.Ask.UserID]

		if !ok {
			return errors.New("user not found")
		}

		// fromAddress := crypto.PubkeyToAddress(fromUser.PrivateKey.PublicKey)

		toUser, ok := users[match.Bid.UserID]

		if !ok {
			return errors.New("user not found")
		}

		toAddress := crypto.PubkeyToAddress(toUser.PrivateKey.PublicKey)

		ammount := big.NewInt(int64(match.SizeFilled))

		common.TransferETH(client, fromUser.PrivateKey, toAddress, ammount)

	}

	return nil
}
