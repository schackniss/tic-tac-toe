// Das Paket 'game' beinhaltet die Tic Tac Toe Spiellogik.
package game

import "strings"

// TicTacToe repräsentiert das Spiel und ist zugleich das API-Antwort-Format.
type TicTacToe struct {
	Finished   bool      `json:"finished"`
	Winner     int       `json:"winner"`
	NextPlayer int       `json:"nextPlayer"`
	Field      [3]string `json:"field"`
}

// Reset setzt die übergebene Variable auf Spielbeginn.
func (ttt *TicTacToe) Reset() {
	ttt.Finished = false
	ttt.Winner = 0
	ttt.NextPlayer = 1
	ttt.Field = [3]string{"   ", "   ", "   "}
}

// Player bestimmt den nächsten Spieler und setzt diesen Wert in der übergebenen Variable.
func (ttt *TicTacToe) Player() {
	if ttt.NextPlayer == 1 {
		ttt.NextPlayer = 2
	} else {
		// Wenn der letzte Spieler 2 war oder ein unerwarteter Spieler hinterlegt ist, so ist Spieler 1 der Nächste.
		ttt.NextPlayer = 1
	}
}

// Win überprüft, ob ein Spieler gewonnen hat und gibt dies als bool zurück.
func (ttt *TicTacToe) Win(mv Move) bool {
	var symbol byte
	switch ttt.NextPlayer {
	case 1:
		symbol = 'X'
	case 2:
		symbol = 'O'
	}

	if ttt.Field[mv.Row][0] == symbol && ttt.Field[mv.Row][1] == symbol && ttt.Field[mv.Row][2] == symbol {
		ttt.Winner = ttt.NextPlayer
		return true
	} else if ttt.Field[0][mv.Column] == symbol && ttt.Field[1][mv.Column] == symbol && ttt.Field[2][mv.Column] == symbol {
		ttt.Winner = ttt.NextPlayer
		return true
	} else if ttt.Field[0][0] == symbol && ttt.Field[1][1] == symbol && ttt.Field[2][2] == symbol {
		ttt.Winner = ttt.NextPlayer
		return true
	} else if ttt.Field[0][2] == symbol && ttt.Field[1][1] == symbol && ttt.Field[2][0] == symbol {
		ttt.Winner = ttt.NextPlayer
		return true
	} else {
		return false
	}
}

// End überprüft, ob das Spielfeld vollständig belegt ist.
func (ttt *TicTacToe) End() bool {
	return !(strings.Contains(ttt.Field[0], " ") && strings.Contains(ttt.Field[1], " ") && strings.Contains(ttt.Field[2], " "))
}

// Set überträgt den Spielzug auf das Spielfeld.
func (ttt *TicTacToe) Set(mv Move) {
	var symbol string
	switch ttt.NextPlayer {
	case 1:
		symbol = "X"
	case 2:
		symbol = "O"
	}

	switch mv.Column {
	case 0:
		ttt.Field[mv.Row] = symbol + ttt.Field[mv.Row][1:]
	case 1:
		ttt.Field[mv.Row] = ttt.Field[mv.Row][:1] + symbol + ttt.Field[mv.Row][2:]
	case 2:
		ttt.Field[mv.Row] = ttt.Field[mv.Row][:2] + symbol
	}
}
