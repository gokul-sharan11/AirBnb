import { serverConfig } from "../config";
import transporter from "../config/mailer.config";

export async function sendMail(subject: string, to: string, body: string) {
    try {
        const info = await transporter.sendMail({
            from: serverConfig.MAILER_USER,
            to : to,
            subject : subject,
            html : body
        });
        return info;
    } catch (error) {
        throw error;
    }
}