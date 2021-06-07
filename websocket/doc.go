// Das Paket 'websocket' dient der Implementierung von Websockets für das Spiel Tic Tac Toe.
// Die Implementierung basiert auf dem Repository https://github.com/gorilla/websocket.
//
// API-Anfragen
//
// API-Anfragen sendet immer ein Client/Browser/Frontend.
// Diese benötigen folgendes Format und stellen einen Spielzug dar:
//
//   {
//       "player": 1,
//       "col": 2,
//       "row": 0,
//       "reset": false
//   }
//
// API-Antworten
//
// API-Antworten sendet immer der Server bzw. das Backend.
// Diese sind im folgenden Format und stellen den Spielstand dar:
//	 {
//	     "finished": false,
//	     "winner": 0,
//	     "nextPlayer": 1,
//	     "field": ["X  ", "O  ", "   "]
//	 }
package websocket
