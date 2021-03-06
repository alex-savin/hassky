package hassky

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/Jeffail/gabs/v2"
	log "github.com/sirupsen/logrus"
)

// Response is a struct
type Response struct {
	OK          bool
	Error       string
	Response    *gabs.Container
	ResponseRaw []byte
}

// ResponseWS is a struct
type ResponseWS struct {
	ID      int             `json:"id"`
	Type    string          `json:"type"`
	Success bool            `json:"success,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   struct {
		ErrorCode    string `json:"code,omitempty"`
		ErrorMessage string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}

// Parse method which parsing output of executed commands
func (res *Response) Parse(req *Request, resp []byte) *Response {
	switch req.method {
	case "websocket":
		response := &ResponseWS{}
		_ = json.Unmarshal(resp, response)

		log.Debugf("WS >> %#v", &response.Result)

		switch response.Type {
		case "result":
			res.Response, _ = gabs.ParseJSON(response.Result)

			if response.Success == true {
				res.OK = true

			} else {
				res.OK = false
				res.Error = response.Error.ErrorMessage
			}
		case "pong":
		}
		switch req.body.Path("type").Data().(string) {
		case "subscribe_events":
		case "unsubscribe_events":
		case "call_service":
		case "get_states":
		case "get_config":
		case "get_services":
		case "get_panels":
		case "camera_thumbnail":
		case "media_player_thumbnail":
		case "ping":
		default:
			log.Fatal("WEBSOCKET >> COULDN'T RECOGNIZE")
		}
	case "post":
		res.Response, _ = gabs.ParseJSON(resp)

		//		switch strings.Split(req.urlPath, "/")[strings.Count(req.urlPath, "/")] {
		switch strings.Split(req.urlPath, "/")[2] {
		case "states":
			if matched, _ := regexp.MatchString(`^(\w+)[\.]{1}(\w+)$`, strings.Split(req.urlPath, "/")[strings.Count(req.urlPath, "/")]); matched == true {
				log.Info("POST >> ENTITY ID IS DETECTED")
			}
		case "events":
			if matched, _ := regexp.MatchString(`^Event\s(\w+)\sfired.$`, strings.Split(req.urlPath, "/")[strings.Count(req.urlPath, "/")]); matched == true {
				log.Info("POST >> EVENT WAS SUCESSFULLY FIRED")
			}
		case "services":
			fmt.Println(res.Response)
			for _, obj := range res.Response.Children() {
				fmt.Printf("%+v", obj)
			}
		case "config":
			if value, ok := res.Response.Path("result").Data().(string); ok == true && value == "valid" {
				log.Info("POST >> TRUE")
				res.OK = true
			} else {
				res.OK = false
				res.Error, _ = res.Response.Path("errors").Data().(string)
			}
		default:
			log.Fatal("POST >> COULDN'T RECOGNIZE")
		}
	case "get":
		res.Response, _ = gabs.ParseJSON(resp)

		switch strings.Split(req.urlPath, "/")[2] {
		case "events":
			for _, child := range res.Response.Children() {
				fmt.Println(child.Data().(map[string]interface{})["event"], child.Data().(map[string]interface{})["listener_count"])
			}
		case "services":
		case "config":
			fmt.Println(res.Response)
		case "discovery_info":
			fmt.Println(res.Response)
		case "error_log":
		case "states":
			if matched, _ := regexp.MatchString(`^(\w+)[\.]{1}(\w+)$`, strings.Split(req.urlPath, "/")[strings.Count(req.urlPath, "/")]); matched == true {
				log.Info("GET >> ENTITY ID IS DETECTED")
			}
		case "camera_proxy":
		case "history":
		default:
			log.Fatal("GET >> COULDN'T RECOGNIZE")
		}
	}
	return res
}
