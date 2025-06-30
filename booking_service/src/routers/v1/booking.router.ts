import { Router } from "express";
import { createBookingHandler, confirmBookingHandler } from "../../controllers/booking.controller";
import { validateRequestBody } from "../../validators";
import bookingSchema from "../../validators/bookingValidators";

const bookingRouter =  Router();

bookingRouter.post('/', validateRequestBody(bookingSchema), createBookingHandler);
bookingRouter.post('/confirm/:idempotencyKey', confirmBookingHandler);

export default bookingRouter;