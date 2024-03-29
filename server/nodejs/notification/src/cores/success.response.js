const {
  ReasonPhrases,
  StatusCodes,
} = require("../utils/httpStatusCodes/httpStatusCode");

class SuccessResponse {
  constructor({
    message,
    statusCode = StatusCodes.OK,
    reasonPhrase = ReasonPhrases.OK,
    metadata = {},
  }) {
    this.message = !message ? reasonPhrase : message;
    this.status = statusCode;
    this.metadata = metadata;
  }

  send(res, headers = {}) {
    return res.status(this.status).json(this);
  }
}

class OK extends SuccessResponse {
  constructor({ message, metadata }) {
    super({ message, metadata });
  }
}

class CREATE extends SuccessResponse {
  constructor({
    options = {},
    message,
    statusCode = StatusCodes.CREATED,
    reasonPhrase = ReasonPhrases.CREATED,
    metadata,
  }) {
    super(message, statusCode, reasonPhrase, metadata);
    this.options = options;
  }
}

module.exports = {
  SuccessResponse,
  OK,
  CREATE,
};
