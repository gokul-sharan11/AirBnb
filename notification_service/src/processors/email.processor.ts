import { Worker } from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE_NAME } from "../queues/mailer.queue";
import { getRedisConnectionObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";

export const setUpMailerWorker = () => {
    const emailProcessor = new Worker<NotificationDto>(
        MAILER_QUEUE_NAME, // Queue Name
        // Function that has to be called to process the job
        async (job) => {
            if(job.name !== MAILER_PAYLOAD) {
                throw new Error("Invalid job name");
            }
            const payload = job.data;
            console.log(`Processing email for ${JSON.stringify(payload)}`);
        }, 
        // connection details
        {
            connection: getRedisConnectionObject()
        }
    );
    
    // on success event emitted by emailProcessor Worker
    emailProcessor.on("completed", (job) => {
        console.log("Email processing completed successfully");
    })
    
    emailProcessor.on("failed", (job) => {
        console.log("Email processing failed")
    })
}