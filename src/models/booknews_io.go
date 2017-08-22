package models

type BooknewsList struct {
	Newid       string  `json:"newid" orm:"pk;size(20);column(newid);"`
	Userid_from	string	`json:"userid_from"`
	Userid_to	string	`json:"userid_to"`
	Bookqid	    string	`json:"bookqid"`
	Books       string  `json:"books"`
	User_from   string	`json:"user_from"`
	User_to     string	`json:"user_to"`
	Order_state uint8   `json:"order_state"`
	Create_time int64	`json:"create_time"`
	Update_time int64	`json:"update_time"`
}

