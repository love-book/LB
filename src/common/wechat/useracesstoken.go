package wechat

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type TicketToken struct {
	AppId     string
	AccessToken  string
	TmpName   string
	LckName   string
}

// get fresh access_token string
func (this *TicketToken) Fresh() (string, error) {
	if this.TmpName == "" {
		this.TmpName = this.AppId + "-tickettoken.tmp"
	}
	if this.LckName == "" {
		this.LckName = this.TmpName + ".lck"
	}
	for {
		if this.locked() {
			time.Sleep(time.Second)
			continue
		}
		break
	}
	fi, err := os.Stat(this.TmpName)
	if err != nil && !os.IsExist(err) {
		return this.fetchAndStore()
	}
	expires := fi.ModTime().Add(2 * time.Hour).Unix()
	if expires <= time.Now().Unix() {
		return this.fetchAndStore()
	}
	tmp, err := os.Open(this.TmpName)
	if err != nil {
		return "", err
	}
	defer tmp.Close()
	data, err := ioutil.ReadAll(tmp)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (this *TicketToken) fetchAndStore() (string, error) {
	if err := this.lock(); err != nil {
		return "", err
	}
	defer this.unlock()
	token, err := this.fetch()
	if err != nil {
		return "", err
	}
	if err := this.store(token); err != nil {
		return "", err
	}
	return token, nil
}

func (this *TicketToken) store(token string) error {
	path := path.Dir(this.TmpName)
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if !fi.IsDir() {
		return errors.New("path is not a directory")
	}
	tmp, err := os.OpenFile(this.TmpName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer tmp.Close()
	if _, err := tmp.Write([]byte(token)); err != nil {
		return err
	}
	return nil
}

func (this *TicketToken) fetch() (string, error) {
	url := fmt.Sprintf("%sticket/getticket?type=jsapi&access_token=", UrlPrefix)
	rtn, err := get(url+this.AccessToken)
	if err != nil {
		return "", err
	}
	return rtn.Ticket, nil
}

func (this *TicketToken) unlock() error {
	return os.Remove(this.LckName)
}

func (this *TicketToken) lock() error {
	path := path.Dir(this.LckName)
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	if !fi.IsDir() {
		return errors.New("path is not a directory")
	}
	lck, err := os.Create(this.LckName)
	if err != nil {
		return err
	}
	lck.Close()
	return nil
}

func (this *TicketToken) locked() bool {
	_, err := os.Stat(this.LckName)
	return !os.IsNotExist(err)
}
