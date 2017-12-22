package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"errors"
)

type Service struct {
	addrList []string
}

func NewService() *Service {
	return &Service{
		addrList: make([]string, 0),
	}
}

func (svr *Service) NextIP() (addr string, err error) {
	if len(svr.addrList) != 0 {
		addr = svr.addrList[0]
		svr.addrList = svr.addrList[1:]
		return
	}
	list, err := svr.GetIP(svr.GetDefaultParams())
	if err != nil {
		return
	}
	if len(list) > 0 {
		addr = list[0]
		svr.addrList = list[1:]
	} else {
		err = errors.New("err, zhima get ip fatal")
		return
	}
	return
}

func (svr *Service) GetIP(num, dataType, pro, city, yys, port, pack, ts, ys, cs, lb, sb, pb, mr int) (addrList []string, err error) {
	reqUrl := fmt.Sprintf("http://webapi.http.zhimacangku.com/getip?num=%d&type=%d&pro=%d&city=%d&yys=%d&port=%d&pack=%d&ts=%d&ys=%d&cs=%d&lb=%d&sb=%d&pb=%d&mr=%d&regions=",
		num, dataType, pro, city, yys, port, pack, ts, ys, cs, lb, sb, pb, mr)
	res, err := http.Get(reqUrl)
	defer res.Body.Close()
	if err != nil {
		return
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	resStr := string(resBody)
	addrList = strings.Split(resStr, "\n")
	return
}

func (svr *Service) GetDefaultParams() (num, dataType, pro, city, yys, port, pack, ts, ys, cs, lb, sb, pb, mr int) {
	num = 1
	dataType = 1
	pro = 0
	city = 0
	yys = 0
	port = 2
	pack = 9538
	ts = 0
	ys = 0
	cs = 0
	lb = 1
	sb = 0
	pb = 45
	mr = 2
	return
}
