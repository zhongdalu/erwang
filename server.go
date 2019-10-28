//@Auth:zdl
package erwang

import (
	"fmt"
	"gitee.com/sky_big/mylog"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/gcron"
	"github.com/zhongdalu/erwang/bll"
	"github.com/zhongdalu/erwang/public"
)

type ErrorCode struct {
	Errno int    `json:"errno"`
	Error string `json:"error"`
}

func Bind(httpUrl, wsbUrl string) {
	public.HttpUrl = httpUrl
	public.WebUrl = wsbUrl
	fmt.Println("使用二网平衡调度系统")
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("欢迎使用二网平衡调度系统\n版本：V1.0")
	})
	type msg struct {
		Status int `json:"status"`
		ErrorCode
		Ts int64 `json:"ts"`
	}

	var entry *gcron.Entry
	s.BindHandler("/erwang", func(r *ghttp.Request) {
		cmd := r.GetInt("cmd")
		var err error
		errCode := msg{}
		switch cmd {
		case -1:
			status := ""
			if entry != nil && entry.Status() == 0 {
				errCode.Errno = 0
				status = "开启"
				errCode.Status = 1
				errCode.Ts = entry.Time.Local().Unix()
			} else {
				errCode.Errno = 0
				status = "关闭"
				errCode.Status = 0
			}
			errCode.Error = "当前状态:" + status
		case 0:
			mylog.Println("关闭二网平衡")
			if entry != nil {
				gcron.Remove(entry.Name)
				entry = nil
			}
			errCode.Errno = 0
			errCode.Error = "关闭成功"
		case 1:
			mylog.Println("开启二网平衡")
			min := r.GetInt("min")
			if min == 0 {
				errCode.Errno = -3
				errCode.Error = "缺少参数:min"
			} else {
				if entry != nil {
					gcron.Remove(entry.Name)
					entry = nil
				}
				entry, err = gcron.Add(getCronStr(min), bll.Transfer)
				if err != nil {
					errCode.Errno = -2
					errCode.Error = "开启失败:" + err.Error()
				}
				g.Dump(gcron.Entries())
			}
		default:
			errCode.Errno = -1
			errCode.Error = "参数错误"
		}
		r.Response.Writeln(errCode)
	})
}

func getCronStr(min int) string {
	return fmt.Sprintf("@every %dm", min)
}
