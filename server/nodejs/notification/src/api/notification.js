const { SuccessResponse } = require("../cores/success.response");
const NotificationService = require("../services/notification.service");
const { subscribeMessage } = require("../utils");

module.exports = (app, channel) => {
  const service = new NotificationService();

  subscribeMessage(channel, service);

  app.get("/create", async (req, res, next) => {
    new SuccessResponse({
      message: "Get notification success",
      metadata: await service.createNotification(req.body),
    }).send(res);
  });
};
