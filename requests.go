package hassky

import (
	"regexp"

	"github.com/Jeffail/gabs/v2"
)

// Request struct
type Request struct {
	method  string
	urlPath string
	body    *gabs.Container
	client  *Client
}

// Ping method sets a type of the request to "ping"
// Hassky current request instance.
//    client.R().Ping()
//
//   Result:
//      {
//        "id": 123,
//        "type": "ping",
//      }
func (req *Request) Ping() *Request {
	req.method = "websocket"
	req.body.Set("ping", "type")
	return req
}

// GetStatesWS method sets a type of the request to "get_states"
// Hassky current request instance.
//    client.R().GetStatesWS()
//
//   Result:
//      {
//        "id": 123,
//        "type": "get_states",
//      }
func (req *Request) GetStatesWS() *Request {
	req.method = "websocket"
	req.body.Set("get_states", "type")
	return req
}

// GetConfigWS method sets a type of the request to "get_config"
// Hassky current request instance.
//    client.R().GetConfig()
//
//   Result:
//      {
//        "id": 123,
//        "type": "get_config",
//      }
func (req *Request) GetConfigWS() *Request {
	req.method = "websocket"
	req.body.Set("get_config", "type")
	return req
}

// GetServicesWS method sets a type of the request to "get_services"
// Hassky current request instance.
//    client.R().GetServicesWS()
//
//   Result:
//      {
//        "id": 123,
//        "type": "get_services",
//      }
func (req *Request) GetServicesWS() *Request {
	req.method = "websocket"
	req.body.Set("get_services", "type")
	return req
}

// GetPanels method sets a type of the request to "get_panels"
// Hassky current request instance.
//    client.R().GetPanels()
//
//   Result:
//      {
//        "id": 123,
//        "type": "get_panels",
//      }
func (req *Request) GetPanels() *Request {
	req.method = "websocket"
	req.body.Set("get_panels", "type")
	return req
}

// GetThumbnailMediaPlayer method sets a type of the request to "media_player_thumbnail"
// Hassky current request instance.
//    client.R().GetThumbnailMediaPlayer()
//
//   Result:
//      {
//        "id": 123,
//        "type": "media_player_thumbnail",
//      }
func (req *Request) GetThumbnailMediaPlayer(entityID string) *Request {
	req.method = "websocket"
	req.body.Set("media_player_thumbnail", "type")
	req.body.Set(entityID, "entity_id")
	return req
}

// GetThumbnailCamera method sets a type of the request to "camera_thumbnail"
// Hassky current request instance.
//    client.R().GetThumbnailCamera()
//
//   Result:
//      {
//        "id": 123,
//        "type": "camera_thumbnail",
//      }
func (req *Request) GetThumbnailCamera(entityID string) *Request {
	req.method = "websocket"
	req.body.Set("camera_thumbnail", "type")
	req.body.Set(entityID, "entity_id")
	return req
}

// IsAPIRunning is a functon
// {
//     "message": "API running."
// }
func (req *Request) IsAPIRunning() *Request {
	req.method = "get"
	req.urlPath = "/api/"
	return req
}

// DiscoveryInfo is a functon
// {
//     "base_url": "http://192.168.0.2:8123",
//     "location_name": "Home",
//     "requires_api_password": true,
//     "version": "0.56.2"
// }
func (req *Request) DiscoveryInfo() *Request {
	req.method = "get"
	req.urlPath = "/api/discovery_info"
	return req
}

// GetConfig is a functon
// {
//     "message": "Entity not found."
// }
func (req *Request) GetConfig() *Request {
	req.method = "get"
	req.urlPath = "/api/config"
	return req
}

// GetEventListeners is a functon
// [
//     {
//       "event": "state_changed",
//       "listener_count": 5
//     },
//     {
//       "event": "time_changed",
//       "listener_count": 2
//     }
// ]
func (req *Request) GetEventListeners() *Request {
	req.method = "get"
	req.urlPath = "/api/events"
	return req
}

// GetServices is a functon
// [
//     {
//       "domain": "browser",
//       "services": [
//         "browse_url"
//       ]
//     },
//     {
//       "domain": "keyboard",
//       "services": [
//         "volume_up",
//         "volume_down"
//       ]
//     }
// ]
func (req *Request) GetServices() *Request {
	req.method = "get"
	req.urlPath = "/api/services"
	return req
}

// GetStates is a functon
// [
//     {
//         "attributes": {},
//         "entity_id": "sun.sun",
//         "last_changed": "2016-05-30T21:43:32.418320+00:00",
//         "state": "below_horizon"
//     },
//     {
//         "attributes": {},
//         "entity_id": "process.Dropbox",
//         "last_changed": "22016-05-30T21:43:32.418320+00:00",
//         "state": "on"
//     }
// ]
func (req *Request) GetStates() *Request {
	req.method = "get"
	req.urlPath = "/api/states"
	return req
}

// GetState is a functon
// {
// 	"attributes":{
// 	   "azimuth":336.34,
// 	   "elevation":-17.67,
// 	   "friendly_name":"Sun",
// 	   "next_rising":"2016-05-31T03:39:14+00:00",
// 	   "next_setting":"2016-05-31T19:16:42+00:00"
// 	},
// 	"entity_id":"sun.sun",
// 	"last_changed":"2016-05-30T21:43:29.204838+00:00",
// 	"last_updated":"2016-05-30T21:50:30.529465+00:00",
// 	"state":"below_horizon"
//  }
func (req *Request) GetState(entityID string) *Request {
	req.method = "get"
	req.urlPath = "/api/states/" + entityID
	return req
}

