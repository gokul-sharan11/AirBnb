import express from 'express';
import { createHotelController, getAllHotelsController, getHotelByIdController } from '../../controllers/hotel.controller';
import { validateRequestBody } from '../../validators';
import { createHotelSchema } from '../../validators/createHotel.validator';

const hotelRouter = express.Router();

hotelRouter.get('/health', (req, res) => {
    res.status(200).send('OK');
});

hotelRouter.post('/', validateRequestBody(createHotelSchema), createHotelController);

hotelRouter.get('/:id',getHotelByIdController);

hotelRouter.get('/', getAllHotelsController);

export default hotelRouter;