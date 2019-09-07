package model

/*
{
	"code": 1,
	"msg": "成功",
	"data": [{
		"guandaoId": 1,
		"name": "管道1",
		"model": "1",//1.手动模式 2 进回水温差 3 进回水均温 4 回水温度
		"p1": 146.5,
		"p2": 126.5,
		"p3": 106.5,
		"cha1": 0.5,
		"cha2": 2,
		"lists": [{
			"fac": "2",
			"dtu": "27",
			"center": "27",
			"mp_no": "5",
			"faNo": "2019060105",
			"faModel": "1",
			"targetValue": 0
		}]
	}]
}
*/

type F struct {
	Center      string  `json:"center"`
	Dtu         string  `json:"dtu"`
	FaModel     string  `json:"faModel"`
	FaNo        string  `json:"faNo"`
	Fac         string  `json:"fac"`
	MpNo        string  `json:"mp_no"`
	TargetValue float64 `json:"targetValue"`
}

type Pipe struct {
	Cha1      float64 `json:"cha1"`
	Cha2      float64 `json:"cha2"`
	GuandaoID int     `json:"guandaoId"`
	Lists     []F     `json:"lists"`
	Model     string  `json:"model"`
	Name      string  `json:"name"`
	P1        float64 `json:"p1"`
	P2        float64 `json:"p2"`
	P3        float64 `json:"p3"`
}
type BsData struct {
	Code int    `json:"code"`
	Data []Pipe `json:"data"`
	Msg  string `json:"msg"`
}
