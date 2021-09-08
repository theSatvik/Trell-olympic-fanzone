const MedalPageDB = require('./../models/MedalPageDB/MedalPageDB.js');


async function MedalTally(req,res){
    try {
        let medalRecords = await MedalPageDB.query(`SELECT COUNTRY, GOLD, SILVER, BRONZE FROM MEDALPAGE ORDER BY GOLD, SILVER, BRONZE DESC;`);
        medalRecords = medalRecords.rows;

        if(medalRecords.length > 0){
            res.send(medalRecords.rows);
        }else {
            res.send(`None of the countries won any medal`);
        }
    } catch (err) {
        res.send("Failed to fetch medal tally");
    }
}

async function UpdateMedalTally(req, res){
   
    try {
        const {COUNTRY, MEDALTYPE, ADDNEWMEDAL} = req.body;
    
        let UpdatedRecords = await SchedulePageDB.query(` UPDATE MedalPage
        SET  ${MEDALTYPE} +=  ADDNEWMEDAL
        WHERE COUNTRY = ${COUNTRY} AND MEDALTYPE = ${MEDALTYPE}; `);
        
        UpdatedRecords = UpdatedRecords.rows;
        if(UpdatedRecords.length > 0){
            res.send(UpdatedRecords);
        }else {
            res.send(`No records were updated`);
        }
    
    } catch (err) {
        res.send("Failed to fetch schedules");
    }
}


module.exports = {MedalTally, UpdateMedalTally};