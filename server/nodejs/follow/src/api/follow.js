const { SuccessResponse } = require("../cores/success.response");
const FollowService = require("../services/follow.service");
const { subscribeMessage } = require("../utils");

module.exports = (app, channel) => {
    const service = new FollowService();

    subscribeMessage(channel, service);

    // app.get("/create", async (req, res, next) => {
    //     new SuccessResponse({
    //         message: "Get Follow success",
    //         metadata: await service.createNotification(req.body),
    //     }).send(res);
    // });
};
