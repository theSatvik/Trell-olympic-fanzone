package controllers;
import (
	// "log"
	// "strconv"
	// "time"
	"gitlab.com/satvikshrivas26/olympic-fanzone/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func MedalTally(c *gin.Context) {
	
	var medalRecords []models.Medal
	_, err := dbmap.Select(&medalRecords, "SELECT * FROM MEDAL ORDER BY GOLD, SILVER, BRONZE DESC;");
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch medal tally"})
		}else if(len(medalRecords) > 0){
			c.JSON(200, medalRecords)
		}else {
			c.JSON(200, gin.H{"message": "None of the countries won any medal"})
		}
}
func UpdateMedalTally(c *gin.Context){

	// _, err := dbmap.Select(&UpdatedRecords, " SELECT * FROM SCHEDULE WHERE CAST(STARTDATE AS DATE) <= CURDATE() AND CAST(ENDDATE AS DATE) >= CURDATE();");
	var UpdatedRecords []models.Medal
	
		if err := c.ShouldBindJSON(&UpdatedRecords); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

}