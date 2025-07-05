import { createBookingDto } from "../dto/booking.dto";
import { prismaClient } from "../prisma/client";
import { createBooking, createIdempotencyKey, getIdempotencyKeywithLocks, finalizeIdempotencyKey, confirmBooking} from "../repo/booking.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";

export async function createBookingService(createBookingDTO : createBookingDto) {
    const booking = await createBooking({
        userId : createBookingDTO.userId,
        hotelId : createBookingDTO.hotelId,
        bookingAmount : createBookingDTO.bookingAmount,
        totalGuests : createBookingDTO.totalGuests
    });
    const idempotencyKey = generateIdempotencyKey();
    await createIdempotencyKey(idempotencyKey, booking.id);
    return {bookingId : booking.id, idemmpotencyKey : idempotencyKey};
}

export async function confirmBookingService(idempotencyKey : string) {
    return await prismaClient.$transaction(async (tx) => {
        const idemmpotencyKeyData = await getIdempotencyKeywithLocks(tx, idempotencyKey);
        if(!idemmpotencyKeyData){
            throw new NotFoundError('Idempotency key not found');
        }
        if(idemmpotencyKeyData.finalized){
            throw new BadRequestError('Idempotency key already finalized');
        }
        const booking = confirmBooking(tx, idemmpotencyKeyData.bookingId);
        await finalizeIdempotencyKey(tx, idempotencyKey);
        return booking;
    })
    
}