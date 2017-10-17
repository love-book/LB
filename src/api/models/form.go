package models

//app
type AccesstokenForm struct {
	Openid    string    `json:"openid"`
}

type WxConfigForm struct {
	Url    string    `json:"url"`
}
type WxcodeForm struct {
	Code    string    `json:"code"`
}

type LoginForm struct {
	Telphone  string	`json:"telphone"`
	Password  string	`json:"password"`
}

type SmsForm struct {
	Telphone  string	`json:"telphone"`
}

type SmsCheckForm struct {
	Telphone  string	`json:"telphone"`
	Code  string	`json:"code"`
}

type UseraddForm struct {
	Openid  	string	 `json:"openid"`
	Wnickname  	string	 `json:"wnickname"`
	Wimgurl    	string 	 `json:"wimgurl"`
	Nickname  	string 	 `json:"nickname"`
	Imgurl    	string 	 `json:"imgurl" `
	Gender  	int64  	 `json:"gender"`
	Age   		int64 	 `json:"age"`
	Telphone  	string   `json:"telphone"`
	Password	string   `json:"password"`
	Qq  		string   `json:"qq"`
	Weino  		string   `json:"weibo"`
	Wechat		string   `json:"wechat"`
	Signature  	 string   `json:"signature"`
	Constellation string `json:"constellation"`
	Province  	string   `json:"province"`
	City  	    string   `json:"city"`
	Address  	string   `json:"address"`
	Long  	    float64  `json:"long"`
	Lat  	    float64  `json:"lat"`
}


type UserupdateForm struct {
	Openid  	string	 `json:"openid"`
	Wnickname  	string	 `json:"wnickname"`
	Wimgurl    	string 	 `json:"wimgurl"`
	Nickname  	string 	 `json:"nickname"`
	Imgurl    	string 	 `json:"imgurl" `
	Gender  	int64  	 `json:"gender"`
	Age   		int64 	 `json:"age"`
	Telphone  	string   `json:"telphone"`
	Code  	    string   `json:"code"`
	Password	string   `json:"password"`
	Qq  		string   `json:"qq"`
	Weino  		string   `json:"weibo"`
	Wechat		string   `json:"wechat"`
	Signature  	 string   `json:"signature"`
	Constellation string `json:"constellation"`
	Province  	string   `json:"province"`
	City  	    string   `json:"city"`
	Address  	string   `json:"address"`
	Long  	    float64  `json:"long"`
	Lat  	    float64  `json:"lat"`
}

//book
type BookaddForm struct {
	Bookname     string `json:"bookname"`
	Author       string	`json:"author"`
	Imageurl     string	`json:"imageurl"`
	Imagehead    string	`json:"imagehead"`
	Imageback    string	`json:"imageback"`
	Isbn         string	`json:"isbn"`
	Depreciation uint8	`json:"depreciation"`
	Price        uint16	`json:"price"`
	Describe     string	`json:"describe"`
	State        uint8	`json:"state"`
}

type AddconcernForm struct {
	UseridFrom   string `json:"userid_from"`
	ConcernType  string `json:"concern_type"`
}

type  DelbookconcernForm struct {
	Concernid   []string `json:"concernid"`
}

type  ConcernBookListForm struct {
	Length int `json:"length"`
	Draw   int `json:"draw"`
	Userid string `json:"userid"`
}
type  BookListForm struct {
	Length int `json:"length"`
	Draw   int `json:"draw"`
}

//Booksrack

type  BookracklistForm struct {
	Length int `json:"length"`
	Draw   int `json:"draw"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
	Radius string `json:"radius"`
}

type  BookrackaddForm struct {
	Bookid string `json:"bookid"`
}

type  BookrackaddbysnForm struct {
	Isbn string `json:"isbn"`
}


type BookStateForm struct {
	Bookqid    []string   `json:"bookqid"`
}

type  GetbookusersForm struct {
	Length int `json:"length"`
	Draw   int `json:"draw"`
	Bookid string  `json:"bookid"`
}


type  GetbookinfoForm struct {
	Bookid string  `json:"bookid"`
}

type  GetuserinfoForm struct {
	Userid string  `json:"userid"`
}

//bookNews

type  NewslistForm struct {
	Length int `json:"length"`
	Draw   int `json:"draw"`
	//OrderState string  `json:"order_state"`
}
type  LibraryrequestForm struct {
	Bookqid     string  `json:"bookqid"`
	Telphone  	string   `json:"telphone"`
	Qq  		string   `json:"qq"`
	Wechat	    string   `json:"wechat"`
}

type  AgreeLibraryrequestForm struct {
	Newid       string  `json:"newid"`
	Bookqid     string  `json:"bookqid"`
	Telphone  	string   `json:"telphone"`
	Qq  		string   `json:"qq"`
	Wechat	    string   `json:"wechat"`
}

type  RefuseLibraryrequestForm struct {
	Newid      string  `json:"newid"`
}
//User
type  UserinfoForm  struct {
	Userid   string  `json:"userid"`
}

type  GetUsersByLocaltionForm  struct {
	Length int `json:"length"`
	Draw   int `json:"draw"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
	Radius string `json:"radius"`
	Logintime string `json:"logintime"`
}


type  AddLocaltionByIDForm  struct {
	Lang   float64  `json:"lang"`
	Lat    float64  `json:"lat"`
}


type  AddOpinionsForm  struct {
	Opinions  string	`json:"opinions"`
	Images    []string	`json:"images"`
}


// order
type  OrderlistForm  struct {
	Length   int `json:"length"`
	Draw     int `json:"draw"`
	OrderState   string `json:"order_state"`
}

type  OrderupdateForm  struct {
	Orderid      string `json:"orderid"`
	OrderState   string `json:"order_state"`
}

type  OrderdeleteForm  struct {
	Orderid      string `json:"orderid"`
}

type  GetuserbookinfoForm struct {
	Bookqid string  `json:"bookqid"`
}