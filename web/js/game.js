const EVENT_NEW_GAME = "new_game";

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
    }

    debugWrapper.innerHTML = '';
    debugWrapper.appendChild(debug);    
}

class BattleshipClient {

    socket;
    game;
    render;

    constructor(render) {
        this.render = render;
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

    create() {
        console.log("Creating game")
        this.send(JSON.stringify({
            "type" : EVENT_NEW_GAME,
            "message": ""
        }))
    }    
    
    proxyEvent() {
        const instance = this;
        return (event) => {
            console.log("Received event", event.data) // TODO add type to returns;
            instance.game = JSON.parse(event.data);
            instance.render(instance.game);
        }
    }
    
    send(event) {
        this.socket.send(event);
    }
}