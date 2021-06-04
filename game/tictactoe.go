package game

type TicTacToe struct {
	Finished   bool      `json:"finished"`
	Winner     int       `json:"winner"`
	NextPlayer int       `json:"nextPlayer"`
	Field      [3]string `json:"field"`
}

type Move struct {
	Player int  `json:"player"`
	Column int  `json:"col"`
	Row    int  `json:"row"`
	Reset  bool `json:"reset"`
}
