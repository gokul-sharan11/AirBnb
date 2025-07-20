import { Queue } from "bullmq";
import { getRedisConnectionObject } from "../config/redis.config";

export const MAILER_QUEUE_NAME = "queue-mailer";

const mailerQueue = new Queue(MAILER_QUEUE_NAME, {
    connection: getRedisConnectionObject()
});

export default mailerQueue;