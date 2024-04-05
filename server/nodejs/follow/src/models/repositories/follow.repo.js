const FollowModel = require("../follow.model");

const createFollow = async ({ account_id, artist_id }) => {
    const follow = await FollowModel.create({
        account_id,
        artist_id,
    });
    return follow;
};

const deleteFollow = async ({ account_id, artist_id }) => {
    const follow = await FollowModel.deleteOne({
        account_id,
        artist_id,
    });
    return follow;
};

module.exports = {
    createFollow,
    deleteFollow,
};
