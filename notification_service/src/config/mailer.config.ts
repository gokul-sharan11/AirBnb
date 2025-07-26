import nodemailer from "nodemailer";
import { serverConfig } from ".";

const transporter = nodemailer.createTransport({
    service: "gmail",
    auth: {
        user: serverConfig.MAILER_USER,
        pass: serverConfig.MAILER_PASSWORD
    }
});

export default transporter;