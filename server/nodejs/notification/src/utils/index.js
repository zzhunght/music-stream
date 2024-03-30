const amqplib = require("amqplib");
const { v4: uuid4 } = require("uuid");
const { Types } = require("mongoose");

let amqplibConnection = null;

const convertToObjectIdMongoDB = (id) => new Types.ObjectId(id);

const createChannelMsgQueue = async () => {
  if (amqplibConnection === null) {
    console.log("check connect");
    amqplibConnection = await amqplib.connect(process.env.RABBITMQ_URL);
  }
  return await amqplibConnection.createChannel();
};

const getChannel = async () => {
  try {
    const channel = await createChannelMsgQueue();
    channel.assertExchange(process.env.EXCHANGE_NAME, "direct", false);
    return channel;
  } catch (error) {
    console.log(error);
  }
};

//Publish, Subscribe RabbitMQ
module.exports.publishMessage = async (channel, routing_key, message) => {
  try {
    channel.publish(
      process.env.EXCHANGE_NAME,
      routing_key,
      Buffer.from(JSON.stringify(message))
    );
  } catch (error) {
    console.log(error);
  }
};

const subscribeMessage = async (channel, service) => {
  try {
    const q = await channel.assertQueue(process.env.QUEUE_NAME);
    channel.bindQueue(
      q.queue,
      process.env.EXCHANGE_NAME,
      process.env.NOTI_ROUTING_KEY
    );
    channel.consume(q.queue, (msg) => {
      service.SubscribeEvents(msg.content.toString());
      // console.log("received message: " + msg.content.toString());
      channel.ack(msg);
    });
  } catch (error) {
    console.log(error);
  }
};

//RPC RabbitMQ

module.exports.RPCObserver = async (RPC_QUEUE_NAME, service) => {
  try {
    const channel = await createChannelMsgQueue();
    await channel.assertQueue(RPC_QUEUE_NAME, {
      durable: true,
    });
    channel.perfetch(1);
    channel.consume(
      RPC_QUEUE_NAME,
      async (msg) => {
        if (msg.content) {
          const payload = JSON.parse(msg.content.toString());
          console.log(payload);
          channel.sendToQueue(
            msg.properties.replyTo,
            Buffer.from(JSON.stringify(payload)),
            {
              corelationId: msg.properties.corelationId,
            }
          );
          channel.ack(msg);
        }
      },
      {
        noAck: false,
      }
    );
  } catch (error) {
    console.log(error);
  }
};

const dataRequest = async (RPC_QUEUE_NAME, requestPayload, uuid) => {
  try {
    const channel = await createChannelMsgQueue();
    const q = await channel.assertQueue("", { exclusive: true });
    channel.sendToQueue(
      RPC_QUEUE_NAME,
      Buffer.from(JSON.stringify(requestPayload)),
      {
        replyTo: q.queue,
        corelationId: uuid,
      }
    );

    return new Promise((resolve, reject) => {
      channel.consume(
        q.queue,
        (msg) => {
          if (msg.properties.corelationId === uuid) {
            resolve(msg.content.toString());
          } else {
            reject("Data not found!");
          }
        },
        {
          noAck: true,
        }
      );
    });
  } catch (error) {
    console.log(error);
  }
};

const RPCRequest = async (RPC_QUEUE_NAME, requestPayload) => {
  try {
    const uuid = new uuid4();
    return await dataRequest(RPC_QUEUE_NAME, requestPayload, uuid);
  } catch (error) {
    console.log(error);
  }
};

module.exports = {
  getChannel,
  subscribeMessage,
  RPCRequest,
  convertToObjectIdMongoDB,
};
