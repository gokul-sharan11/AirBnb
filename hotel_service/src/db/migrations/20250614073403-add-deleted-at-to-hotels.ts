import { QueryInterface } from "sequelize";

module.exports = {
  async up (queryInterface : QueryInterface) {  
    await queryInterface.addColumn('hotels', 'deleted_at', {
      type: 'TIMESTAMP',
      allowNull: true,
      defaultValue: null
    })
  },

  async down (queryInterface : QueryInterface) {
    await queryInterface.removeColumn('hotels', 'deleted_at')
  }
};
