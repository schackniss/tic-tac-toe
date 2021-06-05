// Das Paket 'game' beinhaltet die Tic Tac Toe Spiellogik.
package game

// TicTacToe repräsentiert das Spiel und ist zugleich das API-Antwort-Format.
type TicTacToe struct {
	Finished   bool      `json:"finished"`
	Winner     int       `json:"winner"`
	NextPlayer int       `json:"nextPlayer"`
	Field      [3]string `json:"field"`
}

// Move repräsentiert einen Spielzug und ist zugleich das API-Anfrage-Format.
type Move struct {
	Player int  `json:"player"`
	Column int  `json:"col"`
	Row    int  `json:"row"`
	Reset  bool `json:"reset"`
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
	}
	// Wenn der letzte Spieler 2 war oder ein unerwarteter Spieler hinterlegt ist, so ist Spieler 1 der Nächste.
	ttt.NextPlayer = 1
}

// Allowed prüft, ob der Zug zulässig ist bzw. ob das Feld bereits belegt ist und gibt einen Wahrheitswert (bool) zurück.
func (mv *Move) Allowed(ttt TicTacToe) bool {
	return ttt.Field[mv.Row][mv.Column] == ' '
}

// Symbol gibt das Symbol, dass auf dem Spielfeld gesetzt werden muss zurück.
func (mv *Move) Symbol() string {
	if mv.Player == 1 {
		return "X"
	}
	return "O"
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
		ttt.Finished = true
		return true
	} else if ttt.Field[0][mv.Column] == symbol && ttt.Field[1][mv.Column] == symbol && ttt.Field[2][mv.Column] == symbol {
		ttt.Winner = ttt.NextPlayer
		ttt.Finished = true
		return true
	} else if ttt.Field[0][0] == symbol && ttt.Field[1][1] == symbol && ttt.Field[2][2] == symbol {
		ttt.Winner = ttt.NextPlayer
		ttt.Finished = true
		return true
	} else if ttt.Field[0][2] == symbol && ttt.Field[1][1] == symbol && ttt.Field[2][0] == symbol {
		ttt.Winner = ttt.NextPlayer
		ttt.Finished = true
		return true
	} else {
		return false
	}
}
