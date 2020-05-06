package hassky

import (
	"log"
	"net/url"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
)

// Version is constant for the package
const Version = "0.0.1-dev"

// New function creates a New Hassky client
func New(host string, authToken string, useSSL bool) (*Client, error) {

	wsScheme := "ws"
	httpScheme := "http"

	if useSSL {
		wsScheme = "wss"
		httpScheme = "https"
	}

	httpClient := resty.New()
	httpClient.SetHeader("Accept", "application/json")
	httpClient.SetAuthToken(authToken)
	httpClient.SetHostURL(httpScheme + "://" + host)

	u := url.URL{Scheme: wsScheme, Host: host, Path: "/api/websocket"}

	wsClient, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return nil, err
	}
	//	defer wsClient.Close()

	_, _, _ = wsClient.ReadMessage()

	authMessage := gabs.New()
	authMessage.Set("auth", "type")
	authMessage.Set(authToken, "access_token")

	connectionErr := wsClient.WriteJSON(authMessage)
	if connectionErr != nil {
		log.Fatal("write:", connectionErr)
	}

	_, _, _ = wsClient.ReadMessage()
	// fmt.Printf("Auth Phase #2 >> %+v\n", string(message))

	client := &Client{
		wsClient:   wsClient,
		httpClient: httpClient,
	}

	return client, nil
}
