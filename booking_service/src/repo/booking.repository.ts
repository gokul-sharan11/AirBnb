import { idempotencyKey, Prisma } from "@prisma/client";
import { validate } from "uuid";
import { prismaClient } from "../prisma/client";
import { InternalServerError, NotFoundError } from "../utils/errors/app.error";

export async function createBooking(bookingInput : Prisma.BookingCreateInput) {
    // create a new booking
    const booking = await prismaClient.booking.create({
        data: bookingInput
    });
    return booking;
}

export async function createIdempotencyKey(key : string, bookingId : number) {
    // create a new idempotency key
    const idempotencyKey = await prismaClient.idempotencyKey.create({
        data: {
            idemKey : key,
            booking : {
                connect : {
                    id : bookingId
                }
            } 
        }
    });
    return idempotencyKey;
}

export async function getIdempotencyKeywithLocks(tx : Prisma.TransactionClient, key : string) {
    if(!validate(key)){
        throw new InternalServerError('Invalid idempotency key format');
    }
    // get idempotency key
    const idempotencyKey : Array<idempotencyKey> = await tx.$queryRaw(Prisma.raw(`SELECT * FROM idempotencyKey WHERE idemKey = '${key}' FOR UPDATE;`));
    if(!idempotencyKey || idempotencyKey.length === 0){
        throw new NotFoundError("Idempotent key not found");
    }
    return idempotencyKey[0];
}

export async function getBookingById(bookingId : number) {
    // get booking by id
    const booking = await prismaClient.booking.findUnique({
        where : {
            id : bookingId
        }
    });
    return booking;
}

export async function confirmBooking(tx : Prisma.TransactionClient, bookingId : number) {
    // finalize booking
    const booking = await tx.booking.update({
        where : {
            id : bookingId
        },
        data : {
            status : "CONFIRMED"
        }
    });
    return booking;
}

export async function cancelBooking(bookingId : number) {
    // cancel booking
    const booking = await prismaClient.booking.update({
        where : {
            id : bookingId
        },
        data : {
            status : "CANCELLED"
        }
    });
    return booking;
}

export async function finalizeIdempotencyKey(tx : Prisma.TransactionClient, key : string) {
    // finalize idempotency key
    const idempotencyKey = await tx.idempotencyKey.update({
        where : {
            idemKey : key
        },
        data : {
            finalized : true
        }
    });
    return idempotencyKey;
}