package orderbook

import (
	"reflect"
	"testing"
)

// ids walks the list from head and returns order IDs in queue order.
func ids(l *priceLevel) []uint64 {
	var out []uint64
	for o := l.head; o != nil; o = o.next {
		out = append(out, o.ID)
	}
	return out
}

func TestPushBackKeepsFIFO(t *testing.T) {
	l := &priceLevel{price: 100}
	l.pushBack(&Order{ID: 1, Qty: 5})
	l.pushBack(&Order{ID: 2, Qty: 3})
	l.pushBack(&Order{ID: 3, Qty: 2})

	if got, want := ids(l), []uint64{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("order: got %v, want %v", got, want)
	}
	if l.totalQty != 10 {
		t.Fatalf("totalQty: got %d, want 10", l.totalQty)
	}
}

func TestRemove(t *testing.T) {
	l := &priceLevel{price: 100}
	o1 := &Order{ID: 1, Qty: 5}
	o2 := &Order{ID: 2, Qty: 3}
	o3 := &Order{ID: 3, Qty: 2}
	l.pushBack(o1)
	l.pushBack(o2)
	l.pushBack(o3)

	l.remove(o2) // middle
	if got, want := ids(l), []uint64{1, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("after removing middle: got %v, want %v", got, want)
	}

	l.remove(o1) // head
	l.remove(o3) // tail, list becomes empty
	if !l.empty() || l.tail != nil || l.totalQty != 0 {
		t.Fatalf("expected empty level, got head=%v tail=%v qty=%d", l.head, l.tail, l.totalQty)
	}
}
