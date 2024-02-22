package orderbookEntity

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      float64
}

type Limit struct {
	Price       float64
	Orders      Orders
	TotalVolume float64
}

type Limits []*Limit

type ByBestAsk struct{ Limits }

type Orders []*Order

type Orderbook struct {
	asks []*Limit
	bids []*Limit

	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
	Orders    map[int64]*Order
}

type Order struct {
	ID        int64
	UserID    int64
	Size      float64
	Bid       bool
	Limit     *Limit
	Timestamp int64
}
