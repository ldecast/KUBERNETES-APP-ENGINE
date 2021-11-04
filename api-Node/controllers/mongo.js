require('dotenv').config({ path: __dirname + '/../.env' })
const { MongoClient } = require('mongodb');
const client = new MongoClient(process.env.URI_MONGO, { useNewUrlParser: true, useUnifiedTopology: true });
let collection;

module.exports = {
    connectDB: () => {
        const io = require('./sockets').get();
        client.connect(err => {
            if (!err) {
                console.log("Mongo Database is connected");
                collection = client.db(process.env.DB_MONGO).collection(process.env.COLLECTION_MONGO);
                setInterval(() => io.emit('log-inserted'), 1000 * 2);
                /* changeStream = collection.watch();
                changeStream.on("change", next => {
                    // console.log("received a change to the collection: \t", next);
                    io.emit('log-inserted');
                }); */
            } else {
                console.log("Mongo Database is not connected: " + err);
                throw err;
            }
        });
    },
    getCollection: () => {
        if (!collection)
            throw new Error("Mongo Collection is not initialized");
        return collection;
    },
    close: () => {
        client.close();
    }
};
