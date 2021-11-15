require('dotenv').config({ path: __dirname + '/.env' })
const express = require('express');
const cors = require('cors');
const http = require('http');

const Home = require('./endpoints/home');
const Transactions = require('./endpoints/transactions');
const GamerStats = require('./endpoints/gamerStats');
const Analytics = require('./endpoints/analytics');
const RedisReports = require('./endpoints/redis_reports');

/* init */
const PORT = process.env.PORT || 8080;
const app = express();
app.use(express.static(__dirname + '/public/build'));

const server = http.createServer(app);

/* Middlewares */
app.use(express.json());
app.use(cors());

/* Router */
app.use('/home', Home);
app.use('/transactions', Transactions);
app.use('/stats', GamerStats);
app.use('/analytics', Analytics);
app.use('/redisReports', RedisReports);

/* Inicialización de socket.io */
const io = require('./controllers/sockets').init(server);
io.on('connection', (socket) => {
    console.log('A user is connected');
    socket.on('message', (message) => {
        console.log(`message from ${socket.id} : ${message}`);
    });
    socket.on('disconnect', () => {
        console.log(`Socket ${socket.id} disconnected`);
    });
});

/* Conexión de MongoDB */
const mongo = require('./controllers/mongo');
mongo.connectDB();

/* Conexión a Redis */
const redis = require('./controllers/redis');
redis.connectDB();


app.get('/', (req, res) => {
    res.send('Hello from Node.js Server!');
})

/* Starting */
server.listen(PORT, () => {
    console.log(`Node Server is running on port '${PORT}'.`);
});
