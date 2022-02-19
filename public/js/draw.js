export function drawPointMap(ctx, id, clientMap, list) {
  list.forEach(point => {
    if (point.id === id) return;
    if (clientMap.has(point.id)) {
      draw(ctx, clientMap.get(point.id), point)
    }
    clientMap.set(point.id, { x: point.x, y: point.y });
  });
}

export function draw(ctx, oldXY, newXY) {
  if (oldXY.x === 0 && oldXY.y === 0 || newXY.x === 0 && newXY.y === 0) return;

  ctx.lineWidth = 2;
  ctx.lineCap = 'round';

  ctx.beginPath();
  ctx.moveTo(oldXY.x, oldXY.y);
  ctx.lineTo(newXY.x, newXY.y);
  ctx.stroke();
  ctx.closePath();
}


