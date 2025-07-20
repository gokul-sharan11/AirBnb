import IORedis, {Redis} from 'ioredis';
import Redlock from 'redlock';
import { serverConfig} from '.';

// export const redisClient = new IORedis(serverConfig.REDIS_SERVER_URL);

const connectToRedis = () => {
    try {
        let connection : Redis;

        return () => {
            if(!connection) {
                connection = new IORedis(serverConfig.REDIS_SERVER_URL);
            }
            return connection;
        }
    } catch (error) {
        throw new Error("Error connecting to redis");
    }
}

export const getRedisConnectionObject = connectToRedis();

export const redlockClient = new Redlock([getRedisConnectionObject()], {
    driftFactor: 0.01,
    retryCount: 10,
    retryDelay: 200,
    retryJitter: 200
});