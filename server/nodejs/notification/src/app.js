const express = require("express");
require("dotenv").config();
const notification = require("./api/notification");
const { getChannel } = require("./utils");
const app = express();

app.use(express.json());

const appServer = async () => {
  //init database
  // require("./dbs/init.mongodb");
  require('./configs/config.firebase').initFirebaseApp()
  // const { checkOverLoad } = require("./helpers/check.connect");
  // checkOverLoad();

  //create channel rabbitmq
  const channel = await getChannel();
  notification(app, channel);
};

appServer();

// handling errors
app.use((req, res, next) => {
  const error = new Error("Not Found");
  error.status = 404;
  next(error);
});

app.use((error, req, res, next) => {
  const statusCode = error.status || 500;
  return res.status(statusCode).json({
    status: "error",
    code: statusCode,
    stack: error.stack,
    message: error.message || "Internal Server Error",
  });
});

module.exports = app;
