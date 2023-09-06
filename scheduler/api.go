package scheduler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryRequest struct {
	Query string `json:"query"`
}

type QueryResponse struct {
	Result string `json:"result"`
}

func (p *Scheduler) HandleQueryRequest(c *gin.Context) {
	req := &QueryRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid request"})
		return
	}
	result, _, err := p.HandleQuery(req.Query, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, QueryResponse{Result: result})
	return
}
