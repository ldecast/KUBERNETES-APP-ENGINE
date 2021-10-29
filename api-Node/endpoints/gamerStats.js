const express = require('express');

const router = express.Router();

router.get('/', async (req, res) => {
    try {
        const mongo = require("../controllers/mongo");
        const collection = mongo.getCollection();

        /* Obtener los jugadores que han ganado alguna vez */
        collection.distinct(
            "winner",
            {}, // query object
            (function (err, result) {
                if (err) throw err;
                // console.log(result);
                res.status(200).send(result);
            })
        );

    } catch (error) {
        console.log("ERROR: ", error);
        res.sendStatus(500);
    }
});

router.post('/', async (req, res) => {
    try {
        const _player = req.body.player;
        const mongo = require("../controllers/mongo");
        const collection = mongo.getCollection();

        /* Obtener los juegos en que ganó */
        collection.find({ "winner": _player }).toArray(function (err, result) {
            if (err) throw err;
            // console.log(result, _player)
            res.status(200).send(result);
        });
    } catch (error) {
        console.log("ERROR: ", error);
        res.sendStatus(500);
    }
});

module.exports = router;