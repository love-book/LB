package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/astaxie/beego"
)

const(
	apiRetryNum = 2
)

type BarcodeInfo struct {
	Code       string
	Charge     bool
	Msg        string
	Result     Showapi
}

type Showapi struct {
	Showapi_res_error   string
	Showapi_res_code    int8
	Showapi_res_body    Showapi_res_body
}

type Showapi_res_body struct {
	Spec       string
    ManuName   string
	Ret_code   int8
	Flag       bool
	Price      string
	Trademark  string
	Img        string
	Code       string
	Note	   string
	Zzjb       string
	GoodsName  string
}


//获取条码信息
func  GetBarcodeInfo(code string) (uinf *BarcodeInfo,err error) {
	barcodeUrl  :=  beego.AppConfig.String("barcodeUrl")
	barcodeAppkey  :=  beego.AppConfig.String("barcodeAppkey")
	url := fmt.Sprintf("%s?code="+code+"&appkey=%s", barcodeUrl, barcodeAppkey)
	// retry
	for i := 0; i < apiRetryNum; i++ {
		resp, err := http.Get(url)
		if err != nil {
			if i < apiRetryNum-1 {
				continue
			}
			return uinf, err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if i < apiRetryNum-1 {
				continue
			}
			return uinf, err
		}
		// no
		if err := json.Unmarshal(data, &uinf); err != nil {
			if i < apiRetryNum-1 {
				continue
			}
			return nil, err
		}
		break // success
	}
	return uinf, nil
}

type LatAndLang struct {
	Status int64
	Result  []struct {
	    X float64    //long
		Y float64    //lat
	}
}

//微信gps转换成百度坐标
func GetGps(lat float64,lang float64) (l *LatAndLang,err error){
	ak  :=  beego.AppConfig.String("baiduAppkey")
	lats  := fmt.Sprintf("%f",lat)
	langs := fmt.Sprintf("%f",lang)
	coords := langs+","+lats
	url := "http://api.map.baidu.com/geoconv/v1/?coords="+coords+"&from=1&to=5&ak="+ak
	// retry
	for i := 0; i < apiRetryNum; i++ {
		resp, err := http.Get(url)
		if err != nil {
			if i < apiRetryNum-1 {
				continue
			}
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if i < apiRetryNum-1 {
				continue
			}
		}
		// no
		if err := json.Unmarshal(data, &l); err != nil {
			if i < apiRetryNum-1 {
				continue
			}
			return nil, err
		}
		break // success
	}
	return
}

type Location struct {
	Status int64
	Result  struct{
		Location  struct{
			Lng float64
			Lat float64
		}
		Formatted_address  string
		Business   string
		AddressComponent struct{
			Country string
		    Country_code int64
			Province   string
			City    string
			District   string
			Adcode   string
			Street   string
			Street_number   string
			Direction   string
			Distance   string
		}
	}
}

//通过百度坐标获取地理位置信息
func GetLocation(lat float64,long float64) (l *Location,err error){
	ak  :=  beego.AppConfig.String("baiduAppkey")
	lats  := fmt.Sprintf("%f",lat)
	longs := fmt.Sprintf("%f",long)
	location := lats+","+longs
	url := "http://api.map.baidu.com/geocoder/v2/?location="+location+"&output=json&pois=1&ak="+ak
	// retry
	for i := 0; i < apiRetryNum; i++ {
		resp, err := http.Get(url)
		if err != nil {
			if i < apiRetryNum-1 {
				continue
			}
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if i < apiRetryNum-1 {
				continue
			}
		}
		// no
		if err := json.Unmarshal(data, &l); err != nil {
			if i < apiRetryNum-1 {
				continue
			}
			return nil, err
		}
		break // success
	}
	return
}

//直接转换gps坐标返回百度地址信息
func  GetLocationInfo(lat float64,long float64) (l *Location,err error){
	LatAndLang,err := GetGps(lat,long)
	if  err!= nil{
		return nil,err
	}else{
		for _,v := range LatAndLang.Result{
			long = v.X
			lat  = v.Y
			break
		}
		l,err = GetLocation(lat,long)
        if err!= nil{
			return nil,err
		}
	}
	return
}