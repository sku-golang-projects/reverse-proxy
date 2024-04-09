const express = require("express");

const app = express();

let port = process.env.PORT;

app.get("/", (req, res) => {
    res.send("From " + port + ": Welcome in the node js world.");
})

app.listen(port, () => {
    console.log("listening on port %s.", port);
})