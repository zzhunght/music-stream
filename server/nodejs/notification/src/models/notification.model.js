const { Schema, model, Types } = require("mongoose");

const DOCUMENT_NAME = "Notification";
const COLLECTION_NAME = "Notification";

const NotificaionSchema = new Schema(
  {
    type: { type: String, required: true },
    senderId: { type: Types.ObjectId, required: true },
    receiverId: { type: Types.ObjectId, required: true },
    content: { type: String, required: true },
    read: { type: Boolean, default: false },
  },
  {
    timestamps: {
      createdAt: "created_at",
      updatedAt: "updated_at",
    },
    collection: COLLECTION_NAME,
  }
);

module.exports = model(DOCUMENT_NAME, NotificaionSchema);
