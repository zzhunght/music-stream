const {
  findNotificationById,
} = require("../models/repositories/notification.repo");
const { RPCRequest } = require("../utils");

class NotificaionService {
  async createNotification(data) {
    try {
      // get list follower of artists
      const artists = await RPCRequest(process.env.ARTISTS_RPC, {
        event: "GET_ARTISTS",
        data: {
          shopId: data.shopId,
        },
      });

      console.log(artists);
      // console.log("Data notiservice: ", data);
      let content;
      if (data.type === "SHOP-001") {
        content = `Artists ${artists.name} has just released a new song: ${data.content}`;
      }

      const notiData = {
        type: data.type,
        senderId: data.senderId,
        receiverId: data.receiverId,
        content: content,
        read: false,
      };

      //create notification
      const notification = await findNotificationById(notiData);

      return notification;
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
