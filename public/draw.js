function drawPointList(ctx, id, clientMap, list) {
  list.forEach(point => {
    if (point.id === id) return;
    if (clientMap.has(point.id)) {
      draw(ctx, clientMap.get(point.id), point)
    }
    clientMap.set(point.id, { x: point.x, y: point.y });
  });
}

function setup(canvas, socket) {
  const ctx = canvas.getContext("2d");
  let isPressed = false;
  let x = null, y = null;

  canvas.addEventListener("mousedown", () => {
    isPressed = true;
  });
  canvas.addEventListener("mouseup", () => {
    isPressed = false;
    x = null;
    y = null;
    socket.send(JSON.stringify({ x: 0, y: 0 }));
  });
  canvas.addEventListener("mousemove", (e) => {
    if (!isPressed) return;
    if (x !== null && y !== null) {
      draw(ctx, { x, y }, { x: e.offsetX, y: e.offsetY });
    }
    socket.send(JSON.stringify({ x: e.offsetX, y: e.offsetY }));
    x = e.offsetX;
    y = e.offsetY;
  });
}

function draw(ctx, oldXY, newXY) {
  ctx.lineWidth = 3;
  ctx.lineCap = 'round';

  if (oldXY.x === 0 && oldXY.y === 0 || newXY.x === 0 && newXY.y === 0) return;
  ctx.beginPath();
  ctx.moveTo(oldXY.x, oldXY.y);
  ctx.lineTo(newXY.x, newXY.y);
  ctx.stroke();
  ctx.closePath();
}


