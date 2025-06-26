import { Prisma } from "@prisma/client";

import { prismaClient } from "../prisma/client";

export async function createBooking(bookingInput : Prisma.BookingCreateInput) {
    // create a new booking
    const booking = await prismaClient.booking.create({
        data: bookingInput
    });
    return booking;
}