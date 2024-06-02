package internal

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var writeWait = 10 * time.Second
var pongWait = 60 * time.Second
var pingPeriod = (pongWait * 9) / 10

// Define the WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	WriteBufferPool: &sync.Pool{},
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		fmt.Println(origin)
		return origin == "http://localhost:8081"
	},
}

func wsprocess(ws *websocket.Conn) {
	// defer ws.Close()
	for {
		// Read message from WebSocket
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read:", err, messageType)
			break
		}
		log.Printf("Received: %s", message)

		// Echo the message back to the client
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}

func (s *Server) WsHandler() echo.HandlerFunc {
	// Define a handler for WebSocket connections
	return func(c echo.Context) error {
		// Upgrade the HTTP connection to a WebSocket connection
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Println("Upgrade:", err)
			return err
		}
		ws.SetWriteDeadline(time.Now().Add(pongWait))
		ws.SetPongHandler(func(string) error { ws.SetWriteDeadline(time.Now().Add(pongWait)); fmt.Println("pongers"); return nil })
		ws.SetPingHandler(func(appData string) error {
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			ws.WriteMessage(websocket.PingMessage, nil)
			return nil
		})
		s.WS = ws
		go wsprocess(s.WS)
		return nil
	}
}