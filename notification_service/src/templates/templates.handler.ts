import fs from 'fs/promises'
import Handlebars from 'handlebars';
import path from 'path'

export async function templatesGenerator(templateId : string, params : Record<string, any>) {
    const templatePath = path.join(__dirname, 'mailers',  `${templateId}.hbs`);
    try {
        const templateContent = await fs.readFile(templatePath, 'utf-8');
        const template = Handlebars.compile(templateContent);
        return template(params);
    } catch (error) {
        throw new Error(`Template not found: ${templateId}`);
    }
}