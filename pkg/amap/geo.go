package amap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type amapGeocodes struct {
	FormattedAddress string `json:"formatted_address"` // 结构化地址信息
	Location         string `json:"location"`          // 坐标点：经度，纬度
}

type AmapGeoRsp struct {
	Status   string          `json:"status"`   // 返回结果状态值:返回值为 0 或 1，0 表示请求失败；1 表示请求成功
	Info     string          `json:"info"`     // 返回状态说明: 当 status 为 0 时，info 会返回具体错误原因，否则返回“OK”
	Count    string          `json:"count"`    // 返回结果的个数
	Geocodes *[]amapGeocodes `json:"geocodes"` // 地理编码信息列表
}

// 地理编码：将详细的结构化地址 转换为 高德经纬度坐标。且支持对地标性名胜景区、建筑物名称解析为高德经纬度坐标。
func Geocode(s string) (address string, location string, err error) {

	url := fmt.Sprintf("%s/geocode/geo?key=%s&output=json&address=%s", RestApiV3Url, "aaabbb", s)
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	rsp := AmapGeoRsp{}
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		return
	}
	if rsp.Status == "1" {
		var count int
		count, err = strconv.Atoi(rsp.Count)
		if err != nil {
			return
		}
		if count > 0 && rsp.Geocodes != nil && len(*rsp.Geocodes) > 0 {
			address = (*rsp.Geocodes)[0].FormattedAddress
			location = (*rsp.Geocodes)[0].Location
			return
		}
	}
	errmsg := "amap geo api error"
	if rsp.Status == "0" {
		errmsg = rsp.Info
	}
	err = errors.New(errmsg)
	return
}
