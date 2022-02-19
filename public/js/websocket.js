export function websocketSetup(url, onopen, onclose, onmessage, onerror) {
    const socket = new WebSocket(url);
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
        onopen();
    };

    socket.onclose = (event) => {
        console.log("Socket Closed Connection: ", event);
        onclose();
    };

    socket.onmessage = (message) => {
        onmessage(message);
    };

    socket.onerror = (error) => {
        console.log("Socket Error: ", error);
        onerror();
    };

    return socket;
}