package python

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/index-labs/evalgpt/utils"

	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	PythonInterpreter string
	OpenaiApiKey      string
	OpenaiBaseURL     string
	Model             string
}

type PythonAgent struct {
	pythonInterpreter string
	model             string
	openaiClient      *openai.Client
}

func NewPythonAgent(cfg Config) *PythonAgent {

	openaiConfig := openai.DefaultConfig(cfg.OpenaiApiKey)
	if len(cfg.OpenaiBaseURL) > 0 {
		openaiConfig.BaseURL = cfg.OpenaiBaseURL
	}

	return &PythonAgent{
		pythonInterpreter: cfg.PythonInterpreter,
		openaiClient:      openai.NewClientWithConfig(openaiConfig),
		model:             cfg.Model,
	}
}

func (p *PythonAgent) HandleQuery(query string, fileList []string) (result string, outputFiles []string, err error) {
	log.Infof("handle query: %s", query)
	if len(fileList) > 0 {
		query += "/n/n"
		log.Infof("handle files: %s", strings.Join(fileList, ", "))
		for _, filename := range fileList {
			var fileMeta string
			fileMeta, err = extractFileMeta(filename)
			if err != nil {
				err = fmt.Errorf("extract file meta failed: %v, filename: %s", err, filename)
				return
			}
			query += fileMeta + "\n"
		}
		log.Infof("complete query: %s", query)
	}
	log.Infof("start call openai...")
	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: "Provide Python code that satisfies the user's request, ensuring the output is directed to stdout. Exclude any library installation instructions, only pure Python code is required"},
		{Role: openai.ChatMessageRoleUser, Content: query},
	}
	resp, err := p.openaiClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:       p.model,
		Messages:    messages,
		Temperature: 0,
	})

	if err != nil {
		log.Errorf("create chat completion failed: %v", err)
		return
	}
	msg := resp.Choices[0].Message
	log.Infof("call openai success, response message:\n%s", msg.Content)

	log.Infof("extract python code from openai response msg")
	code, err := extractPythonCode(msg.Content)
	if err != nil {
		log.Errorf("extract python code failed: %v, content: %s", err, msg.Content)
		return
	}
	log.Infof("extract python code success, code:\n%s", code)
	result, outputFiles, err = p.Interpret(code, fileList)
	if err != nil {
		log.Errorf("interpret code failed: %v", err)
		return
	}
	return
}

func (p *PythonAgent) Interpret(code string, fileList []string) (output string, outputFiles []string, err error) {
	workDir, err := os.MkdirTemp("/tmp", "interpreter_")
	if err != nil {
		err = fmt.Errorf("create tmp work dir failed: %v", err)
		return
	}
	defer func() {
		_ = os.RemoveAll(workDir)
	}()

	inputFilenamesDict := map[string]bool{}
	for _, filename := range fileList {
		baseFilename := path.Base(filename)
		inputFilenamesDict[baseFilename] = true
		err = os.Link(filename, path.Join(workDir, baseFilename))
		if err != nil {
			err = fmt.Errorf("link file failed: %v, filename: %s", err, filename)
			return
		}
	}

	pyFilename := "__main__.py"
	inputFilenamesDict[pyFilename] = true

	codeFilepath := path.Join(workDir, pyFilename)
	err = os.WriteFile(codeFilepath, []byte(code), 0644)
	if err != nil {
		err = fmt.Errorf("write code file failed: %v", err)
		return
	}
	output, err = utils.RunCmdWithTimeout(workDir, -1, p.pythonInterpreter, codeFilepath)
	if err != nil {
		err = fmt.Errorf("run cmd failed: %v", err)
		return
	}

	allFiles, err := os.ReadDir(workDir)
	if err != nil {
		err = fmt.Errorf("read work dir failed: %v", err)
		return
	}
	for _, f := range allFiles {
		if f.IsDir() {
			continue
		}
		if _, ok := inputFilenamesDict[f.Name()]; ok {
			continue
		}

		err = os.Rename(path.Join(workDir, f.Name()), f.Name())
		if err != nil {
			err = fmt.Errorf("rename output file failed: %v", err)
			return
		}
		outputFiles = append(outputFiles, f.Name())
	}
	return
}

func extractFileMeta(filename string) (meta string, err error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return
	}

	meta = fmt.Sprintf("filename: %s", fileInfo.Name())
	if strings.HasSuffix(filename, ".csv") {
		var lines []string
		lines, err = utils.ReadFileLines(filename, 1)
		if err != nil {
			return
		}
		if len(lines) > 0 {
			meta += fmt.Sprintf(", csv header: %s", lines[0])
		}
	}
	return
}

func extractPythonCode(content string) (code string, err error) {
	ss := strings.SplitN(content, "```python", -1)
	if len(ss) < 2 {
		err = fmt.Errorf("cant't find python code")
		return
	}
	ss = strings.SplitN(ss[1], "```", -1)
	code = strings.TrimSpace(ss[0])
	return
}
