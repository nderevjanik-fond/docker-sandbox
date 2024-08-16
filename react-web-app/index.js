const express = require("express");
const app = express();
const port = process.env.PORT || 3000;

app.get("/", (req, res) => {
  const environment = process.env["NODE_ENV"].toUpperCase();
  res.send(`(${environment}) Hello from Dockerized Node.js app!`);
});

app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});
