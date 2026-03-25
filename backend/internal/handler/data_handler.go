package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gaokao-advisor/backend/internal/service"
	"github.com/gaokao-advisor/backend/pkg/response"
)

type DataHandler struct {
	dataService *service.DataService
}

func NewDataHandler(dataService *service.DataService) *DataHandler {
	return &DataHandler{dataService: dataService}
}

func (h *DataHandler) GetColleges(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	req := &service.CollegeListRequest{
		Page:     page,
		PageSize: pageSize,
		Keyword:  c.Query("keyword"),
		Province: c.Query("province"),
		Type:     c.Query("type"),
		Level:    c.DefaultQuery("level", "本科"),
	}

	resp, err := h.dataService.GetColleges(req)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *DataHandler) GetCollegeDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "invalid college id")
		return
	}

	college, err := h.dataService.GetCollegeByID(id)
	if err != nil {
		response.NotFound(c, "college not found")
		return
	}

	response.Success(c, college)
}

func (h *DataHandler) GetCollegeMajors(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "invalid college id")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	category := c.Query("category")

	majors, total, err := h.dataService.GetCollegeMajors(id, page, pageSize, category)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"total": total,
		"list":  majors,
	})
}

func (h *DataHandler) GetMajors(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	req := &service.MajorListRequest{
		Page:     page,
		PageSize: pageSize,
		Keyword:  c.Query("keyword"),
		Category: c.Query("category"),
		Degree:   c.Query("degree"),
	}

	resp, err := h.dataService.GetMajors(req)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, resp)
}

func (h *DataHandler) GetMajorDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ParamError(c, "invalid major id")
		return
	}

	major, err := h.dataService.GetMajorByID(id)
	if err != nil {
		response.NotFound(c, "major not found")
		return
	}

	response.Success(c, major)
}

func (h *DataHandler) GetScores(c *gin.Context) {
	collegeID, _ := strconv.ParseUint(c.Query("college_id"), 10, 64)
	majorID, _ := strconv.ParseUint(c.Query("major_id"), 10, 64)
	year, _ := strconv.Atoi(c.DefaultQuery("year", "2024"))

	req := &service.ScoreQueryRequest{
		CollegeID: collegeID,
		MajorID:   majorID,
		Year:      year,
		Province:  c.DefaultQuery("province", "河南省"),
		Batch:     c.Query("batch"),
		Category:  c.Query("category"),
	}

	resp, err := h.dataService.GetScores(req)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, gin.H{"list": resp})
}

func (h *DataHandler) GetProbability(c *gin.Context) {
	collegeID, err := strconv.ParseUint(c.Query("college_id"), 10, 64)
	if err != nil {
		response.MissingParam(c, "college_id is required")
		return
	}

	score, _ := strconv.Atoi(c.Query("score"))
	rank, _ := strconv.Atoi(c.Query("rank"))

	req := &service.ProbabilityRequest{
		CollegeID: collegeID,
		MajorID:   0,
		Score:     score,
		Rank:      rank,
		Category:  c.DefaultQuery("category", "理科"),
	}

	resp, err := h.dataService.GetProbability(req)
	if err != nil {
		response.Error(c, response.CodeInternalError, err.Error())
		return
	}

	response.Success(c, resp)
}
