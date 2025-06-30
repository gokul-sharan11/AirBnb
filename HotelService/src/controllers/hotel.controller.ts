import { Request, Response, NextFunction } from "express";
import { createHotelService, getAllHOtelsService, getHotelByIdService } from "../service/hotel.service";
import { StatusCodes } from "http-status-codes";

export async function createHotelController(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await createHotelService(req.body);
    res.status(StatusCodes.CREATED).json({
        message : "Hotel created successfully",
        data : hotelResponse,
        success : true
    });
}

export async function getHotelByIdController(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await getHotelByIdService(Number(req.params.id));
    res.status(200).json({
        message : "Hotel found successfully",
        data : hotelResponse,
        success : true
    });
}

export async function getAllHotelsController(req: Request, res: Response, next: NextFunction) {
    const hotelResponse = await getAllHOtelsService();
    res.status(200).json({
        message : "Hotels found successfully",
        data : hotelResponse,
        success : true
    });
}