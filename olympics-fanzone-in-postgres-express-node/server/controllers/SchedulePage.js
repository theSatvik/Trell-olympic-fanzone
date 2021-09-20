const SchedulePageDB = require('./../models/SchedulePageDB/SchedulePageDB.js');
SchedulePageDB.connect();
const moment = require('moment');
function checkIsDate(obj) { 
    return moment.isDate(obj); 
 } 

async function scheduleAllEvents(req,res){
    try {
        let schudules = await SchedulePageDB.query(`SELECT _EVENT, STARTDATE,ENDDATE,COUNTRY FROM SCHEDULEPAGE WHERE CAST(STARTDATE AS DATE) > NOW()::date ORDER BY STARTDATE;`);
        schudules = schudules.rows;

        if(schudules.length > 0){
            res.send(schudules.rows);
        }else {
            res.send(`No events are available`);
        }
    } catch (err) {
        res.send("Failed to fetch schedules");
    }
}

async function filteredSchedule(req, res){
   
    try {
        const {STARTDATE, ENDDATE} = req.body;
        if( checkIsDate(STARTDATE) && checkIsDate(ENDDATE)){
            let filteredSchedules = await SchedulePageDB.query(`SELECT _EVENT, STARTDATE,ENDDATE,COUNTRY FROM SCHEDULEPAGE WHERE CAST(${STARTDATE} AS DATE) >= STARTDATE AND CAST(${ENDDATE} AS DATE) <= ENDDATE;`);
            filteredSchedules = filteredSchedules.rows;
            if(filteredSchedules.length > 0){
                res.send(filteredSchedules);
            }else {
                res.send(`No events are available for given filter`);
            }
        }else{
            res.send(`Invalid Date`);
        }
    } catch (err) {
        res.send("Failed to fetch schedules");
    }
}

async function liveSchedule(req, res){
   
    try {
        
        let liveSchedules = await SchedulePageDB.query(`SELECT _EVENT, STARTDATE,ENDDATE,COUNTRY FROM SCHEDULEPAGE WHERE CAST(STARTDATE AS DATE) <= NOW()::date AND CAST(ENDDATE AS DATE) >= NOW()::date;`);
        
        liveSchedules = liveSchedules.rows;
        if(liveSchedules.length > 0){
            res.send(liveSchedules);
        }else {
            res.send(`No events are live`);
        }
        
    } catch (err) {
        res.send("Failed to fetch Events");
    }
}

module.exports = {scheduleAllEvents, filteredSchedule, liveSchedule};