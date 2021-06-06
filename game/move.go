package game

// Move repräsentiert einen Spielzug und ist zugleich das API-Anfrage-Format.
type Move struct {
	Player int  `json:"player"`
	Column int  `json:"col"`
	Row    int  `json:"row"`
	Reset  bool `json:"reset"`
}

// Allowed prüft, ob der Zug zulässig ist bzw. ob das Feld bereits belegt ist und gibt einen Wahrheitswert (bool) zurück.
func (mv *Move) Allowed(ttt TicTacToe) bool {
	return ttt.Field[mv.Row][mv.Column] == ' ' && mv.Player == ttt.NextPlayer
}
