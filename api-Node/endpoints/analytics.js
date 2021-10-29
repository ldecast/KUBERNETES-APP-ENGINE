const express = require('express');

const router = express.Router();

router.get('/', async (req, res) => {
    try {
        let data = {
            top3_games: [],
            inserts_workers: [],
            isLoaded: true
        }
        const mongo = require("../controllers/mongo");
        const collection = mongo.getCollection();

        /* Obtener el top 3 de jugadores */
        collection.aggregate(
            [{ $group: { _id: "$gamename", "count": { $sum: 1 } } },
            { $sort: { "count": -1 } },
            { $limit: 3 }]
        ).toArray(function (err, result) {
            if (err) throw err;
            data.top3_games = result;

            /* Obtener la cantidad de inserciones de cada worker */
            collection.aggregate(
                [{ $group: { _id: "$worker", "count": { $sum: 1 } } }]
            ).toArray(function (err, result) {
                if (err) throw err;
                data.inserts_workers = result;

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