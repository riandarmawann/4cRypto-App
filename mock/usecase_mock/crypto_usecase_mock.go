package usecasemock

import (
	"4crypto/model/entity"
	cryptoEntity "4crypto/model/entity/crypto"
	"4crypto/utils/common"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
)

type CryptoUseCaseMock struct {
	mock.Mock
}

func (cum *CryptoUseCaseMock) Orderbooks(ob *common.Orderbook) cryptoEntity.OrderbookData {
	args := cum.Called(ob)
	return args.Get(0).(cryptoEntity.OrderbookData)
}

func (cum *CryptoUseCaseMock) HandlePlaceLimitOrder(ob *common.Orderbook, price float64, o *common.Order) error {
	args := cum.Called(ob, price, o)
	return args.Error(0)
}

func (cum *CryptoUseCaseMock) HandlePlaceMarketOrder(ob *common.Orderbook, order *common.Order) ([]common.Match, []*cryptoEntity.MatchedOrders) {
	args := cum.Called(ob, order)
	return args.Get(0).([]common.Match), args.Get(1).([]*cryptoEntity.MatchedOrders)
}

func (cum *CryptoUseCaseMock) HandleMatches(matches []common.Match, users map[int64]*cryptoEntity.User, client *ethclient.Client) error {
	args := cum.Called(matches, users)
	return args.Error(0)
}

func (cum *CryptoUseCaseMock) HandleCryptoRank() ([]entity.CmcRank, error) {
	args := cum.Called()
	return args.Get(0).([]entity.CmcRank), args.Error(1)
}
