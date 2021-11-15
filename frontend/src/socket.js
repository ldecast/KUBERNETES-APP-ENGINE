import { io } from 'socket.io-client';

const socket = io(process.env.REACT_APP_API_URL);

socket.on('connnection', () => {
    console.log('Connected to socket');
});

socket.on('disconnect', (reason) => {
    console.log('Socket disconnected because of ' + reason);
});

socket.on('message', (message) => {
    console.log('Socket message: ' + message);
});

export default socket;