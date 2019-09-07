package model

/*
发送命令格式
{
	"uuid": "12121212",
	"fac": "deer_1_1",
	"order_id": 10074,
	"dt": "2019-03-27 14:34:00",
	"priority": 1,
	"timeout": 20,
	"recopy": 3,
	"data": [{
		"dtu": "5",
		"cjq": "1",
		"value": ["175,01,手动,90.1"]
	}]
}
*/
type CommandItem struct {
	Cjq   string   `json:"cjq"`
	Dtu   string   `json:"dtu"`
	Value []string `json:"value"`
}
type Command struct {
	Data     []CommandItem `json:"data"`
	Dt       string        `json:"dt"`
	Fac      string        `json:"fac"`
	OrderID  int           `json:"order_id"`
	Priority int           `json:"priority"`
	Recopy   int           `json:"recopy"`
	Timeout  int           `json:"timeout"`
	UUID     string        `json:"uuid"`
}

/*
ci第一次返回
{
	"xyh": 0,
	"uuid": "12121212",
	"dt": "2019-07-29 10:56:20",
	"cmd_id": ["bkv60l110j9327716fng"],
	"errno": 0,
	"error": ""
}
*/
type CiReturn0 struct {
	CmdID []string `json:"cmd_id"`
	Dt    string   `json:"dt"`
	Errno int      `json:"errno"`
	Error string   `json:"error"`
	UUID  string   `json:"uuid"`
	Xyh   int      `json:"xyh"`
}

/*
ci第二次返回
{
	"xyh": 1,
	"dt": "2019-07-29 10:57:48",
	"fac": "deer_1_1",
	"order_id": 0,
	"dtu": "5",
	"cjq": "1",
	"socket": "",
	"cmd_id": "bkv60l110j9327716fng",
	"data": null,
	"errno": -8,
	"error": "向网关发送命令失败=\u003e为网络错误=\u003ehttp://122.5.30.252:7109/receiveCmd:500 Internal Privoxy Error",
	"is_complete": true
}
*/
type CiReturn1 struct {
	Cjq        string      `json:"cjq"`
	CmdID      string      `json:"cmd_id"`
	Data       interface{} `json:"data"`
	Dt         string      `json:"dt"`
	Dtu        string      `json:"dtu"`
	Errno      int         `json:"errno"`
	Error      string      `json:"error"`
	Fac        string      `json:"fac"`
	IsComplete bool        `json:"is_complete"`
	OrderID    int         `json:"order_id"`
	Socket     string      `json:"socket"`
	Xyh        int         `json:"xyh"`
}
