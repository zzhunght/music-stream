const { Schema, model, Types } = require("mongoose");

const DOCUMENT_NAME = "Follow";
const COLLECTION_NAME = "Follows";

const FollowSchema = new Schema(
    {
        account_id: { type: Number, required: true },
        artist_id: { type: Number, required: true },
    },
    {
        timestamps: {
            createdAt: "created_at",
            updatedAt: "updated_at",
        },
        collection: COLLECTION_NAME,
    }
);

module.exports = model(DOCUMENT_NAME, FollowSchema);
