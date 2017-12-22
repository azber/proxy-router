package server

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"errors"
	"sync"
	"regexp"
)

type Service struct {
	addrList []string
	mu       *sync.Mutex
}

func NewService() (svr *Service) {
	nextIP := "127.0.0.1:12458"

	svr = &Service{
		addrList: make([]string, 0),
		mu:       &sync.Mutex{},
	}
	svr.addrList = append(svr.addrList, nextIP)
	return
}

func (svr *Service) NextIP(curAddr string) {
	svr.mu.Lock()
	defer svr.mu.Unlock()

	if len(svr.addrList) != 0 {
		if curAddr != svr.addrList[0] {
			return
		}
		svr.addrList = svr.addrList[1:]
	}

	list, err := svr.requestGetIP(svr.GetDefaultParams())
	if err != nil {
		return
	}
	for _, it := range list {
		r, _ := regexp.Compile(`((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d))))`)
		if r.Match([]byte(it)) {
			svr.addrList = append(svr.addrList, it)
		}
	}
	return
}

func (svr *Service) GetIP() (addr string, err error) {
	svr.mu.Lock()
	defer svr.mu.Unlock()

	if len(svr.addrList) != 0 {
		return svr.addrList[0], nil
	}
	list, err := svr.requestGetIP(svr.GetDefaultParams())
	if err != nil {
		return
	}
	if len(list) > 0 {
		addr = list[0]
		svr.addrList = append(svr.addrList, list...)
		return
	} else {
		err = errors.New("err, zhima get ip fatal")
		return
	}
}

func (svr *Service) requestGetIP(num, dataType, pro, city, yys, port, pack, ts, ys, cs, lb, sb, pb, mr int) (addrList []string, err error) {
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
	addrList = strings.Split(resStr, "\r\n")
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
