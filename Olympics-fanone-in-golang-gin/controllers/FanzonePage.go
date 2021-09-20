package controllers;


import (
	"fmt"
	"gitlab.com/satvikshrivas26/olympic-fanzone/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCheers(c *gin.Context) {
	var user []models.Fanzone
	_, err := dbmap.Select(&user, "select * from fanzone")
	if err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func AddCheers(c *gin.Context) {
	var user models.Fanzone
	country := c.PostForm("country")
	err := dbmap.SelectOne(&user, "select * from fanzone WHERE country=?", country)
	fmt.Println(err, user.Cheers)
	if err == nil {
		user.Cheers += 1
		_, err := dbmap.Exec("UPDATE fanzone SET cheers=?  WHERE country=? ", user.Cheers, user.Country)
		if err != nil {
			c.JSON(500, gin.H{"success": "DB request failed", "err": err.Error()})
		}else {
			c.JSON(200, gin.H{"success": "Updated"})
		}
	} else {
		_, err := dbmap.Exec("INSERT INTO  fanzone (country,cheers) VALUES(?,?) ", country, 1)
		if err != nil {
			c.JSON(500, gin.H{"success": "DB request failed", "err": err.Error()})
		}else {
			c.JSON(200, gin.H{"success": "Added"})
		}
	}
}
