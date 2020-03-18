package bll

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"gitee.com/sky_big/mylog"
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/gorilla/websocket"

	. "github.com/zhongdalu/erwang/model"
	"github.com/zhongdalu/erwang/public"
	"github.com/zhongdalu/erwang/util"
)

// 获取web接口的数据 判断是否需要调平衡 需要的话 发送命令给命令交互
func Transfer() {
	bs, err := util.HttpGet(public.HttpUrl + "/api/server/getConduitData")
	if err != nil {
		mylog.Error(err)
		return
	}
	err = record()
	if err != nil {
		return
	}
	var data BsData
	err = json.Unmarshal(bs, &data)
	if err != nil {
		mylog.Error(err)
		return
	}
	for _, p := range data.Data {
		// 2、	查看p1、p2、p3的差是否有超出范围的 如果有执行下一步
		if math.Abs(p.P1-p.P2) < p.Cha1 && math.Abs(p.P1-p.P3) < p.Cha1 && math.Abs(p.P3-p.P2) < p.Cha1 {
			pa, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", (p.P1+p.P2+p.P3)/3.0), 64)
			for _, f := range p.Lists {
				// 3、	将p1的值与当前单元阀的目标值比较 如果超出范围 则执行下一步
				if math.Abs(p.P1-f.TargetValue) > p.Cha2 {
					// 4、	向单元阀 发送命令 设置阀模式为管道模式 目标值为(p1+p2+p3)/3
					sendCommand(f.Fac, f.Dtu, f.Center, f.MpNo, p.Model, pa)
				}
			}
		} else {
			mylog.Println(p.Name + "p1、p2、p3的差超出范围,温度尚未稳定")
		}
	}
}

func record() error {
	urls := public.HttpUrl + "/api/balance/getzhixing_pingheng"
	_, err := util.HttpGet(urls)
	if err != nil {
		mylog.Error(err)
		return err
	}
	return nil
}

func sendCommand(fac, dtu, center, mpno, mod string, value float64) {
	/*
		 命令
		 "01",    //"手动",
		 "02",    //"温差",
		 "03",    //"均温",
		 "04",    //"回水",

		1.手动模式 2 进回水温差 3 进回水均温 4 回水温度
	*/
	arr := []string{"175", mpno}
	switch mod {
	case "1":
		// arr = append(arr, "01")
		// 手动模式不做处理
		return
	case "2":
		arr = append(arr, "02")
	case "3":
		arr = append(arr, "03")
	case "4":
		arr = append(arr, "04")
	default:
		mylog.Error("未知命令:", mod)
	}
	arr = append(arr, gconv.String(value))
	v := strings.Join(arr, ",")
	data := Command{
		Data: []CommandItem{
			{
				Cjq:   center,
				Dtu:   dtu,
				Value: []string{v},
			},
		},
		Dt:       gtime.Datetime(),
		Fac:      fac,
		OrderID:  10074,
		Priority: 3,
		Recopy:   3,
		Timeout:  50,
		UUID:     util.Rand().Hex(),
	}
	bs, _ := json.Marshal(data)

	c, _, err := websocket.DefaultDialer.Dial(public.WebUrl, nil)
	if err != nil {
		mylog.Error("dial:", err)
	}
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, bs)
	if err != nil {
		mylog.Error(err)
		return
	} else {
		mylog.Println("发送给CI：" + string(bs))
	}
}
