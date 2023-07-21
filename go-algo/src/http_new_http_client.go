/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Create new client
 */

package src

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Client *http.Client
}

func (c *Client) Init() {
	c.Client = &http.Client{
		Timeout: time.Second * 5,
	}
}

func (c *Client) GetCEPWithNewClient() {
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/58429073/json/", nil)

	if err != nil {
		panic(err)
	}

	req.SetBasicAuth("user", "password")

	res, err := c.Client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func (c *Client) PostCEPWithNewClient() {
	payload := bytes.NewBuffer([]byte(`{"name": "Tiago"}`))

	req, err := http.NewRequest("POST", "https://httpbin.org/post", payload)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	res, err := c.Client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
