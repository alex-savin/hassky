package hassky

import (
	"errors"
	"time"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var requestID int = 0

// Client is a struct
type Client struct {
	wsClient   *websocket.Conn
	httpClient *resty.Client
}

// R method creates a new Request instances to be use lates while the firing a request with client
func (c *Client) R() *Request {
	req := &Request{
		body:   gabs.New(),
		client: c,
	}
	return req
}

// Exec method executes a Client instance with the Request
func (c *Client) execute(req *Request) (string, bool, error) {
	defer timeTrack(time.Now(), "Executing")

	res := &Response{}

	switch req.method {

	case "websocket":
		requestID++
		req.body.Set(requestID, "id")

		err := c.wsClient.WriteJSON(req.body)
		if err != nil {
			log.Fatal("ERROR >> %s\n", err)
		}

		_, message, _ := c.wsClient.ReadMessage()
		if len(string(message)) > 0 {
			res.Parse(req, message)
			return string(message), true, nil
		}

		return string(message), true, nil

	case "post":
		if isNil(req.body) {
			resp, err := c.httpClient.R().Post(req.urlPath)
			if err != nil {
				log.Fatal("ERROR >> %s", err)
			}

			if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
				log.Fatal("PARSE ERROR HERE >> %s", err)
			}
			res.OK = true
			res.Parse(req, resp.Body())
		} else {
			resp, err := c.httpClient.R().SetBody(req.body).Post(req.urlPath)
			if err != nil {
				log.Fatal("ERROR >> %s", err)
			}
			if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
				log.Fatal("PARSE ERROR HERE >> %s", err)
			}
			res.OK = true
			res.Parse(req, resp.Body())
		}
		return "", true, nil

	case "get":
		resp, err := c.httpClient.R().Get(req.urlPath)
		if err != nil {
			log.Fatal("ERROR >> %s", err)
		}
		if resp.StatusCode() != 200 {
			log.Fatal("PARSE ERROR HERE >> %s", err)
		}
		res.OK = true
		res.Parse(req, resp.Body())

		return string(resp.Body()), true, nil

	default:
		return "", false, errors.New("Something went wrong")
	}
}
