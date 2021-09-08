const express = require('express');
const app = express();
const cors = require("cors");
const moment = require('moment');

const SchedulePage = require('./controllers/SchedulePage');

const MedalPage = require('./controllers/MedalPage');

const FanzonePage = require('./controllers/Fanzone');


const port = 3000;

app.use(express.urlencoded({ extended: false }));
app.use(express.json());
app.use(cors());

//Middlewares
// Schudulespage
app.get('/schedulePage', SchedulePage.scheduleAllEvents , (req, res) => { });

app.get('/schedulePage/filter', SchedulePage.filteredSchedule , (req, res) => { });


app.get('/schedulePage/filter', SchedulePage.liveSchedule , (req, res) => { });


//Medal Page 
// Medal Tally
app.get('/MedalPage', MedalPage.MedalTally , (req, res) => { });

// update medal tally
app.post('/MedalPage', MedalPage.UpdateMedalTally, (req, res) => { });



// Fanzone Page 
// Medal Tally
app.post('/FanzonePage', FanzonePage.increaseFanCheers , (req, res) => { });

// update medal tally
app.post('/FanzonePage', FanzonePage.updateCheersInTenSeconds, (req, res) => { });


app.listen(port, () => {
    console.log(`Server is listening at port : ${port}`);
});