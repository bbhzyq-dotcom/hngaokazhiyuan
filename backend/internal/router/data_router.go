package router

import (
	"github.com/gin-gonic/gin"

	"github.com/gaokao-advisor/backend/internal/handler"
)

type DataRouter struct {
	dataHandler *handler.DataHandler
}

func NewDataRouter(dataHandler *handler.DataHandler) *DataRouter {
	return &DataRouter{dataHandler: dataHandler}
}

func (r *DataRouter) RegisterRoutes(rg *gin.RouterGroup) {
	colleges := rg.Group("/colleges")
	{
		colleges.GET("", r.dataHandler.GetColleges)
		colleges.GET("/:id", r.dataHandler.GetCollegeDetail)
		colleges.GET("/:id/majors", r.dataHandler.GetCollegeMajors)
	}

	majors := rg.Group("/majors")
	{
		majors.GET("", r.dataHandler.GetMajors)
		majors.GET("/:id", r.dataHandler.GetMajorDetail)
	}

	scores := rg.Group("/scores")
	{
		scores.GET("", r.dataHandler.GetScores)
		scores.GET("/probability", r.dataHandler.GetProbability)
	}

	rankings := rg.Group("/rankings")
	{
		rankings.GET("/colleges", r.dataHandler.GetColleges)
		rankings.GET("/majors", r.dataHandler.GetMajors)
	}
}
