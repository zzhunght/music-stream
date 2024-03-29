const {
  findNotificationById,
} = require("../models/repositories/notification.repo");
const { RPCRequest } = require("../utils");

class NotificaionService {
  async createNotification(data) {
    try {
      // get list follower of artists
      const artists = await RPCRequest("SHOP_RPC", {
        event: "GET_SHOP",
        data: {
          shopId: data.shopId,
        },
      });

      console.log(artists);
      // console.log("Data notiservice: ", data);

      //create notification
      // const notification = await findNotificationById(data);

      return data;
    } catch (error) {
      throw error;
    }
  }

  async SubscribeEvents(payload) {
    payload = JSON.parse(payload);
    const { event, data } = payload;

    switch (event) {
      case "CREATE_NOTIFICATION":
        this.createNotification(data);
        break;
      default:
        break;
    }
  }
}

module.exports = NotificaionService;
