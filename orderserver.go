package exchange

import (
	"encoding/json"
	"net/http"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/valyala/fasthttp"
)

type client struct {
	ip     string
	id     string
	orders int
}

type OrderServer struct {
	clientManager *clientManager

	orders  chan<- Order
	clients chan<- client

	cleanup <-chan string
}

func NewOrderServer(orders chan Order) *OrderServer {
	or := new(OrderServer)
	clientManager, clients := NewClientManager()
	or.orders = orders
	or.clientManager = clientManager
	or.clients = clients
	or.cleanup = make(chan string, 100)
	return or
}

func (or *OrderServer) HTTPHandler(ctx *fasthttp.RequestCtx) {
	resp := &ctx.Response
	buf := ctx.Request.Body()
	order := Order{}
	err := json.Unmarshal(buf, &order)
	if err != nil {
		resp.Header.SetStatusCode(http.StatusInternalServerError)
	}
	or.orders <- order
	or.clients <- client{ip: ctx.RemoteIP().String(), id: order.id}
	resp.Header.SetStatusCode(http.StatusAccepted)
}

// client stores all recent clients(traders) interacting with the server
type clientManager struct {
	ips cmap.ConcurrentMap

	clientQueue chan client
	cleanup     <-chan string
}

func NewClientManager() (*clientManager, chan client) {
	cm := new(clientManager)
	cm.ips = cmap.New()
	cm.clientQueue = make(chan client, 3000)
	cm.ips = cmap.New()
	return cm, cm.clientQueue
}

// register takes a client from a http go routine and adds it to clientManager
func (c *clientManager) register() {
	for {
		for client := range c.clientQueue {
			c.ips.Set(client.id, client.ip)
		}
	}
}

/*
func (c *clientManager) cleanUp() {





}



func (c *clientManager) cleanUp() {
	for {
		select {
		case client := <-c.cleanup:
			c.ips.Remove(client)
		case
		}
	}
}

func (c *clientManager) removeOrder() {



}


func (c *clientManager) addOrder(client {
	for {
		select {
		case client := <-c.cleanup:
			c.ips.Remove(client)
		case
		}
	}
}
*/
