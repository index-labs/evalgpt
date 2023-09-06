package scheduler

import (
	"github.com/gin-gonic/gin"
	"github.com/index-labs/evalgpt/agent/python"
)

type Scheduler struct {
	pythonAgent *python.PythonAgent
}

type Config struct {
	PythonAgent *python.PythonAgent
}

func NewScheduler(cfg Config) *Scheduler {
	return &Scheduler{
		pythonAgent: cfg.PythonAgent,
	}
}

func (p *Scheduler) HandleQuery(query string, fileList []string) (result string, outputFiles []string, err error) {
	// TODO: add schedule logic here
	return p.pythonAgent.HandleQuery(query, fileList)
}

func (p *Scheduler) Run(addr string) (err error) {
	r := gin.Default()
	r.POST("/query", p.HandleQueryRequest)
	return r.Run(addr)
}
