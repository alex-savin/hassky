package hassky

import (
	"net/url"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// Version is constant for the package
const Version = "0.1.0-dev"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

// New function creates a New Hassky client
func New(host string, authToken string, useSSL bool, logLevel string) (*Client, error) {

	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		log.Warnf("[HASSKY] Invalid log level supplied: '%s'", logLevel)
	}

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
		log.Fatalf("WSCLIENT ERROR >> %s", err)
		return nil, err
	}
	//	defer wsClient.Close()

	_, _, _ = wsClient.ReadMessage()

	authMessage := gabs.New()
	authMessage.Set("auth", "type")
	authMessage.Set(authToken, "access_token")

	err = wsClient.WriteJSON(authMessage)
	if err != nil {
		log.Fatalf("WSCLIENT ERROR >> %s", err)
	}

	_, _, _ = wsClient.ReadMessage()

	client := &Client{
		wsClient:   wsClient,
		httpClient: httpClient,
	}

	return client, nil
}
