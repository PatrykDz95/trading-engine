package orderbook

type Side uint8

const (
	Buy Side = iota
	Sell
)

func (s Side) opposite() Side {
	if s == Buy {
		return Sell
	}
	return Buy
}

type OrderType uint8

const (
	Limit OrderType = iota
	Market
)

type Order struct {
	ID    uint64
	Side  Side
	Type  OrderType
	Price int64 // in ticks
	Qty   int64 // remaining quantity
}

type Trade struct {
	MakerID uint64
	TakerID uint64
	Price   int64
	Qty     int64
}
