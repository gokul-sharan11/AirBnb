import logger from "../config/logger.config";
import { NextFunction, Request, Response } from "express";
import { createBookingService, confirmBookingService } from "../services/booking.service";

export async function createBookingHandler(req: Request, res: Response, next: NextFunction) {
    logger.info("Booking request received");
    const booking = await createBookingService(req.body);
    res.status(201).json({bookingId : booking.bookingId, idemmpotencyKey : booking.idemmpotencyKey});
}

export async function confirmBookingHandler(req: Request, res: Response, next: NextFunction) {
    logger.info("Confirm booking request received");
    const booking = await confirmBookingService(req.params.idempotencyKey);
    res.status(200).json({
        bookingId : booking.id,
        status : booking.status
    });
}