package main

import (
	"fmt"
	"os"

	"github.com/index-labs/evalgpt/agent/python"
	"github.com/index-labs/evalgpt/scheduler"
	"github.com/index-labs/evalgpt/utils"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "evalgpt",
		HelpName:    "evalgpt help",
		Description: "description",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "openai-api-key",
				Usage:    "Openai Api Key, if you use open ai model gpt3 or gpt4, you must set this flag",
				EnvVars:  []string{"OPENAI_API_KEY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "model",
				Usage:   "LLM name",
				Value:   "gpt-4-0613",
				EnvVars: []string{"MODEL"},
			},
			&cli.StringFlag{
				Name:    "python-interpreter",
				Usage:   "python interpreter path",
				Value:   "/usr/bin/python3",
				EnvVars: []string{"PYTHON_INTERPRETER"},
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Usage:       "print verbose log",
				DefaultText: "false",
				EnvVars:     []string{"VERBOSE"},
			},
			&cli.StringFlag{
				Name:     "query",
				Aliases:  []string{"q"},
				Usage:    "what you want to ask",
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:  "file",
				Usage: "the path to the file to be parsed and processed, eg. --file /tmp/a.txt --file /tmp/b.txt",
			},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Errorf("error: %v", err)
	}
}

func run(ctx *cli.Context) error {
	openaiApiKey := ctx.String("openai-api-key")
	model := ctx.String("model")
	pythonInterpreter := ctx.String("python-interpreter")
	query := ctx.String("query")
	verbose := ctx.Bool("verbose")
	fileList := ctx.StringSlice("file")

	if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	log.SetFormatter(&log.TextFormatter{
		ForceQuote: true,
	})

	pipeData, err := utils.ReadStdinPipeData()
	if err != nil {
		log.Errorf("read stdin failed: %v", err)
		return err
	}

	if len(pipeData) > 0 {
		query += fmt.Sprintf(`\n"""%s"""`, pipeData)
	}
	log.Infof("model: %s, pythonInterpreter: %s", model, pythonInterpreter)

	pyAgent := python.NewPythonAgent(python.Config{
		PythonInterpreter: pythonInterpreter,
		OpenaiApiKey:      openaiApiKey,
		Model:             model,
	})

	sched := scheduler.NewScheduler(scheduler.Config{
		PythonAgent: pyAgent,
	})

	result, outputFiles, err := sched.HandleQuery(query, fileList)
	if err != nil {
		return err
	}
	log.Infof("===== result =====")
	if len(result) > 0 {
		fmt.Println(result)
	}
	if len(outputFiles) > 0 {
		for _, filename := range outputFiles {
			fmt.Println("created file:", filename)
		}
	}
	return nil
}
