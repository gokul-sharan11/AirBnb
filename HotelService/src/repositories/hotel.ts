import { where } from "sequelize";
import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { createHotelDTO } from "../dto/hotel.dto";
import { NotFoundError } from "../utils/errors/app.error";


export async function createHotel(hotelData : createHotelDTO) {
    const hotel = await Hotel.create({name : hotelData.name, location : hotelData.location, address : hotelData.address, rating : hotelData.rating, rating_count : hotelData.rating_count});
    logger.info(`Hotel created successfully ${hotel.id}`);
    return hotel;
}

export async function findHotelById(id : number) {
    const hotel = await Hotel.findByPk(id);
    if(!hotel) {
        logger.error(`Hotel with id ${id} not found`);
        throw new NotFoundError(`Hotel with id ${id} not found`);
    }
    return hotel;
}

export async function findAllHotels() {
    const hotels = await Hotel.findAll({where: {deleted_at : null}});
    return hotels;
}

export async function softDeleteHotels(id : number){
    const hotel = await Hotel.findByPk(id);
    if(!hotel) {
        logger.error(`Hotel with id ${id} not found`);
        throw new NotFoundError(`Hotel with id ${id} not found`);
    }
    hotel.deleted_at = new Date();
    await hotel.save();
    return true;
}

