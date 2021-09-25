const EVENT_ENTER_LOBBY = "enter_lobby";
const EVENT_NEW_MATCH = "new_match";
const EVENT_ATTACK = "attack";

var rawRender = (game) => {
    let boardWrapper = document.getElementById("board");
    let board = document.createElement("table");
    board.style.width  = '500px';
    board.style.border = '2px solid navy';

    boardSettings = game["board"]["settings"];
    for (r = 0; r < boardSettings["rows"]; r++){
        var tr = board.insertRow();
        for(c = 0; c < boardSettings["columns"]; c++) {
            var attackBtn = document.createElement("button");
            attackBtn.appendChild(document.createTextNode("Attack " + r + ":" + c))
            attackBtn.onclick = window.game.attack(r,c);

            var td = tr.insertCell();
            td.style.border = '1px solid #ccc';
            td.style.color = '#ccc';
            td.appendChild(attackBtn);
        }
    }

    boardWrapper.innerHTML = '';
    boardWrapper.appendChild(board);    
}


var debugRender = (game) => {
    let debugWrapper = document.getElementById("debug");
    let debug = document.createElement("table");
    debug.style.width  = '500px';
    debug.style.border = '1px solid black';

    let debugAddRow = (name, value) => {
        var tr = debug.insertRow();
        [name, value].forEach((text) => tr.insertCell().appendChild(document.createTextNode(text))); 
    }

    debugAddRow("Property", "Value"); 
    if (game) {
        debugAddRow("Game ID", game["id"]); 
        debugAddRow("Board > Columns", game["board"]["settings"]["columns"]); 
        debugAddRow("Board > Rows", game["board"]["settings"]["rows"]); 
    }

    debugWrapper.innerHTML = '';
    debugWrapper.appendChild(debug);    
}

class BattleshipClient {

    socket;
    game;
    renders;

    constructor(renders) {
        this.renders = Array.isArray(renders)? renders: [renders];
        this.loadWebsocket();
    }

    loadWebsocket() {
        console.log("Attempting Connection...");
        this.socket = new WebSocket("ws://localhost:8080/game/ws");

        this.socket.onopen = () => {
            console.log("Successfully Connected");
        };
        
		this.socket.onmessage = this.proxyEvent();

        this.socket.onclose = event => {
            console.log("Socket Closed Connection: ", event);
        };

        this.socket.onerror = error => {
            console.log("Socket Error: ", error);
        };
    }

    enterLobby() {
        var playerName = prompt("What is your name?");
        this.send(JSON.stringify({
            "type" : EVENT_ENTER_LOBBY,
            "message": JSON.stringify({
                "name" : playerName
            })
        }))
    }

    create() {
        console.log("Creating game")
        this.send(JSON.stringify({
            "type" : EVENT_NEW_MATCH,
            "message": ""
        }))
    }    
    
    proxyEvent() {
        const instance = this;
        return (event) => {
            console.log("Received event", event.data) // TODO add type to returns;
            instance.game = JSON.parse(event.data);
            instance.renders.forEach(render => render(instance.game));
        }
    }
    
    attack(row, column) {
        return () => {
            console.log("attacking ", row, column);
            this.send(JSON.stringify({
                "type" : EVENT_ATTACK,
                "message": JSON.stringify({ "column" : column, "row" : row })
            }))
        }
    }
    
    send(event) {
        this.socket.send(event);
    }

}