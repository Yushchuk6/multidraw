import { websocketSetup } from "./websocket.js"
import { drawPointMap, draw } from "./draw.js"
import { mouseHandler } from "./mouse.js"

const url = location.protocol.replace("http", "ws") + "//" + location.host + "/ws";

const id = uuidv4();
let x = 0, y = 0;
const clientMap = new Map();

const canvas = document.getElementById("canvas");
canvas.width = canvas.clientWidth;
canvas.height = canvas.clientHeight;
const ctx = canvas.getContext("2d");

mouseHandler(canvas,
    () => [x, y] = onMouseUp(websocket, x, y),
    () => { },
    (e) => [x, y] = onMouseMove(ctx, websocket, x, y, e),
    () => [x, y] = onMouseOut(websocket, x, y));

const onMouseUp = (websocket, x, y) => {
    x = 0, y = 0;
    websocket.send(JSON.stringify({ x, y }));
    return [x, y]
}
const onMouseOut = onMouseUp;

const onMouseMove = (ctx, websocket, x, y, e) => {
    draw(ctx, { x, y }, { x: e.offsetX, y: e.offsetY });
    websocket.send(JSON.stringify({ x: e.offsetX, y: e.offsetY }));
    return [e.offsetX, e.offsetY];
}

const websocket = websocketSetup(url,
    () => onOpen(websocket, id),
    () => { },
    (m) => onMessage(ctx, id, clientMap, JSON.parse(m.data)),
    () => { },
    () => { });

const onOpen = (websocket, id) => {
    websocket.send(id);
}

const onMessage = (ctx, id, clientMap, data) => {
    drawPointMap(ctx, id, clientMap, data);
}



