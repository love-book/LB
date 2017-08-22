package models

type BookrackList struct {
	Bookqid      string  `json:"bookqid" orm:"pk;size(20);column(bookqid);"`
	Userid       string  `json:"userid"`
	Bookid       string  `json:"bookid"`
	Book_state   string  `json:"bookstate"`
	Is_borrow    string  `json:"is_borrow"`
	Create_time  int64   `json:"create_time"`
	Update_time  int64   `json:"update_time"`
	Bookname     string `json:"bookname"`
	Author       string	`json:"auhtor""`
	Imageurl     string	`json:"imageurl"`
	Imagehead    string	`json:"imagehead"`
	Imageback    string	`json:"imageback"`
	Isbn         string	`json:"isbn"`
	Depreciation uint8	`json:"depreciation"`
	Price        uint16	`json:"price"`
	Describe     string	`json:"describe"`
	State        uint8	`json:"state" `
}


type BookExist struct {
	Userid       string   `json:"userid"`
	Bookid       string   `json:"bookid"`
}


