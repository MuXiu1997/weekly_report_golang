package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weekly_report_golang/report"
)

var Router *gin.Engine

func init() {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.StaticFile("/", "./template/index.html")
	r.GET("/report", getReport)
	r.PATCH("/report", updateReport)

	Router = r
}

func getReport(c *gin.Context) {
	r := report.New().Load()
	c.JSON(http.StatusOK, gin.H{
		"labels": r.Labels(),
		"works":  r.Works(),
		"keys":   r.Keys(),
	})
}

func updateReport(c *gin.Context) {
	request := struct {
		Works map[string]string `json:"works"`
	}{}
	err := c.BindJSON(&request)
	if err != nil {
		panic(err)
	}
	works := request.Works
	r := report.New()
	r.Update(works).Save()
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
