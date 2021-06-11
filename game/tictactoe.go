// Das Paket 'game' beinhaltet die Tic Tac Toe Spiellogik.
package game

import "strings"

// TicTacToe repräsentiert den Spielstand und ist zugleich das API-Antwort-Format.
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

// Play stellt den Spiel-Ablauf dar.
func (ttt *TicTacToe) Play(mv Move) {
	// Prüfen, ob der Zug erlaubt ist (noch kein X oder O auf Feld).
	if mv.allowed(*ttt) {
		ttt.set(mv)
		// Prüfen, ob es einen Gewinner gibt oder ob alle Felder belegt sind.
		if ttt.win(mv) || ttt.end() {
			ttt.Finished = true
		}
		// Nächsten Spieler setzen.
		ttt.player()
	}
}

// player bestimmt den nächsten Spieler und setzt diesen Wert in der übergebenen Variable.
func (ttt *TicTacToe) player() {
	if ttt.NextPlayer == 1 {
		ttt.NextPlayer = 2
	} else {
		// Wenn der letzte Spieler 2 war oder ein unerwarteter Spieler hinterlegt ist, so ist Spieler 1 der Nächste.
		ttt.NextPlayer = 1
	}
}

// win überprüft, ob ein Spieler gewonnen hat und gibt dies als bool zurück.
func (ttt *TicTacToe) win(mv Move) bool {
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

// end überprüft, ob das Spielfeld vollständig belegt ist.
func (ttt *TicTacToe) end() bool {
	return !(strings.Contains(ttt.Field[0], " ") && strings.Contains(ttt.Field[1], " ") && strings.Contains(ttt.Field[2], " "))
}

// set überträgt den Spielzug auf das Spielfeld.
func (ttt *TicTacToe) set(mv Move) {
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
