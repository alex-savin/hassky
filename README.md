# Hassky
Is a simple API client to interact with Home Assistant instance (https://www.home-assistant.io/) via HTTP(s) / Websocket, written in Go

## News
  * v0.0.1-dev First public release on May 6, 2020

## Features
  * Simple and chainable methods for settings and request

## Installation
```bash
# Go Modules
require github.com/alex-savin/hassky
```

## Usage
The following samples will assist you to become as comfortable as possible with hassky library.
```go
// Import hassky into your code and refer it as `hassky`.
import "github.com/alex-savin/hassky"
```

#### Simple PING request
```go
// Create a Hassky Client
h, _ := hassky.New("[HASS-IP-ADDESS]:[PORT]", "[HASS-API-TOKEN]", "[USE-SSL]", "[LOG-LEVEL]")

h.R().Ping().
      Exec()
```

#### Simple CALL SERVICE request
```go
h.R().CallService("light.toggle", "light.porch_01").
      Exec()
```

#### Advanced CALL SERVICE request
```go
h.R().CallService("light.toggle", "light.porch_01").
      SetParameters(map[string]interface{}{
        "transition":15,
        "brightness_pct":100
      }).
      Exec()
```
