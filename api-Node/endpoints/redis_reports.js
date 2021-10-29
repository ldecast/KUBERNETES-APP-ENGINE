require('dotenv').config({ path: __dirname + '/../.env' })
const express = require('express');
const router = express.Router();
const redis = require('redis');
const { promisifyAll } = require('bluebird');

promisifyAll(redis);

router.get('/', async (req, res) => {
    try {
        const client = redis.createClient({
            host: process.env.HOST_REDIS,
            port: process.env.PORT_REDIS,
            password: process.env.PASSWORD_REDIS
        });
        client.on('error', err => {
            console.log('Redis Error: ' + err);
        });
        /*  
        data.last_request = await client.getAsync('last_request');
        data.last_gameid = await client.getAsync('last_gameid');
        data.last_gamename = await client.getAsync('last_gamename');
        data.last_winner = await client.getAsync('last_winner');
        data.last_players = await client.getAsync('last_players');
        data.last_worker = await client.getAsync('last_worker');
        */
        let data = {
            last_request: await client.getAsync('last_request'),
            last_gameid: await client.getAsync('last_gameid'),
            last_gamename: await client.getAsync('last_gamename'),
            last_winner: await client.getAsync('last_winner'),
            last_players: await client.getAsync('last_players'),
            last_worker: await client.getAsync('last_worker'),
            isLoaded: true
        }
        client.quit();
        res.status(200).send(data);

    } catch (error) {
        console.log("ERROR: ", error);
        res.sendStatus(500);
    }
})

module.exports = router;