package amap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

/*
// wifi 接入网络：1
type WhenAccessType1 struct {
	// 非必填，但强烈建议填写。已连 热点 mac 信息： mac/signal/ssid  如：
	// mac：    f0:7d:68:9e:7d:18
	// signal： -41
	// ssid：   TPLink
	Mmac string `json:"mmac"`

	Macs string `json:"macs"` // accesstype=1 时，必填。 wifi 列表中 mac 信息。 单 mac 信息同 mmac，mac 之间使用“|” 分隔。 必须填写 2 个及 2 个以上，30 个以内的方可正常定位。请不要包含移动 wifi 信息
}

// 移动接入网络：0
type WhenAccessType0 struct {
	Cdma    uint8  `json:"cdma"`    // accesstype=0 时，必填。 是否为 cdma，缺省值：0             非 cdma：0                是 cdma：1
	Network string `json:"network"` // accesstype=0 时，必填。 无线网络类型。                    GSM/GPRS/EDGE/HSUPA/HSDPA/WC DMA

	// accesstype=0 时，必填。接入基站信息。
	// 非 CDMA 格式为：mcc/mnc/lac/cellid/signal
	//    CDMA 格式为：sid/nid/bid/lon/lat/signal  其中 lon/lat 可为空，格式为：sid/nid/bid/,,,/signal
	Bts string `json:"bts"`

	Nearbts string `json:"nearbts"` // 非必填，周边基站信息（不含接入基站信息）。基站信息1|基站信息2|基站信息3…..
}
*/

// type GetPositionReq struct {
// 	Key        string `json:"key"`
// 	Output     string `json:"output"`
// 	AccessType *uint8 `json:"accesstype"` // 必填。    移动端 接入网络 方式。        可选值：  0：移动 接入网络        1：wifi 接入网络

// 	Imei     string `json:"imei"`     // 必填。 手机 imei 号。此参数能够提高定位精度和准确度，且定位不到时可依据此参数进行跟踪排查，如没有则无法排查和跟踪问题
// 	Idfa     string `json:"idfa"`     // ios   手机的 idfa，与 imei 必填一个。此参数能够提高定位精度和准确度，且定位不到时可依据此参数进行跟踪排查，如没有则无法排查和跟踪问题
// 	Smac     string `json:"smac"`     // 非必填，但建议填写。手机 mac 码。此参数能够提高定位精度和准确度，且定位不到时可依据此参数进行跟踪排查，如没有则无法排查和跟踪问题
// 	ServerIp string `json:"serverip"` // 非必填，但建议填写。设备接入基站时对应的网关 IP。 此参数能够提高定位精度和准确度，且定位不到时可依据此参数进行跟踪排查，如没有则无法排查和跟踪问题
// 	Imsi     string `json:"imsi"`     // 非必填，但建议填写。移动用户识别码。            此参数能够提高定位精度和准确度，且定位不到时可依据此参数进行跟踪排查，如没有则无法排查和跟踪问题
// 	Tel      string `json:"tel"`      // 非必填。手机号码。
// }

type PositionResult struct {
	ResultType string `json:"type"`     // 定位类型  0：没有得到定位结果；其他数字为：正常获取定位结果
	Location   string `json:"location"` // 定位经纬度
	Desc       string `json:"desc"`     // 位置描述
	Province   string `json:"province"` // 省份
}

type GetPositionRsp struct {
	Status   string          `json:"status"` // 返回结果状态值: 0 表示失败；1 表示成功
	Info     string          `json:"info"`   // 返回状态说明: status 为 0 时，info 返回错误原因，否则返回“OK”
	InfoCode string          `json:"infocode"`
	Result   *PositionResult `json:"result"` // 定位结果列表
}

func GetPosition(imei string, bts, nearbts string, mmac, macs string, network string) (result *PositionResult, err error) {

	accesstype := 0
	if bts == "" {
		accesstype = 1
	}
	url := fmt.Sprintf("%s/position?key=%s&output=json&accesstype=%d", ApiLocateUrl, AmapKeyApiLocate, accesstype)
	url += "&cdma=0"
	if imei != "" {
		url += "&imei=" + imei
	}
	if bts != "" {
		url += "&bts=" + bts
	}
	if nearbts != "" {
		url += "&nearbts=" + nearbts
	}
	if mmac != "" {
		url += "&mmac=" + mmac
	}
	if macs != "" {
		url += "&macs=" + macs
	}
	if network != "" {
		url += "&network=" + network
	}

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	rsp := GetPositionRsp{}
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		return
	}
	if rsp.Status == "1" { // 成功
		result = rsp.Result
		return
	}
	errmsg := "amap position api error"
	if rsp.Status == "0" {
		errmsg = rsp.Info
	}
	err = errors.New(errmsg)
	return
}
