package orderbook

// priceLevel holds all resting orders at one price(FIFO).
type priceLevel struct {
	price    int64
	totalQty int64
	head     *Order // oldest, matched first
	tail     *Order // newest
}

func (l *priceLevel) pushBack(o *Order) {
	o.level = l
	o.prev = l.tail
	o.next = nil
	if l.tail != nil {
		l.tail.next = o // dotychczasowy ostatni wskazuje na nowy
	} else {
		l.head = o // lista była pusta: nowy jest też pierwszy
	}
	l.tail = o
	l.totalQty += o.Qty
}
func (l *priceLevel) remove(o *Order) {
	if o.prev != nil {
		o.prev.next = o.next
	} else {
		l.head = o.next // usuwamy pierwszy
	}
	if o.next != nil {
		o.next.prev = o.prev
	} else {
		l.tail = o.prev // usuwamy ostatni
	}
	l.totalQty -= o.Qty
	o.prev, o.next, o.level = nil, nil, nil
}

func (l *priceLevel) empty() bool { return l.head == nil }
