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

	if req.Reset {
		res.Reset()
	} else {
		if req.Allowed(*res) {
			res.Player()
		}
	}

	msg, err = json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("API-Response: ")
	fmt.Println(res)
	return msg
}
