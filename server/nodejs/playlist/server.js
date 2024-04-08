const app = require("./src/app");

const PORT = 3059;

const server = app.listen(PORT, () => {
    console.log("Notificaion Service listening on port: " + PORT);
});

process.on("SIGINT", () => {
    server.close(() => console.log("Exit Notification server Express!"));
});
