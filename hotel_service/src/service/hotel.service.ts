import { createHotelDTO } from "../dto/hotel.dto";
import { createHotel, findAllHotels, findHotelById } from "../repositories/hotel";

export async function createHotelService(hotelData : createHotelDTO) {
    const hotel = await createHotel(hotelData);
    return hotel;
}

export async function getHotelByIdService(id : number){
    const hotel = await findHotelById(id);
    return hotel;
}

export async function getAllHOtelsService(){
    const hotels = await findAllHotels();
    return hotels;
}