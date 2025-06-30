import { Prisma } from "@prisma/client";

import { prismaClient } from "../prisma/client";

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
            key,
            booking : {
                connect : {
                    id : bookingId
                }
            } 
        }
    });
    return idempotencyKey;
}

export async function getIdempotencyKey(key : string) {
    // get idempotency key
    const idempotencyKey = await prismaClient.idempotencyKey.findUnique({
        where : {
            key
        }
    });
    return idempotencyKey;
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

export async function confirmBooking(bookingId : number) {
    // finalize booking
    const booking = await prismaClient.booking.update({
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

export async function finalizeIdempotencyKey(key : string) {
    // finalize idempotency key
    const idempotencyKey = await prismaClient.idempotencyKey.update({
        where : {
            key
        },
        data : {
            finalized : true
        }
    });
    return idempotencyKey;
}