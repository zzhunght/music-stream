const NotiModel = require("../notification.model");

const findNotificationById = async ({
  type,
  senderId,
  receiverId,
  content,
  read = false,
}) => {
  const noti = await NotiModel.create({
    type,
    senderId,
    receiverId,
    content,
    read,
  });
  return noti;
};

module.exports = {
  findNotificationById,
};
