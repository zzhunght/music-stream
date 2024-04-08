const { SuccessResponse } = require("../cores/success.response");
const PlaylistService = require("../services/playlist.service");
const { subscribeMessage } = require("../utils");

module.exports = (app, channel) => {
    const service = new PlaylistService();

    subscribeMessage(channel, service);

    app.post("/create", async (req, res, next) => {
        new SuccessResponse({
            message: "Get notification success",
            metadata: await service.createNotification(req.body),
        }).send(res);
    });
};
