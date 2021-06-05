package websocket

import "tic-tac-toe/game"

// handleMessage verarbeitet die Anfrage/Nachricht eines Clients und gibt eine Antwort zurück.
//
// Die Nachricht ist der vom Client durchgeführte Spielzug. handleMessage ruft die Spiellogik auf und gibt den neuen Spielstand zurück.
func handleMessage(msg []byte, res *game.TicTacToe) []byte {
	return msg
}
