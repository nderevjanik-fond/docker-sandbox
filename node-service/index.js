const express = require("express");
const app = express();
const port = process.env.PORT || 3000;

app.get("/hello", (req, res) => {
  res.send({ message: "hello world" });
});

app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});
