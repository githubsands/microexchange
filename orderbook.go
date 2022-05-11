package exchange

import (
	"errors"
	"fmt"
	"math"

	"github.com/davecgh/go-spew/spew"
)

type price struct {
	price int
	side  string

	nextOrder *orderLocation
	prevPrice *price
}

type orderLocation struct {
	prevOrder *orderLocation
	nextOrder *orderLocation

	size  float64
	id    string
	price int
}

type Order struct {
	id     string
	action string
	size   float64
	pair   string
	price  int
	time   float64
}

type Orderbook struct {
	asset    string
	maxPrice int
	size     int

	bids map[int]*price
	asks map[int]*price

	trading bool
}

// NewOrderbook mallocs a new order book on the heap. Its data structure is a map of price's for each side of the spread (BID, ASK),
// Where each price point contains a row of orderLocations.
//
// P1	 o1 o2 o3 o4 o5 o6 o7 o8 o9 ... oRowSize
// P2	 o1 o2 o3 o4 o5 o6 o7 o8 o9 ... oRowSize
// PHigh o1 o2 o3 o4 o5 o6 o7 o8 o9 ... oRowSize
//
func NewOrderbook(asset string, high int, low int, rowSize int) *Orderbook {
	orderBook := new(Orderbook)
	orderBook.size = rowSize
	orderBook.maxPrice = high
	orderBook.asset = asset
	orderBook.trading = true

	// allocate mmeory on the heap for our bid side

	var currentPricePoint *price
	var prevPrice *price
	var currentOrder *orderLocation
	orderBook.bids = make(map[int]*price)
	var prevOrder *orderLocation
	for i := low; i < high+1; i++ {
		currentPricePoint = new(price)
		currentPricePoint.price = i
		currentPricePoint.side = "BID"
		currentPricePoint.prevPrice = prevPrice
		orderBook.bids[i] = currentPricePoint

		for j := 0; j < rowSize+1; j++ {
			if j == 0 {
				currentPricePoint.nextOrder = new(orderLocation)
				currentOrder = currentPricePoint.nextOrder
				currentOrder.price = j
			} else {
				prevOrder = currentOrder
				currentOrder = new(orderLocation)
				currentOrder.price = j
				prevOrder.nextOrder = currentOrder
				currentOrder.prevOrder = prevOrder
			}
		}

		prevPrice = currentPricePoint
	}

	currentPricePoint = nil
	prevPrice = nil
	orderBook.asks = make(map[int]*price)
	for i := low; i < high+1; i++ {
		currentPricePoint = new(price)
		currentPricePoint.price = i
		currentPricePoint.side = "ASK"
		currentPricePoint.prevPrice = prevPrice
		orderBook.asks[i] = currentPricePoint

		for j := 0; j < rowSize+1; j++ {
			if j == 0 {
				currentPricePoint.nextOrder = new(orderLocation)
				currentOrder = currentPricePoint.nextOrder
				currentOrder.price = j
			} else {
				prevOrder = currentOrder
				currentOrder = new(orderLocation)
				currentOrder.price = j
				prevOrder.nextOrder = currentOrder
				currentOrder.prevOrder = prevOrder
			}
		}

		prevPrice = currentPricePoint
	}

	return orderBook
}

func (o *Orderbook) Depths(price int) float64 {
	return o.bids[price].nextOrder.size
}

func (o *Orderbook) Ask(order Order) error {
	if order.action != "ASK" {
		return errors.New("Wrong order type expecting order action ASK")
	}

	if order.price < 1 {
		return errors.New(fmt.Sprintf("Order price must be greater then or equal to 1"))
	}

	if order.pair != o.asset {
		return errors.New(fmt.Sprintf("Wrong asset expecting %v, received %v", o.asset, order.pair))
	}

	if order.price > o.maxPrice {
		return errors.New("Too large of an error")
	}

	if o.bids[order.price].nextOrder == nil {
		return errors.New("Have exceeded the length of orders on this price point")
	}

	if o.trading {
		order, state := o.match(o.bids[order.price], o.bids[order.price].nextOrder, order)
		if state != "filled" {
			o.insertAsk(order.id, order.price, order.size)
		}
	} else {
		o.insertAsk(order.id, order.price, order.size)
	}

	return nil
}

