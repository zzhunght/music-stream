const NotificationService = require("../services/notification.service");
const { subscribeMessage } = require("../utils");

module.exports = (app, channel) => {
  const service = new NotificationService();

  subscribeMessage(channel, service);

  app.get("/create", async (req, res, next) => {
    try {
      const newNoti = await service.createNotification(req.body);

      return res.status(200).json({
        msg: "Created notification successfully",
        metadata: newNoti,
      });
    } catch (error) {
      return res.status(500).json({
        msg: "Failed to create notification",
        error,
      });
    }
  });
};
