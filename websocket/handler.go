package websocket

import (
	"encoding/json"
	"fmt"
	"tic-tac-toe/game"
)

// handleMessage verarbeitet die Anfrage/Nachricht eines Clients und gibt eine Antwort zurück.
//
// Die Nachricht ist der vom Client durchgeführte Spielzug. handleMessage ruft die Spiellogik auf und gibt den neuen Spielstand zurück.
func handleMessage(msg []byte, res *game.TicTacToe) []byte {
	req := game.Move{}
	err := json.Unmarshal(msg, &req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("API-Request: ")
	fmt.Println(req)

	// Prüfen, ob Reset angefordert wurde
	if req.Reset {
		res.Reset()
	} else {
		// Prüfen, ob der Zug erlaubt ist (noch kein X oder O auf Feld)
		if req.Allowed(*res) {
			res.Player()
		}
	}

	msg, err = json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("API-Response: ")
	fmt.Println(*res)
	return msg
}

// welcome formatiert den aktuellen Spielstand als JSON type []byte.
func welcome(ttt *game.TicTacToe) []byte {
	msg, err := json.Marshal(ttt)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("Willkommen: ")
	fmt.Println(*ttt)
	return msg
}
