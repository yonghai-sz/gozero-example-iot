package amap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type amapAddressComponent struct {
	Province string      `json:"province"` // 坐标点所在省名称
	City     interface{} `json:"city"`     // 坐标点所在城市名称
	// District string      `json:"district"` // 坐标点所在区
}

type amapReGeoCode struct {
	FormattedAddress string                `json:"formatted_address"` // 结构化地址信息
	AddressComponent *amapAddressComponent `json:"addressComponent"`  // 地址元素列表
}

type AmapReGeoRsp struct {
	Status    string         `json:"status"` // 返回结果状态值: 0 表示请求失败；1 表示请求成功
	Info      string         `json:"info"`   // 返回状态说明: 当 status 为 0 时，info 会返回具体错误原因，否则返回“OK”
	InfoCode  string         `json:"infocode"`
	ReGeoCode *amapReGeoCode `json:"regeocode"` // 逆地理编码列表
}

// 逆地理编码：将经纬度转换为详细结构化的地址，且返回附近周边的POI、AOI信息。
// location 经纬度坐标  传入内容规则：经度在前，纬度在后，经纬度间以“,”分割，经纬度小数点后不要超过 6 位。
func ReGeocode(location string) (address string, province string, city string, err error) {

	url := fmt.Sprintf("%s/geocode/regeo?key=%s&output=json&location=%s", RestApiV3Url, "aaabbb", location)
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	rsp := AmapReGeoRsp{}
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		return
	}
	if rsp.Status == "1" { // 成功
		if rsp.ReGeoCode != nil {
			address = rsp.ReGeoCode.FormattedAddress
			province = rsp.ReGeoCode.AddressComponent.Province
			cityUnknow := rsp.ReGeoCode.AddressComponent.City
			switch cityUnknow := cityUnknow.(type) {
			case string:
				city = cityUnknow
			default:
				city = ""
			}
			return
		}
	}
	errmsg := "amap regeo api error"
	if rsp.Status == "0" {
		errmsg = rsp.Info
	}
	err = errors.New(errmsg)
	return
}
