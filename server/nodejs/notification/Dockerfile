FROM node:latest

WORKDIR /app

COPY . .
COPY .env .env
RUN npm install

CMD [ "node", "server.js" ]