const FanzonePageDB = require('./../models/FanzoneDB/FanzoneDB.js');
  

async function increaseFanCheers(req,res){
    try {
        const {COUNTRY, ADDCHEERS} = req.body;
        let cheerRecords = await FanzonePageDB.query(`UPDATE FANZONEPAGE SET CHEERCOUNTER += ADDCHEERS WHERE COUNTRY = '${COUNTRY}'`);
        cheerRecords = cheerRecords.rows;

        if(cheerRecords.length > 0){
            res.send(cheerRecords.rows);
        }else {
            res.send(`None of the countries got cheers`);
        }
    } catch (err) {
        res.send("Failed to fetch cheers of countires");
    }
}

async function updateCheersInTenSeconds(req,res){
    setInterval(async function(){ 
        try {
            const {COUNTRY, ADDCHEERS} = req.body;
            let cheerRecords = await FanzonePageDB.query(`UPDATE FANZONEPAGE SET CHEERCOUNTER += ADDCHEERS WHERE COUNTRY = '${COUNTRY}'`);
            cheerRecords = cheerRecords.rows;
    
            if(cheerRecords.length > 0){
                res.send(cheerRecords.rows);
            }else {
                res.send(`None of the countries got cheers`);
            }
        } catch (err) {
            res.send("Failed to update cheers of countires");
        }
    }, 10000);
}

module.exports = {increaseFanCheers, updateCheersInTenSeconds};