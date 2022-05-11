package exchange

import (
	"testing"
)

func TestOrderbookTraverseBidPricePointWhenEmptyAsksHighToLow(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 2)

	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13436", action: "BID", size: 15, pair: "BTC/ETH", price: 3}
	err := orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	if orderbook.bids[3].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.bids[3].nextOrder.size)
	}
	order = Order{id: "13435", action: "BID", size: 15, pair: "BTC/ETH", price: 2}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	if orderbook.bids[3].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.bids[2].nextOrder.size)
	}
	order = Order{id: "13434", action: "BID", size: 15, pair: "BTC/ETH", price: 1}
	err = orderbook.Bid(order)
	if orderbook.bids[1].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.bids[1].nextOrder.size)
	}
}

func TestOrderbookTraverseBidPricePointWhenEmptyAsksLowToHigh(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 2)

	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13436", action: "BID", size: 15, pair: "BTC/ETH", price: 1}
	err := orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	if orderbook.bids[1].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.bids[1].nextOrder.size)
	}
	order = Order{id: "13435", action: "BID", size: 15, pair: "BTC/ETH", price: 2}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	if orderbook.bids[2].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.bids[2].nextOrder.size)
	}
	order = Order{id: "13434", action: "BID", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Bid(order)
	if orderbook.bids[3].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.bids[3].nextOrder.size)
	}
}

func TestOrderbookTraverseAsksPricePointWhenEmptyBidsHighToLow(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 2)

	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13436", action: "ASK", size: 15, pair: "BTC/ETH", price: 3}
	err := orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	if orderbook.asks[3].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.asks[3].nextOrder.size)
	}
	order = Order{id: "13435", action: "ASK", size: 15, pair: "BTC/ETH", price: 2}
	err = orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	if orderbook.asks[3].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.asks[2].nextOrder.size)
	}
	order = Order{id: "13434", action: "ASK", size: 15, pair: "BTC/ETH", price: 1}
	err = orderbook.Ask(order)
	if orderbook.asks[1].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.asks[1].nextOrder.size)
	}
}

func TestOrderbookTraverseAsksPricePointWhenEmptyBidsLowToHigh(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 2)

	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13436", action: "ASK", size: 15, pair: "BTC/ETH", price: 1}
	err := orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	if orderbook.asks[1].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.asks[1].nextOrder.size)
	}
	order = Order{id: "13435", action: "ASK", size: 15, pair: "BTC/ETH", price: 2}
	err = orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	if orderbook.asks[2].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.asks[2].nextOrder.size)
	}
	order = Order{id: "13434", action: "ASK", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Ask(order)
	if orderbook.asks[3].nextOrder.size != 15 {
		t.Fatalf("Orderbook size should be equal to %v, got %v", 15, orderbook.asks[3].nextOrder.size)
	}
}

