const { createFollow } = require("../models/repositories/follow.repo");

class FollowService {
    async createFollow(data) {
        try {
            //create notification
            const follow = await createFollow(data);

            return follow;
        } catch (error) {
            throw error;
        }
    }

    async deleteFollow(data) {
        try {
            //create notification
            const follow = await deleteFollow(data);

            return follow;
        } catch (error) {
            throw error;
        }
    }

    async SubscribeEvents(payload) {
        payload = JSON.parse(payload);
        const { event, data } = payload;

        switch (event) {
            case "CREATE_FOLLOW":
                this.createFollow(data);
                break;
            case "DELETE_FOLLOW":
                this.deleteFollow(data);
                break;
            default:
                break;
        }
    }
}

module.exports = FollowService;
