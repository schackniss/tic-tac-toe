# tic-tac-toe


## Klausurersatzleistung Web-Programmierung
- Marc Glocker (4693611)
- Clara Lo Curto (8484335)
- Johannes Schackniß (4175962)

## Enthaltene Dateien
In jedem Verzeichnis ist nochmals eine `README.md`-Datei, zur genaueren Erläuterung des jeweiligen Verzeichnis, enthalten.
- `docs/`: Dokumentation, Beispiele
- `game/`: Tic-Tac-Toe Spiellogik
- `web/`: HTML, CSS, JavaScript usw. zur browserseitigen Ausführung
- `websocket/`: Websocket-Verbindung
- `main.go`: Aufruf der serverseitigen Funktionen
- `go.mod`: Go-Pakete bzw. Abhängigkeiten
- `tic-tac-toe.exe`: kompiliertes Projekt als ausführbare Datei

## Anleitung: Projekt ausführen
1. Verzeichnis entpacken
2. `tic-tac-toe.exe` starten
3. Bei Firewall-Warnung `Zugriff zulassen` klicken
4. Im Terminal-Fenster erscheint folgende Ausgabe:
   ```
   Starting webserver...
   http://localhost:8080
   ```
5. URL [http://localhost:8080](http://localhost:8080) im Web-Browser öffnen.
6. Die Webanwendung kann auch von anderen Clients im Netzwerk aufgerufen werden. Hierfür muss `localhost` durch die IP-Adresse des Gerätes, auf dem `tic-tac-toe.exe` ausgeführt wird, ersetzt werden.
7. Spaß beim Spielen nicht vergessen!

## Lastenheft
- Umsetzung der Spiellogik
- Darstellung im Web-Browser --> Spiel ermöglichen ohne Neuladen der Seite 
- Beliebig viele Spieler und ein Spielfeld 
- Jeder verbundene Client kann sowohl für Spieler 1 als auch für Spieler 2 Züge ausführen. 
- Anzeige des Spielfelds, des aktuellen Spielers und gegebenenfalls des Spielergebnisses

## Pflichtenheft
Das Projekt wird in drei Teilaufgaben unterteilt. Es wird in Server (Backend) und Browser (Frontend) strukturiert. Browserseitig findet eine getrennte Entwicklung der visuellen Darstellung und der Logik statt. 

### Backend
- Realisierung mit GO
- Tic Tac Toe Spiellogik implementieren
- Implementierung eines WebSockets für die bidirektionale Verbindung zwischen Front- und Backend
- Frontend hosten (Webserver)
- API-Anfragen auswerten und basierend auf der Logik eine API-Antwort senden (Kommunikation zwischen Front und Backend)

### Frontend - Logik
- Realisierung mit JavaScript
- WebSocket Verbindung zum Backend implementieren 
- API-Anfragen senden (bei Spielzug)   
- API-Antworten auswerten: 
  - Darstellung des aktuellen Spielstands (Pop-up bei Gewinn/Unentschieden)

### Frontend - Visuell 
- Realisierung mit HTML/CSS 
- Überschrift
- Anzeige des Spielfeldes für eine Userinteraktionsmöglichkeit (Buttons)
- Pop-up Fenster des Spielergebnisse

## Screenshots und Dokumentation
Screenshots und Dokumentation sind im Verzeichnis `docs/` zu finden.