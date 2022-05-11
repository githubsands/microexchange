package exchange

type O interface {
	Ask(Order)
	Bid(Order)
	Modify(Order)
}
