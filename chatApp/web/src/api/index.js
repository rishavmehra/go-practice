var socket = new WebSocket("ws://localhost:8080/ws")

let connect =() => {
    console.log("creating connnection...");

    socket.onopen = () =>{
        console.log("Successfully connected!");
    }

    socket.onmessage = msg =>{
        console.log(msg);
    }

    socket.onclose = event =>{
        console.log("Connection Closed: ", event);
        
    }

    socket.onerror = error =>{
        console.log("Socker Error: ", error);
    };
}

let sentMsg = msg =>{
    console.log("sending msg: ", msg);
    socket.send(msg)
}

export{connect, sentMsg}

