import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import { sequelize } from "./sequelize";

// Typescript layer - This enforces the following rules before object creation in typescript object
class Hotel extends Model<InferAttributes<Hotel>, InferCreationAttributes<Hotel>>{
    declare id : CreationOptional<number>;
    declare name : string;
    declare location : string;
    declare address : string;
    declare created_at : CreationOptional<Date>;
    declare updated_at : CreationOptional<Date>;
    declare rating ?: number;
    declare rating_count ?: number;
    declare deleted_at ?: CreationOptional<Date> | null;
}

Hotel.init(
    {
        id : {
            type : "INTEGER",
            autoIncrement : true,
            primaryKey : true
        },
        name : {
            type : "STRING",
            allowNull : false
        },
        location : {
            type : "STRING",
            allowNull : false
        },
        address : {
            type : "STRING",
            allowNull : false
        },
        created_at : {
            type : "DATE",
            defaultValue : new Date(),
        },
        updated_at : {
            type : "DATE",
            defaultValue : new Date(),
        },
        deleted_at : {
            type : "DATE",
            allowNull : true
        },
        rating : {
            type : "FLOAT",
            defaultValue : 0
        },
        rating_count : {
            type : "INTEGER",
            defaultValue : 0
        }
    },
    {
        tableName : "hotels",
        sequelize : sequelize,
        underscored : true,
        timestamps : true
    });

export default Hotel;