const express = require('express');

const router = express.Router();

router.get('/', async (req, res) => {
    try {
        const mongo = require("../controllers/mongo");
        const collection = mongo.getCollection();
        collection.find({}).toArray(function (err, result) {
            if (err) throw err;
            // console.log(result);
            res.status(200).send(result);
        });
    } catch (error) {
        console.log("ERROR: ", error);
        res.sendStatus(500);
    }
})

module.exports = router;