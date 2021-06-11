var conn; // Enthält eine Referenz zum Websocket
var msg; // Enthält die API-Request
var res; // Enthält die API-Response
var playernumber = 1; // Enthält den aktuellen Spieler

// Verwaltet Eingaben der Clients und bereitet die Daten zum Senden für den API-Request vor.
function handleClient(element)
{
    changeField(); // muss wieder entfernt werden!
    row = element.name[0];
    column = element.name[1];
    sendAPIRequest(row, column, false);
}

// Verfasst eine Nachricht im JSON-Format und sendet einen API-Request mit der Nachricht.
function sendAPIRequest(row, column, reset)
{
    

    if (!conn) {
        window.alert("KEINE VERBINDUNG!");
    }
    else
    {
        msg = {
            player: playernumber,
            col: column,
            row: row,
            reset: reset,
        };
        conn.send(JSON.stringify(msg));
        //window.alert("Anfrage: " + reset);
    }


    
}

// Verwaltung der API-Antwort
function handleAPIResponse()
{
    changeField();
    evaluateGame();
    changePlayer();
}

// Ändert den Inhalt des Spielfelds abhängig von der API-Antwort.
function changeField()
{    
    buttons = document.getElementById("spielfeld").getElementsByTagName("input"); // Buttons ist ein Array, der eine Referenz zu allen Buttons des Spielfelds enthält
    for(i=0; i<buttons.length; i++)
    {
        buttons[i].value = res.field[Math.floor(i/3)][Math.floor(i%3)];
    }
}

// Auswertung des Spiels. Wenn es fertig ist, soll eine Anzeige aufpoppen und ..........
function evaluateGame()
{
    if(res.finished == true)
    {  
        var endmsg;
        if(res.winner == 0) endmsg = "Es gibt keinen Gewinner (Unentschieden).";
        else endmsg = "Der Gewinner ist Spieler " + res.winner; 
        window.alert("Das Spiel ist fertig. " + endmsg)

        // Hier noch entscheiden, was am Ende passiert (Alles abbrechen??)
    }
}

// Passt den aktuellen Spielernamen abhängig von der API-Antwort an.
function changePlayer()
{
    playernumber = res.nextPlayer;
    document.getElementById("spielernummer").innerHTML = playernumber;
}

// Wird beim Starten des Servers ausgeführt
window.onload = function () 
{
    // Brauchen wir das??
    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    // Sorgt für eine Verbindung mit dem Websocket
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (event) { // Fängt API-Responses ab
            res = JSON.parse(event.data);
            handleAPIResponse();
            
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};