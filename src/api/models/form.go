package models

//app
type AccesstokenForm struct {
	Openid    string    `json:"openid"`
}

type LoginForm struct {
	Telphone  string	`json:"telphone"`
	Password  string	`json:"password"`
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
	Author       string	`json:"auhtor"`
	Imageurl     string	`json:"imageurl"`
	Imagehead    string	`json:"imagehead"`
	Imageback    string	`json:"imageback"`
	Isbn         string	`json:"isbn"`
	Depreciation uint8	`json:"depreciation"`
	Price        uint16	`json:"price"`
	Describe     string	`json:"describe"`
	Userid		 string	`json:"userid"`
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
	Bookqid        []string   `json:"bookqid"`
	Book_state     string   `json:"book_state"`
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
	OrderState string  `json:"order_state"`
}
type  LibraryrequestForm struct {
	Bookqid     string  `json:"bookqid"`
	Telphone  	string   `json:"telphone"`
	Qq  		string   `json:"qq"`
	Wechat	    string   `json:"wechat"`
}

type  AgreeLibraryrequestForm struct {
	Newid       string  `json:"newid"`
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
	Nickname string  `json:"nickname"`
	Telphone string  `json:"telphone"`
}

type  GetUsersByLocaltionForm  struct {
	Length int `json:"length"`
	Draw   int `json:"draw"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
	Radius string `json:"radius"`
}


type  AddLocaltionByIDForm  struct {
	Lang   float64  `json:"lang"`
	Lat    float64  `json:"lat"`
}


// order
type  OrderlistForm  struct {
	Length   int `json:"length"`
	Draw     int `json:"draw"`
	OrderState   string `json:"order_state"`
	Isbn	     string `json:"isbn"`
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