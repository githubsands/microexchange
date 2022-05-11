package exchange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Dispatcher struct {
	events chan event

	asyncDispatchFrequency time.Duration
	transactor             func(*http.Request) (*http.Response, error)
}

func NewDispatcher() (*Dispatcher, chan event) {
	d := new(Dispatcher)
	d.events = make(chan event, 3000)
	d.transactor = http.DefaultClient.Do
	client := http.Client{}
	client.Timeout = 1
	d.transactor = client.Do
	go d.send()
	return d, d.events
}

func (d *Dispatcher) send() {
	for {
		go d.transact()
		time.Sleep(d.asyncDispatchFrequency)
	}
}

func (d *Dispatcher) transact() {
	for v := range d.events {
		go func() {
			buf, err := json.Marshal(v)
			reader := bytes.NewReader(buf)
			req, err := http.NewRequest("POST", v.ip, reader)
			if err != nil {
				fmt.Println(err)
			}
			resp, err := d.transactor(req)
			if err != nil {
				fmt.Println(err)
			}

			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}()
	}
}
