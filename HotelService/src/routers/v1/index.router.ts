import express from 'express';
import hotelRouter from './hotel.router';
import pingRouter from './ping.router';

const v1Router = express.Router();


v1Router.get('/health', (req, res) => {
    res.status(200).send('OK');
})
v1Router.use('/hotels',  hotelRouter);
v1Router.use('/ping',  pingRouter);

export default v1Router;