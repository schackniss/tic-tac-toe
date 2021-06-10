var playernumber = 1; // Enthält den aktuellen Spieler

// Diese Funktion wird ausgeführt, wenn ein Button auf dem Spielfeld gedrückt wird.
// Abhängig vom aktuellen Spieler wird ein 'X' oder ein 'O' in das Feld geschrieben wird und das Feld wird für weitere Eingaben gesperrt.
function changeField(element){
    if(playernumber == 1)
    {
        element.value = "O";
        element.disabled = true;
        playernumber = 2;
    } 
    else if(playernumber == 2)
    {
        element.value = "X";
        element.disabled = true;
        playernumber = 1;
    }
    document.getElementById("spielernummer").innerHTML = playernumber;
}

// Diese Funktion wird ausgeführt, wenn der Reset-Button gedrückt wird. Sie leert alle Felder des Spielsfelds und entsperrt diese ggf. wieder.
function resetField()
{
    buttons = document.getElementById("spielfeld").getElementsByTagName("input"); // Buttons ist ein Array, der eine Referenz zu allen Buttons des Spielfelds enthält

    for(i=0; i<buttons.length; i++)
    {
        buttons[i].value = " ";
        buttons[i].disabled = false;
    }
}

window.onload = function () {
    var conn;
    var msg = document.getElementById("msg"); // Request
    var log = document.getElementById("log"); //Answer

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    // Schickt die Nachricht
    document.getElementById("form").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        conn.send(msg.value);
        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};