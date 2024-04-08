const { Schema, model, Types } = require("mongoose");

const DOCUMENT_NAME = "Playlist";
const COLLECTION_NAME = "Playlists";

const PlaylistSchema = new Schema(
    {
        playlist_name: { type: String, required: true },
        playlist_account_id: { type: Number, required: true },
        playlist_desc: { type: String, required: true },
        playlist_songs: [
            {
                song_id: { type: Number, required: true },
            },
        ],
    },
    {
        timestamps: {
            createdAt: "created_at",
            updatedAt: "updated_at",
        },
        collection: COLLECTION_NAME,
    }
);

module.exports = model(DOCUMENT_NAME, PlaylistSchema);
