const express = require("express");
require("dotenv").config();
const notification = require("./api/notification");
const { getChannel } = require("./utils");
const app = express();

app.use(express.json());

const appServer = async () => {
  //init database
  require("./dbs/init.mongodb");
  // const { checkOverLoad } = require("./helpers/check.connect");
  // checkOverLoad();

  //create channel rabbitmq
  const channel = await getChannel();

  notification(app, channel);
};

appServer();

module.exports = app;
