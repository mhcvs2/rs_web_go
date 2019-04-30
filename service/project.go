package service

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/httplib"
	"github.com/beego/bee/logger"
	"strings"
	"time"
)

const (
	PROJECT_SERVICE_URL_KEY = "project_url"
)

var (
	projectServiceUrl string
	bm                cache.Cache
)

func init() {
	projectServiceUrl = beego.AppConfig.String(PROJECT_SERVICE_URL_KEY)
	bm, _ = cache.NewCache("memory", `{"interval":60}`)
}

type ProjectInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProjectServiceResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ProjectServiceGet(action string, headers map[string]string, params []string, v interface{}) error {
	var url string
	if params != nil && len(params) > 0 {
		url = fmt.Sprintf("%s/%s?%s", projectServiceUrl, action, strings.Join(params, "&"))
	} else {
		url = fmt.Sprintf("%s/%s", projectServiceUrl, action)
	}
	beeLogger.Log.Infof("get request url %s", url)
	req := httplib.Get(url)
	for k, v := range headers {
		req.Header(k, v)
	}
	var response ProjectServiceResponse
	responseBytes, err := req.Bytes()
	if err != nil {
		return err
	}
	beeLogger.Log.Infof("response: %s", string(responseBytes))
	err = req.ToJSON(&response)
	if err != nil {
		return err
	}
	if response.Status != 200 {
		return fmt.Errorf(response.Message)
	}
	if err = json.Unmarshal(responseBytes, v); err != nil {
		return err
	}
	return nil
}

type GetAllProjectsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []ProjectInfo
}

func GetAllProjects() ([]ProjectInfo, error) {
	if bm.IsExist("projects") {
		return bm.Get("projects").([]ProjectInfo), nil
	}
	var rep GetAllProjectsResponse
	err := ProjectServiceGet("api/projects", make(map[string]string), nil, &rep)
	if err != nil {
		return nil, err
	} else {
		err = bm.Put("projects", rep.Result, 100*time.Second)
		if err != nil {
			beeLogger.Log.Warnf("project put cache error: %s", err.Error())
		}
	}
	return rep.Result, nil
}
