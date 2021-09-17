const EVENT_NEW_GAME = "new_game";


class BattleshipClient {

    socket;
    game;

    constructor() {
        console.log("Attempting Connection...");
        this.socket = new WebSocket("ws://localhost:8080/game/ws");

        this.socket.onopen = () => {
            console.log("Successfully Connected");
            this.create()
        };
        
		this.socket.onmessage = this.proxyEvent;

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
    
    proxyEvent(event) {
        console.log("Received event", event.data)
        this.game = JSON.parse(event.data);
        console.log(this.game);
    }
    
    send(event) {
        this.socket.send(event);
    }
}