const Pool = require('pg').Pool;

const pool = new Pool({
    user: "postgres",
    password: "Satvik26#",
    port: 5432,
    database: "Olympics_Fanzone_Trell"
});

module.exports = pool;