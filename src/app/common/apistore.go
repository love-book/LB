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
	Charge     string
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
func  GetBarcodeInfo(code string) (BarcodeInfo, error) {
	barcodeUrl  :=  beego.AppConfig.String("barcodeUrl")
	barcodeAppkey  :=  beego.AppConfig.String("barcodeAppkey")
	var uinf BarcodeInfo
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
			return uinf, err
		}
		break // success
	}
	return uinf, nil
}

