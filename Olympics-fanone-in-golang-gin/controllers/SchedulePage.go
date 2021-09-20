package controllers
import (
	// "log"
	// "strconv"
	"time"
	"gitlab.com/satvikshrivas26/olympic-fanzone/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
func ScheduleAllEvents(c *gin.Context) {
	var countries []models.Schedule
	_, err := dbmap.Select(&countries, "select * from Schedule")
	if err == nil {
		c.JSON(200, countries)
	} else {
		c.JSON(404, gin.H{"error": "countries not found"})
	}
}
// func validateDate(myDateString string){
// 	myDate, err := time.Parse("2006-01-02", myDateString)
// 	if err != nil {
// 		panic(err)
// 		return err;
// 	}
// 	return myDate
// }

func FilteredSchedule(c *gin.Context) {
	startDate := c.Query("STARTDATE")
	endDate := c.Query("ENDDATE")
	
	_, err1 := time.Parse("2006-01-02", startDate)

	_, err2 := time.Parse("2006-01-02", endDate)

	if err1 != nil || err2 != nil {
		// panic("Failed")
		c.JSON(404, gin.H{"error": "Invalid filter"})
	}else {
		var filteredSchedules []models.Schedule
	_, err := dbmap.Select(&filteredSchedules, "SELECT * FROM SCHEDULE WHERE ? >= STARTDATE AND ? <= ENDDATE;", startDate, endDate);
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch"})
		}else if(len(filteredSchedules) > 0){
			c.JSON(200, filteredSchedules)
		}else {
			c.JSON(200, gin.H{"error": "No events are available for given filter"})
		}
	}
}

func LiveSchedule(c *gin.Context){
	var liveSchedules []models.Schedule
	_, err := dbmap.Select(&liveSchedules, " SELECT * FROM SCHEDULE WHERE CAST(STARTDATE AS DATE) <= CURDATE() AND CAST(ENDDATE AS DATE) >= CURDATE();");
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch"})
		}else if(len(liveSchedules) > 0){
			c.JSON(200, liveSchedules)
		}else {
			c.JSON(200, gin.H{"error": "No events are available for given filter"})
		}
}

// func GetUserDetail(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=? LIMIT 1", id)
// 	if err == nil {
// 		user_id, _ := strconv.ParseInt(id, 0, 64)
// 		content := &models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: user.Firstname,
// 			Lastname:  user.Lastname,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }
// func Login(c *gin.Context) {
// 	var user models.User
// 	c.Bind(&user)
// 	err := dbmap.SelectOne(&user, "select * from user where Username=? LIMIT 1", user.Username)
// 	if err == nil {
// 		user_id := user.Id
// 		content := &models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: user.Firstname,
// 			Lastname:  user.Lastname,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }
// func PostUser(c *gin.Context) {
// 	var user models.User
// 	c.Bind(&user)
// 	log.Println(user)
// 	if user.Username != "" && user.Password != "" && user.Firstname != "" && user.Lastname != "" {
// 		if insert, _ := dbmap.Exec(`INSERT INTO user (Username, Password, Firstname, Lastname) VALUES (?, ?, ?, ?)`, user.Username, user.Password, user.Firstname, user.Lastname); insert != nil {
// 			user_id, err := insert.LastInsertId()
// 			if err == nil {
// 				content := &models.User{
// 					Id:        user_id,
// 					Username:  user.Username,
// 					Password:  user.Password,
// 					Firstname: user.Firstname,
// 					Lastname:  user.Lastname,
// 				}
// 				c.JSON(201, content)
// 			} else {
// 				checkErr(err, "Insert failed")
// 			}
// 		}
// 	} else {
// 		c.JSON(400, gin.H{"error": "Fields are empty"})
// 	}
// }
// func UpdateUser(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user models.User
// 	err := dbmap.SelectOne(&user, "SELECT * FROM user WHERE id=?", id)
// 	if err == nil {
// 		var json models.User
// 		c.Bind(&json)
// 		user_id, _ := strconv.ParseInt(id, 0, 64)
// 		user := models.User{
// 			Id:        user_id,
// 			Username:  user.Username,
// 			Password:  user.Password,
// 			Firstname: json.Firstname,
// 			Lastname:  json.Lastname,
// 		}
// 		if user.Firstname != "" && user.Lastname != "" {
// 			_, err = dbmap.Update(&user)
// 			if err == nil {
// 				c.JSON(200, user)
// 			} else {
// 				checkErr(err, "Updated failed")
// 			}
// 		} else {
// 			c.JSON(400, gin.H{"error": "fields are empty"})
// 		}
// 	} else {
// 		c.JSON(404, gin.H{"error": "user not found"})
// 	}
// }