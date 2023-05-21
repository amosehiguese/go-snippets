package snippet

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/amosehiguese/go-snippet/dbutils"
	"github.com/gin-gonic/gin"
	// _ "github.com/mattn/go-sqlite3"
)

var nDb *sql.DB

type StationResource struct {
	ID  int `json:"id"`
	Name string `json:"name"`
	OpeningTime  string  `json:"opening_time"`
	ClosingTime  string  `json:"closing_time"`
}

func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")
	err := nDb.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

func CreateStation(c *gin.Context) {
	var station StationResource
	if err := c.BindJSON(&station); err == nil {
		statement, _ := nDb.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		}else {
			c.String(http.StatusInternalServerError, err.Error())
		} 
	} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
}

func RemoveStation(c *gin.Context) {
	id := c.Param("station-id")
	statement, _ := nDb.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func DBmain() {
	var err error 
	nDb, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Init(Db)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/stations/:station_id", GetStation)
		v1.POST("/stations", CreateStation)
		v1.DELETE("/station/:station_id", RemoveStation)

	}

	r.Run(":8080")

}