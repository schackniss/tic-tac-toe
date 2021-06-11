package websocket

import "tic-tac-toe/game"

// Hub verwaltet die Menge der aktiven Clients und sendet Nachrichten an die Clients.
type Hub struct {
	// Registrierte Clients.
	clients map[*Client]bool

	// Eingehende Nachrichten von den Clients.
	broadcast chan []byte

	// Registrieren von Anfragen der Clients.
	register chan *Client

	// Aufheben der Registrierung von Anfragen der Clients.
	unregister chan *Client
}

// NewHub ist der Konstruktor zur Initialisierung eines neuen Hubs.
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run startet den Hub, um die Anfragen (register, unregister, message) der Clients zu verwalten.
//
// Die Anwendung startet eine Go-Routine für die Funktion Run.
func (h *Hub) Run() {
	ttt := game.TicTacToe{
		NextPlayer: 1,
		Field:      [3]string{"   ", "   ", "   "},
	}

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true

			// Aktuellen Spielstand an neuen Client schicken.
			select {
			case client.send <- welcome(&ttt):
			default:
				close(client.send)
				delete(h.clients, client)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.broadcast:

			// Nachricht entgegennehmen, verarbeiten und Antwort an alle Clients schicken.
			message = handleMessage(message, &ttt)

			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