func TestOrderbookInsertOrderRowAsksUnfilled(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 10)
	var order = Order{id: "13436", action: "ASK", size: 15, pair: "BTC/ETH", price: 3}
	err := orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	order = Order{id: "13437", action: "ASK", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	order = Order{id: "13438", action: "ASK", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	order = Order{id: "13439", action: "ASK", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	if orderbook.asks[3].nextOrder.nextOrder.nextOrder.size != 15 {
		t.Fatalf("Ask at this order should be 15")
	}
}

func TestOrderbookInsertOrderRowBidsUnfilled(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 10)
	var order = Order{id: "13436", action: "BID", size: 15, pair: "BTC/ETH", price: 3}
	err := orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	order = Order{id: "13437", action: "BID", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	order = Order{id: "13438", action: "BID", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	order = Order{id: "13439", action: "BID", size: 15, pair: "BTC/ETH", price: 3}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	if orderbook.bids[3].nextOrder.nextOrder.nextOrder.size != 15 {
		t.Fatalf("Ask at this order should be 15")
	}
}

func TestOrderbookSingleMatchNoPricePointTraversalFilled(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 5)
	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13436", action: "ASK", size: 15, pair: "BTC/ETH", price: 4}
	err := orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	order = Order{id: "13437", action: "BID", size: 15, pair: "BTC/ETH", price: 4}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}

	var pricePoint int = 1
	// the bid order should never get placed sense it gets matched
	if orderbook.bids[pricePoint].nextOrder.size != 0 {
		t.Fatalf("Bids at size %v should be zero sense it should of matched", pricePoint)
	}

	// the ask order should be filled by the incoming bid order
	if orderbook.asks[pricePoint].nextOrder.size != 0 {
		t.Fatalf("Asks at price point %v should be zero sense it should of matched with the incoming bid order", pricePoint)
	}
}

func TestOrderbookSingleMatchPricePointTraversalFilled(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 2)
	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13436", action: "ASK", size: 15, pair: "BTC/ETH", price: 1}
	err := orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	order = Order{id: "13437", action: "BID", size: 15, pair: "BTC/ETH", price: 2}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}

	// the bid order should never get placed sense it gets matched
	if orderbook.bids[2].nextOrder.size != 0 {
		t.Fatalf("Bids at size %v should be zero sense it should of matched", 1)
	}

	// the ask order should be filled by the incoming bid order
	if orderbook.asks[1].nextOrder.size != 0 {
		t.Fatalf("Asks at price point %v should be zero sense it should of matched with the incoming bid order", 2)
	}
}

func TestOrderbookSingleMatchPricePointTraversalPartialFilled(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 5, 0, 2)
	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13436", action: "ASK", size: 15, pair: "BTC/ETH", price: 1}
	err := orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	order = Order{id: "13437", action: "BID", size: 15, pair: "BTC/ETH", price: 2}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	order = Order{id: "13438", action: "BID", size: 15, pair: "BTC/ETH", price: 2}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}

	// the ask order should be filled by the incoming bid order
	if orderbook.asks[1].nextOrder.size != 0 {
		t.Fatalf("Asks at price point %v should be zero sense it should of matched with the incoming bid order", 2)
	}

	// the order book on the BID side should have an order of 15
	if orderbook.bids[2].nextOrder.size != 15 {
		t.Fatalf("Bids at size %v should be zero sense it should of matched", 1)
	}
}

func TestCancelOrder(t *testing.T) {
	orderbook := NewOrderbook("BTC/ETH", 5, 0, 2)
	order := Order{id: "13437", action: "ASK", size: 15, pair: "BTC/ETH", price: 2}
	// stock the orderbook
	orderbook.Ask(order)
	order = Order{id: "13438", action: "CANCEL", size: 15, pair: "BTC/ETH", price: 2}
	status := orderbook.Cancel(order)
	if status != canceledOrderStatusFail {
		t.Fatalf("Failed received status: %v", status)
	}
}

/*
func TestOrderbookEmptyOrderBookBids(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 2, 0, 2)

	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13435", action: "BID", size: 15, pair: "BTC/ETH", price: 1}
	err := orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}

	// stock the orderbook with a second order - orderbook is not empty it should
	// traverse to the next node
	order = Order{id: "13436", action: "BID", size: 20, pair: "BTC/ETH", price: 1}
	err = orderbook.Bid(order)
	if err != nil {
		t.Fatalf("Order should be action BID: %v\n", err)
	}
	orderbook.PrintBids(1)
}
*/

/*
func TestOrderbookEmptyOrderBookAsks(t *testing.T) {
	var orderbook = NewOrderbook("BTC/ETH", 3, 0, 2)

	// stock the orderbook with the first order - order book is empty
	var order = Order{id: "13435", action: "ASK", size: 12, pair: "BTC/ETH", price: 1}
	err := orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}

	// stock the orderbook with a second order - orderbook is not empty it should
	// traverse to the next node
	err = orderbook.Ask(order)
	if err != nil {
		t.Fatalf("Order should be action ASK: %v\n", err)
	}
	orderbook.PrintAsks(1)
}
*/