func (o *Orderbook) Bid(order Order) error {
	if order.action != "BID" {
		return errors.New("Wrong order type expecting order action BID")
	}

	if order.pair != o.asset {
		return errors.New(fmt.Sprintf("Wrong asset expecting %v, received %v", o.asset, order.pair))
	}

	if order.price > o.maxPrice {
		return errors.New("Too large of an error")
	}

	if order.price < 1 {
		return errors.New(fmt.Sprintf("Order price must be greater then or equal to 1"))
	}

	if o.asks[order.price].nextOrder == nil {
		return errors.New("Have exceeded the length of orders on this price point")
	}

	if o.trading {
		order, state := o.match(o.asks[order.price], o.asks[order.price].nextOrder, order)
		if state != "filled" {
			o.insertBid(order.id, order.price, order.size)
		}
	} else {
		o.insertBid(order.id, order.price, order.size)
	}

	return nil
}

/*
func (o *Orderbook) Modify(order Order) {
	switch order.action {
	case "ask":
		next := o.asks[order.price].nextOrder
		for {
			if order.id != next.id {
				next = next.nextOrder
			} else {
				next.size = order.size
				next.prevOrder.nextOrder = next.nextOrder // now that the order has been updated move it to the end of the list due to time priority
				next.tail.nextOrder = next
				continue
			}
		}

	case "bids":
		next := o.bids[order.price].nextOrder
		for {
			if order.id != next.id {
				next = next.nextOrder
			} else {
				next.size = order.size
				next.prevOrder.nextOrder = next.nextOrder // now that the order has been updated move it to the end of the list due to time priority
				next.tail.nextOrder = next
				continue
			}
		}
	}
}
*/

const (
	canceledOrderStatusFail    = "order does not exist"
	canceledOrderStatusSuccess = "order filled"
)

func (o *Orderbook) Cancel(order Order) string {
	var status string
	var orderbookOrder *orderLocation
	switch order.action {
	case "CANCEL":
		price := o.asks[order.price]
		orderbookOrder = price.nextOrder
		for {
			if orderbookOrder == nil {
				status = canceledOrderStatusFail
				break
			}
			if orderbookOrder.id == order.id {
				orderbookOrder.id = ""
				orderbookOrder.size = 0
				status = canceledOrderStatusSuccess
				break
			} else {
				orderbookOrder = price.nextOrder
				continue
			}
		}
	}
	return status
}

// match handles the core matching of the orderbook and is divided into X states:
//
// 1. order size is not fulfilled (=0); no order entrys exist on any pricepoints
// 2. order size is not fulfilled; no order entries exist on this pricepoint but exist on the next pricepoint
// 3. order size is not fulfilled; order entries exist on this pricepoint

func (o *Orderbook) match(price *price, insertion *orderLocation, order Order) (Order, string) {
	orderEntry := price.nextOrder
	fmt.Println(orderEntry.nextOrder)
	var size float64 = order.size
	var filled string
MATCHING_ENGINE:
	for order.size > 0 {
		if price.prevPrice == nil {
			break MATCHING_ENGINE
		}

		if order.size > 0 && orderEntry == nil {
			price = price.prevPrice
			orderEntry = price.nextOrder
			continue
		}

		if order.size > 0 && orderEntry.size > 0 {
			subtractor := order.size
			order.size = subtractor - orderEntry.size
			orderEntry.size = orderEntry.size - subtractor
			order.size = 0
			if math.Signbit(orderEntry.size) || orderEntry.size == 0 {
				orderEntry.size = 0
				orderEntry.id = ""
			}
			orderEntry = orderEntry.nextOrder
		}

		orderEntry = orderEntry.nextOrder
	}

	if order.size == size {
		filled = "unfilled"
		return order, filled
	}

	if order.size > 0 {
		filled = "partial-filled"
		return order, filled
	}

	return Order{}, "filled"
}

func (o *Orderbook) insertBid(id string, price int, size float64) {
	if price == 0 {
		return
	}
	fmt.Println(id)
	currentOrder := o.bids[price].nextOrder
	for {
		if currentOrder.size > 0 {
			currentOrder = currentOrder.nextOrder
			continue
		} else {
			currentOrder.id = id
			currentOrder.size = size
			break
		}
	}
}

func (o *Orderbook) insertAsk(id string, price int, size float64) {
	if price == 0 {
		return
	}
	currentOrder := o.asks[price].nextOrder
	for {
		if currentOrder.size > 0 {
			currentOrder = currentOrder.nextOrder
			continue
		} else {
			currentOrder.id = id
			currentOrder.size = size
			break
		}
	}
}

func (o *Orderbook) PrintAsks(price int) {
	fmt.Println(spew.Sdump(o.asks[price]))
}

func (o *Orderbook) PrintBids(price int) {
	fmt.Println(spew.Sdump(o.bids[price]))
}
