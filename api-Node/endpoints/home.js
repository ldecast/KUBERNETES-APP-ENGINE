const express = require('express');

const router = express.Router();

router.get('/', async (req, res) => {
    try {
        let data = {
            last_games: [],
            top10_players: []
        }
        const mongo = require("../controllers/mongo");
        const collection = mongo.getCollection();

        /* Obtener los últimos 10 juegos */
        await collection.find({}).sort({ _id: -1 }).limit(10).toArray(function (err, result) {
            if (err) throw err;
            data.last_games = result;

            /* Obtener el top 10 de jugadores */
            collection.aggregate(
                [{ $group: { _id: "$winner", "wins": { $sum: 1 } } },
                { $sort: { "wins": -1 } },
                { $limit: 10 }]
            ).toArray(function (err, result) {
                if (err) throw err;
                data.top10_players = result;

                // console.log(data)
                res.status(200).send(data);
            });
        });
    } catch (error) {
        console.log("ERROR: ", error);
        res.sendStatus(500);
    }
})

module.exports = router;