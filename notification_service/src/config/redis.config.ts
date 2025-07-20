import ioredis, { Redis } from 'ioredis';
import { serverConfig } from '.';

// Singleton Design Pattern 
const connectToRedis = () => {
    try {
        let connection : Redis;
        const redisConfig = {
            host: serverConfig.REDIS_HOST,
            port: serverConfig.REDIS_PORT,
            maxRetriesPerRequest: null
        };

        return () => {
            if(!connection) {
                connection = new ioredis(redisConfig);
            }
            return connection;
        }
    } catch (error) {
        throw new Error("Error connecting to redis");
    }
}

export const getRedisConnectionObject = connectToRedis();



