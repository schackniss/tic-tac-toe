// websocket dient der Implementierung von Websockets für das Spiel Tic Tac Toe
// Die Implementierung basiert auf dem Repository https://github.com/gorilla/websocket.
package websocket

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Deinition von Konstanten.
const (
	// Zeit, die für das Schreiben einer Nachricht an die Gegenstelle zur Verfügung steht.
	writeWait = 10 * time.Second

	// Zeit, die zum Lesen der nächsten Pong-Nachricht von der Gegenstelle erlaubt ist.
	pongWait = 60 * time.Second

	// Sendet Pings an die Gegenstelle mit dieser Zeitspanne. Muss kleiner sein als pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximal zulässige Nachrichtengröße von der Gegenstelle.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Definition der Lese- und Schreib-Puffer-Größe.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client ist ein Verbinder zwischen der Websocket-Verbindung und dem Hub.
type Client struct {
	hub *Hub

	// Die Websocket-Verbindung.
	conn *websocket.Conn

	// Gepufferter Kanal für ausgehende Nachrichten.
	send chan []byte
}

// readPump pumpt Nachrichten von der Websocket-Verbindung zum Hub.
//
// Die Anwendung führt readPump, je Verbindung, in einer Go-Routine aus. Die Anwendung stellt sicher, dass es höchstens einen Leser auf einer Verbindung gibt, indem sie alle Lesevorgänge aus dieser Go-Routine ausführt.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumpt Nachrichten vom Hub an die Websocket-Verbindung.
//
// Für jede Verbindung wird eine Go-Routine gestartet, die writePump ausführt. Die Anwendung stellt sicher, dass es höchstens einen Schreiber für eine Verbindung gibt, indem alle Schreibvorgänge von dieser Go-Routine aus ausgeführt werden.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Der Hub hat den Kanal geschlossen.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Hinzufügen von Chatnachrichten aus der Warteschlange zur aktuellen Websocket-Nachricht.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Serve verarbeitet Websocket-Anfragen von der Gegenstelle.
func Serve(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Erlaubt das Sammeln des vom Aufrufer referenzierten Speichers, indem alle Arbeiten in neuen Go-Routinen ausgeführt werden.
	go client.writePump()
	go client.readPump()
}