// GetErrors is a functon
// 15-12-20 11:02:50 homeassistant.components.recorder: Found unfinished sessions
// 15-12-20 11:03:03 netdisco.ssdp: Error fetching description at http://192.168.1.1:8200/rootDesc.xml
// 15-12-20 11:04:36 homeassistant.components.alexa: Received unknown intent HelpIntent
func (req *Request) GetErrors() *Request {
	req.method = "get"
	req.urlPath = "/api/error_log"
	return req
}

// CallService is a function
func (req *Request) CallService(service string, entityID string) *Request {
	r, _ := regexp.Compile(`(\w+)[\.\/]{1}(\w+)`)
	serviceData := r.FindStringSubmatch(service)
	if len(serviceData) == 3 {
		req.urlPath = "/api/services/" + serviceData[1] + "/" + serviceData[2]
	}

	matched, _ := regexp.MatchString(`^(\w+)[\.]{1}(\w+)$`, entityID)
	if matched {
		req.body.Set(entityID, "entity_id")
	}

	req.method = "post"

	return req
}

// CallServiceWS is a function
func (req *Request) CallServiceWS(service string, entityID string) *Request {
	r, _ := regexp.Compile(`(\w+)[\.\/]{1}(\w+)`)
	serviceData := r.FindStringSubmatch(service)

	if len(serviceData) == 3 {
		req.body.Set(serviceData[1], "domain")
		req.body.Set(serviceData[2], "service")
	}

	matched, _ := regexp.MatchString(`^(\w+)[\.]{1}(\w+)$`, entityID)
	if matched {
		req.body.Set(entityID, "service_data", "entity_id")
	}

	req.method = "websocket"
	req.body.Set("call_service", "type")

	return req
}

// SetParameters method sets a domain of a calling service
// Hassky current request instance.
//    client.R().SetParameters(map[string]interface{"entity_id": "light.kitchen"})
//
//   Result:
//      {
//        "id": 123,
//        "type": "call_service",
//        "domain": "light",
//        "service": "turn_on",
//        // Optional
//        "service_data": {
//          "entity_id": "light.kitchen"
//        }
//      }
func (req *Request) SetParameters(params map[string]interface{}) *Request {
	for k, v := range params {
		if req.method == "websocket" {
			req.body.Set(v, "service_data", k)
		} else {
			req.body.Set(v, k)
		}
	}
	return req
}

// Subscribe is a function
func (req *Request) Subscribe(eventType string) *Request {
	req.method = "websocket"
	req.body.Set("subscribe_events", "type")
	req.body.Set(eventType, "event_type")
	return req
}

// Unsubscribe method sets a type of the request to "unsubscribe_events"
// Hassky current request instance.
//    client.R().Unsubscribe()
//
//   Result:
//      {
//        "id": 123,
//        "type": "unsubscribe_events",
//        "subscription": 234
//      }
func (req *Request) Unsubscribe(subscribtionID int) *Request {
	req.method = "websocket"
	req.body.Set("unsubscribe_events", "type")
	req.body.Set(subscribtionID, "subscription")
	return req
}

// SetState method sets a type of the request to "set_panels"
// Hassky current request instance.
//    client.R().SetState()
//
//   Result:
//     {
//       "state": "below_horizon",
//       "attributes": {
//         "next_rising":"2016-05-31T03:39:14+00:00",
//         "next_setting":"2016-05-31T19:16:42+00:00"
//       }
//     }
func (req *Request) SetState(entityID string, state interface{}) *Request {

	matched, _ := regexp.MatchString(`^(\w+)[\.]{1}(\w+)$`, entityID)
	if matched {
		req.urlPath = "/api/states/" + entityID
	}

	req.method = "post"
	req.body.Set(state, "state")

	return req
}

// SetAttributes method sets a domain of a calling service
// Hassky current request instance.
//    client.R().SetAttributes(map[string]interface{"entity_id": "light.kitchen"})
//
//   Result:
//      {
//        "id": 123,
//        "type": "call_service",
//        "domain": "light",
//        "service": "turn_on",
//        // Optional
//        "service_data": {
//          "entity_id": "light.kitchen"
//        }
//      }
func (req *Request) SetAttributes(attributes map[string]interface{}) *Request {
	for k, v := range attributes {
		req.body.Set(v, "attributes", k)
	}
	return req
}

// FireEvent is a function
// {
//     "message": "Event download_file fired."
// }
func (req *Request) FireEvent(eventType string) *Request {
	req.method = "post"
	req.urlPath = "/api/events/" + eventType

	return req
}

// SetEventData method sets a domain of a calling service
// Hassky current request instance.
//    client.R().SetEventData(map[string]interface{"direction": "up"})
//
//   Result:
//      {
//        "id": 123,
//        "type": "call_service",
//        "domain": "light",
//        "service": "turn_on",
//        // Optional
//        "service_data": {
//          "entity_id": "light.kitchen"
//        }
//      }
func (req *Request) SetEventData(attributes map[string]interface{}) *Request {
	for k, v := range attributes {
		req.body.Set(v, k)
	}
	return req
}

// CheckConfig is a fuction
// {
//     "errors": null,
//     "result": "valid"
// }
// {
//     "errors": "Integration not found: frontend:",
//     "result": "invalid"
// }
func (req *Request) CheckConfig() *Request {
	req.method = "post"
	req.urlPath = "/api/config/core/check_config"
	return req
}

// Exec is a method
func (req *Request) Exec() (string, bool, error) {
	resp, ok, err := req.client.execute(req)
	return resp, ok, err
}

// GetHistory is a functon
// func (req *Request) GetHistory(entityIDs ...string) *Request {
// 	req.method = "get"
// 	req.urlPath = "/api/history/period/<timestamp>"
// 	return req
// }
