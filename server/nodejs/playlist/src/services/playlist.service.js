const {
    createPlaylist,
    getPlaylistByUserId,
    deletePlaylistByUserId,
} = require("../models/repositories/playlist.repo");
const { RPCRequest } = require("../utils");

class NotificaionService {
    async createPlaylist(data) {
        try {
            const playlist = await createPlaylist(data);
            return playlist;
        } catch (error) {
            throw error;
        }
    }

    async updateSongPlaylist(data) {
        try {
            const playlist = await updateSongPlaylist(data);
            return playlist;
        } catch (error) {
            throw error;
        }
    }

    async getPlaylistByUserId(userId) {
        try {
            const playlists = await getPlaylistByUserId(userId);
            return playlists;
        } catch (error) {
            throw error;
        }
    }

    async deletePlaylistByUserId(userId) {
        try {
            const deletePlaylist = await deletePlaylistByUserId(userId);
            return deletePlaylist;
        } catch (error) {
            throw error;
        }
    }

    async SubscribeEvents(payload) {
        payload = JSON.parse(payload);
        const { event, data } = payload;

        switch (event) {
            case "CREATE_PLAYLIST":
                this.createPlaylist(data);
                break;
            case "GET_PLAYLIST":
                this.getPlaylistByUserId(data);
                break;
            case "DELETE_PLAYLIST":
                this.deletePlaylistByUserId(data);
                break;
            default:
                break;
        }
    }
}

module.exports = NotificaionService;
