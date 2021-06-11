# Paket: websocket
Das Paket gehört zur serverseitigen Implementierung, dem sogenannten Backend.
Die Umsetzung ist mit Hilfe des Repositories [`github.com/gorilla/websocket`](https://github.com/gorilla/websocket) realisiert.

## Enthaltene Dateien
- `client.go`:
  - Implementierung der Struktur `Client` (Verbinder zwischen der Websocket-Verbindung und dem Hub)
  - Implementierung aller Methoden mit Empfängerzeiger vom Typ `Client` und der Funktion `Serve`
- `doc.go`:
  - allgemeine Dokumentation zur Verwendung der API (formatierte Version in `docs/websocket-doc.pdf`)
- `handler.go`:
  - Implementierung der API
- `hub.go`:
  - Implementierung der Struktur `Hub` (verwaltet die Menge der aktiven Clients)
  - Implementierung aller Methoden mit Empfängerzeiger vom Typ `Hub`

## Code-Dokumentation
- im Code als Kommentare
- in der PDF `docs/websocket-doc.pdf`
- im UML-Diagramm `docs/uml-diagram.png`