require('dotenv').config({ path: __dirname + '/../.env' })
const redis = require('redis');
const { promisifyAll } = require('bluebird');
promisifyAll(redis);

let client;

module.exports = {
    connectDB: () => {
        /* Conexión a Redis */
        const io = require('./sockets').get();
        client = redis.createClient({
            host: process.env.HOST_REDIS,
            port: process.env.PORT_REDIS,
            password: process.env.PASSWORD_REDIS
        });
        client.config('set', 'notify-keyspace-events', 'KEA');
        client.subscribe('__keyevent@0__:set');
        client.on('message', function (channel, key) {
            // console.log("key updated:", key);
            io.emit('redis-report');
        });
        // client.set("key", "value", redis.print);
        // client.get("key", redis.print);
        // console.log("Redis service is connected");
        client.on('error', err => {
            console.log('Redis Error: ' + err);
        });
        return client;
    },
    getClient: () => {
        if (!client)
            throw new Error("Redis Client is not initialized");
        return client;
    },
    close: () => {
        client.quit();
    }
};