package common

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func filterUser(ctx *context.Context) {
	if ctx.GetCookie("userId") == "" {
		response := NewNoLoginResponse(nil)
		res, _ := json.Marshal(response)
		ctx.ResponseWriter.Write(res)
	}

}

func init() {

	beego.InsertFilter("/*", beego.BeforeRouter, filterUser)
}
