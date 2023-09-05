package scheduler

import "github.com/index-labs/evalgpt/agent/python"

type Schedule struct {
	pythonAgent *python.PythonAgent
}

type Config struct {
	PythonAgent *python.PythonAgent
}

func NewScheduler(cfg Config) *Schedule {
	return &Schedule{
		pythonAgent: cfg.PythonAgent,
	}
}

func (p *Schedule) HandleQuery(query string, fileList []string) (result string, outputFiles []string, err error) {
	// TODO: add schedule logic here
	return p.pythonAgent.HandleQuery(query, fileList)
}
