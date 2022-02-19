export function mouseHandler(element, mouseup, mousedown, mousemove, mouseout) {
    let isPressed = false;

    element.onmousedown = () => {
        isPressed = true;
        mousedown();
    };

    element.onmouseup = () => {
        isPressed = false;
        mouseup();
    };

    element.onmousemove = (e) => {
        if (!isPressed) return;
        mousemove(e);
    };

    element.onmouseout = () => {
        isPressed = false;
        mouseout();
    };
}