package worker

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"code.lstaas.com/lightspeed/atom/log"

	"github.com/MrVegeta/go-playground/analyst/common"

	"github.com/MrVegeta/go-playground/analyst/feature/worker"

	"code.lstaas.com/lightspeed/atom"
)

const (
	galaxyPrefix  = "client"
	voyagerPrefix = "voyager"
)

type Worker struct {
	SDKLogPath      string
	KeywordTemplate string
}

func New(ctx context.Context, config *Config) (*Worker, error) {
	return &Worker{
		SDKLogPath:      config.SDKLogPath,
		KeywordTemplate: config.KeywordTemplate,
	}, nil
}

// Type implements atom.HasType
func (h *Worker) Type() interface{} {
	return worker.HandlerType()
}

// Start implements atom.Runnable.
func (h *Worker) Start() error {
	log.New("*******************************************************").AtInfo().WriteToLog()
	log.New("****** 重要提示：请在SDK同级路径下运行本程序！！！*****").AtInfo().WriteToLog()
	log.New("*******************************************************").AtInfo().WriteToLog()

	var err error
	var logList []string
	var templates []*common.Template
	if templates, err = h.loadTemplate(); err != nil {
		return err
	}
	if logList, err = h.getLogFilenames(h.SDKLogPath, logList); err != nil {
		return err
	}
	if err := h.parseLogs(logList, templates); err != nil {
		log.NewF("%v", err).AtInfo().WriteToLog()
	}
	return nil
}

// Close implements atom.Closable.
func (h *Worker) Close() error {
	return nil
}

func (h *Worker) loadTemplate() ([]*common.Template, error) {
	var err error
	var data []byte
	if data, err = ioutil.ReadFile(h.KeywordTemplate); err != nil {
		return nil, err
	}
	var templates []*common.Template
	if err := json.Unmarshal(data, &templates); err != nil {
		return nil, err
	}

	return templates, nil
}

func (h *Worker) getLogFilenames(path string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := path + "/" + fi.Name()
			s, err = h.getLogFilenames(fullDir, s)
			if err != nil {
				return s, err
			}
		} else {
			fullName := path + "/" + fi.Name()
			if strings.Contains(fullName, ".log") {
				s = append(s, fullName)
			}
		}
	}
	return s, nil
}

func (h *Worker) parseLogs(logList []string, templates []*common.Template) error {
	if len(logList) == 0 {
		return fmt.Errorf("***** 未检测到SDK日志文件，请确认SDK是否正常运行 *****")
	}
	if len(templates) == 0 {
		return fmt.Errorf("empty template")
	}

	t := time.Now()
	log.New("开始检测....").AtInfo().WriteToLog()

	// 1. search keyword: match
	matched := make(map[string]int, 0)
	for _, v := range templates {
		if v.Match == "" {
			continue
		}
		matched[v.Match] = 0
	}
	for _, v := range logList {
		if !strings.Contains(v, galaxyPrefix) && !strings.Contains(v, voyagerPrefix) {
			continue
		}

		filename := v
		file, err := os.Open(filename)
		if err != nil {
			log.NewF("failed opening file: %v %v", filename, err).AtDebug().WriteToLog()
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			for _, v2 := range templates {
				if v2.Match != "" && strings.Contains(line, v2.Match) {
					matched[v2.Match]++
				}
			}
		}
	}
	// 2. print summary
	for k, v := range matched {
		log.NewF("[%v] 匹配到 [%v] 条记录", k, v).AtInfo().WriteToLog()
	}
	if len(matched) == 0 {
		log.New("*******************************************************").AtInfo().WriteToLog()
		log.New("********** 匹配结果为空，SDK可能未正确运行！***********").AtInfo().WriteToLog()
		log.New("*******************************************************").AtInfo().WriteToLog()
	}
	// 3. search keyword: error
	errors := make(map[string]int, 0)
	for _, v := range logList {
		filename := v
		if !strings.Contains(filename, galaxyPrefix) && !strings.Contains(filename, voyagerPrefix) {
			continue
		}

		var logContent []string
		if data, err := ioutil.ReadFile(filename); err == nil {
			logContent = strings.Split(string(data), "\n")
		}

		for _, v2 := range templates {
			for _, v3 := range logContent {
				keyword := v2.Error
				line := v3
				if keyword != "" && strings.Contains(line, keyword) {
					errors[keyword]++
				}
			}
		}
	}
	if len(errors) > 0 {
		for k, v := range errors {
			log.NewF("[%v] 匹配到 [%v] 条记录", k, v).AtInfo().WriteToLog()
		}
		log.New("*******************************************************").AtInfo().WriteToLog()
		log.New("*******************************************************").AtInfo().WriteToLog()
		log.New("*******************************************************").AtInfo().WriteToLog()
		log.New("****** 发现异常，请与我们联系以获得技术支持！！！******").AtInfo().WriteToLog()
		log.New("*******************************************************").AtInfo().WriteToLog()
		log.New("*******************************************************").AtInfo().WriteToLog()
		log.New("*******************************************************").AtInfo().WriteToLog()
	} else {
		log.NewF("****** 未匹配到错误 ******").AtInfo().WriteToLog()
	}
	log.NewF("检测结束，耗时：%v", time.Since(t)).AtInfo().WriteToLog()

	return nil
}

func (h *Worker) extractSessionID(json string) (string, error) {
	result := gjson.Get(json, "message")
	sub := strings.Split(result.String(), " ")
	if len(sub) == 0 {
		return "", nil
	}
	sessionID := sub[0]
	return sessionID, nil
}

func init() {
	atom.Must(atom.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*Config))
	}))
}
