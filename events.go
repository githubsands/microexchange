package exchange

/*
func (p *publish) order() &publishEvent {} // publish order or cancel acknowledgement format
func (p *publish) changes() &publishEvent {} // publish changes in Top Of Book per side using format, use ‘-‘ for side elimination
func (p *publish) reject() &publishEvent {} // publish rejects for orders that would make or book crossed
func (p *publish) trades() &publishEvent {} // publish trades (matched orders) format
*/

type event struct {
	userOrderID int `json:'userOrderID'`
	userIDSell  int `json:'userIDSell'`

	price    int `json:'price'`
	quantity int `json:'quantity'`

	ip string `json:'ip'`
}

type eventHandler struct {
	orderbookEvents <-chan event
	dispatchs       chan<- event
	clientEvents    chan<- event
}

func (e *eventHandler) handle() {
	for {
		select {
		case _ = <-e.orderbookEvents:
		default:
		}
	}
}
