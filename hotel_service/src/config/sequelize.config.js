require('ts-node/register') // This line enables TypeScript support, since our db.config and sequelize rc supports js we need to convert the migration files from ts to js on the go 
const config = require('./db.config')
module.exports = config