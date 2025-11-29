// Используем location.host, чтобы работало и на localhost, и на реальном сервере
let socket = new WebSocket("ws://" + location.host + "/ws");

socket.onopen = () => {
    console.log("Successfully connected to Signaling Server");
    // Отправим тестовое сообщение
    socket.send(JSON.stringify({ type: "test", content: "Hello Go!" }));
};

socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
};

socket.onerror = (error) => {
    console.log("Socket Error: ", error);
};

// Слушаем входящие сообщения (пока что это наше же эхо)
socket.onmessage = (event) => {
    console.log("Message from server: ", event.data);
};