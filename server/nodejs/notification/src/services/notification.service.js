const {
    findNotificationById,
} = require("../models/repositories/notification.repo");
const { RPCRequest } = require("../utils");

const { messaging } = require('firebase-admin')
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

            case 'CREATE_NEW_SONG':

                const message = {
                    notification: {
                        title: 'Bài hát ' + payload.data.song_name + " của " + payload.data.artist.name + " vừa được ra mắt",
                        body: payload.data.song_name
                    },
                    android: {
                        notification: {
                            imageUrl: payload.data.thumbnail
                        }
                    },
                    topic: "public"
                };
                console.log("message :", message)
                this.sendNotification(message);
                break;
            default:
                break;
        }
    }

    async sendNotification(message) {
        messaging().send(message)
            .then((response) => {
                console.log('Successfully sent message:', response);
            })
            .catch((error) => {
                console.error('Error sending message:', error);
            });
    }
}

module.exports = NotificaionService;
