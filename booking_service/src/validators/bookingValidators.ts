import { z } from "zod";

const bookingSchema = z.object({
    userId : z.number({message : "User Id is required"}),
    hotelId : z.number({message : "Hotel Id is required"}),
    bookingAmount : z.number({message : "Booking Amount is required"}).min(1, {message : "Booking Amount must be greater than 1"}),
    totalGuests : z.number({message : "Total Number of Guests is required"}).min(1, {message : "Total Number of Guests must be greater than 1"})
})

export default bookingSchema;