package models

type BookrackList struct {
	Userid       string   `json:"userid" valid:"Required" orm:"pk;size(20);column(userid);"`
	Bookid       string   `json:"bookid" valid:"Required"`
	Bookstate    string   `json:"bookstate"`
	Create_time  int64    `json:"create_time"`
	Update_time  int64    `json:"update_time"`
	Bookname     string `json:"bookname"`
	Author       string	`json:"auhtor" `
	Imageurl     string	`json:"imageurl"`
	Depreciation uint8	`json:"depreciation"`
	State        uint8	`json:"state"`
	Nickname  	string 	`json:"nickname"`
	Imgurl    	string 	`json:"imgurl" `
	Gender  	int8  	`json:"gender"`
	Age   		int32 	`json:"age"`
}

type BookrackInfo struct {
	Userid       string   `json:"userid" valid:"Required" orm:"pk;size(20);column(userid);"`
	Bookid       string   `json:"bookid" valid:"Required"`
	Bookstate    string   `json:"bookstate"`
	Create_time  int64    `json:"create_time"`
	Update_time  int64    `json:"update_time"`
	Books
	Users
}

func  BooksrackList(start int,length int,conditions string,filed string) (book []BookrackList,total int){
	var books []BookrackList
	var  TableName = "lb_bookrack as r left join lb_books  as b on r.bookid = b.bookid  left join lb_users as u on u.userid = r.userid "
	totalItem, res :=GetPagesInfo(TableName,start,length,conditions,filed)
	res.QueryRows(&books)
	return  books,totalItem
}


func  BooksrackInfo(start int,length int,conditions string,filed string) (book []BookrackInfo,total int){
	var books []BookrackInfo
	var  TableName = "lb_bookrack as r left join lb_books  as b on r.bookid = b.bookid  left join lb_users as u on u.userid = r.userid "
	totalItem, res :=GetPagesInfo(TableName,start,length,conditions,filed)
	res.QueryRows(&books)
	return  books,totalItem
}