const PlaylistModel = require("../playlist.model");

const createPlaylist = async ({ name, account_id, song_id, desc }) => {
    const playlist = await PlaylistModel.create({
        playlist_name: name,
        playlist_account_id: account_id,
        playlist_songs: [
            {
                song_id: song_id,
            },
        ],
        playlist_desc: desc,
    });
    return playlist;
};

const updateSongPlaylist = async ({ playlist_id, name, song_id }) => {
    const update = await PlaylistModel.findByIdAndUpdate(playlist_id, {
        song_id: song_id,
        playlist_name: name,
    });

    return update;
};

const getPlaylistByUserId = async (userId) => {
    const playlistsByUserId = await PlaylistModel.findAll({
        where: {
            playlist_account_id: userId,
        },
    });

    return playlistsByUserId;
};

const deletePlaylistByUserId = async (userId) => {
    const del = await PlaylistModel.deleteOne({
        where: {
            playlist_account_id: userId,
        },
    });

    return del;
};

module.exports = {
    createPlaylist,
    updateSongPlaylist,
    getPlaylistByUserId,
    deletePlaylistByUserId,
};
