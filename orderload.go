package exchange

import (
	"sort"
	"time"
)

type OrderLoad struct {
	orderQueueSize int
	orderQueue     chan Order

	sortedOrderQueueSize int
	sortedOrderQueue     chan Order
	sortFrequency        time.Duration
	sortBatchSize        int

	submitFrequency time.Duration
	submitter       chan Order
	orderbook       *Orderbook
}

type orders []Order

func (os orders) Len() int {
	return len(os)
}

func (os orders) Less(i, j int) bool {
	var less bool = false
	switch {
	case os[i].price < os[j].price:
		less = true
		fallthrough
	case os[i].time < os[j].time:
		less = true
	}

	return less
}

func (os orders) Swap(i, j int) {
	os[i], os[j] = os[j], os[i]
}

func NewOrderLoad() (*OrderLoad, chan Order) {
	ol := new(OrderLoad)
	ol.orderQueueSize = 3000000
	ol.orderQueue = make(chan Order, ol.orderQueueSize)             // order queue channel here
	ol.sortedOrderQueueSize = 3000                                  // order queue channel here
	ol.sortedOrderQueue = make(chan Order, ol.sortedOrderQueueSize) // order queue channel here
	ol.submitFrequency = 1 * time.Microsecond                       // determines how often an order is taken from the queue and is submitted to the orderbook
	ol.submitter = make(chan Order, 1)                              // creating blocking channel here
	go ol.queue()
	go ol.sort()
	go ol.submit()
	return ol, ol.orderQueue
}

func (ol *OrderLoad) sort() {
	for {
		ol.sorter()
		time.Sleep(ol.sortFrequency)
	}
}

// sorter sorts orders from the queue by price then priority in batches of size sortBatchSize
func (ol *OrderLoad) sorter() {
	orders := make(orders, 0, ol.sortBatchSize)
	i := 0
	for order := range ol.orderQueue {
		i++
		if i > ol.sortBatchSize+1 {
			break
		}
		orders = append(orders, order)
	}

	sort.Sort(orders)

	for _, sortedOrder := range orders {
		ol.sortedOrderQueue <- sortedOrder
	}
}

func (ol *OrderLoad) queue() {
	for order := range ol.sortedOrderQueue {
		time.Sleep(ol.submitFrequency)
		ol.submitter <- order
	}
}

// submit, submits the order to the orderbook the time to submit a new order from the queue depends on
// the orderbook's time complexity
func (ol *OrderLoad) submit() {
	for {
		select {
		case order := <-ol.submitter:
			switch order.action {
			case "BID":
				ol.orderbook.Bid(order)
			case "ASK":
				ol.orderbook.Ask(order)
			case "CANCEL":
				//		ol.orderbook.Cancel(order)
			default:
			}
		default:
		}
	}
}
