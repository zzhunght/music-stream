const dev = {
    app: {
        port: process.env.DEV_APP_PORT,
    },
    db: {
        host: process.env.DEV_DB_HOST || "0.0.0.0",
        port: process.env.DEV_DB_PORT || 27017,
        name: process.env.DEV_DB_NAME || "follow-music-dev",
    },
};

const prod = {
    app: {
        port: process.env.PROD_APP_PORT || 3000,
    },
    db: {
        host: process.env.PROD_DB_HOST || "0.0.0.0",
        port: process.env.PROD_DB_PORT || 27017,
        name: process.env.PROD_DB_NAME || "follow-music-prod",
    },
};

const config = { dev, prod };
const env = process.env.NODE_ENV || "dev";

module.exports = config[env];
