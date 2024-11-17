package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type TableData struct {
	ID          int            `json:"id"`
	RangeFrom   int            `json:"range_from"`
	RangeTo     int            `json:"range_to"`
	DaysColumns map[string]int `json:"days_columns"`
}

var (
	db                                              *sql.DB
	DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME string
)

func main() {
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")

	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err = sql.Open("mysql", connection)
	if err != nil {
		log.Fatal("error connecting to db", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to db at %s:%s", DB_HOST, DB_PORT)

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/data", getData)
	r.POST("/data", postData)
	log.Fatal(r.Run(":6969"))
}

func getData(c *gin.Context) {
	rows, err := db.Query("SELECT id, range_from, range_to, days_columns FROM admin_main")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query data"})
		return
	}
	defer rows.Close()

	var result []TableData
	for rows.Next() {
		var data TableData
		var daysColumns string
		err := rows.Scan(&data.ID, &data.RangeFrom, &data.RangeTo, &daysColumns)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list data"})
			return
		}

		data.DaysColumns = parseJSON(daysColumns)
		result = append(result, data)
	}

	c.JSON(http.StatusOK, result)
}

func postData(c *gin.Context) {
	var newData []TableData
	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, data := range newData {
		daysColumnsJSON := toJSON(data.DaysColumns)
		query := `INSERT INTO admin_main (range_from, range_to, days_columns) VALUES (?, ?, ?)`
		_, err := db.Exec(query, data.RangeFrom, data.RangeTo, daysColumnsJSON)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
			return
		}
	}

	c.JSON(http.StatusOK, newData)
}

func toJSON(data map[string]int) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func parseJSON(jsonStr string) map[string]int {
	var data map[string]int
	json.Unmarshal([]byte(jsonStr), &data)
	return data
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
