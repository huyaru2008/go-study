package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"github.com/hsyan2008/go-logger"
)

func main() {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		// colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36"),
		colly.Async(true),
	)
	//随机UserAgent
	extensions.RandomUserAgent(c)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
		Parallelism: 15,
	})
	//代理
	rp, err := proxy.RoundRobinProxySwitcher("代理url")
	if err != nil {
		logger.Error(err, "代理错误", err.Error())
	}
	c.SetProxyFunc(rp)
	StepParent(c)
	c.Visit("https://stockx.com/")
	c.Wait()
}

func StepParent(c *colly.Collector) {
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	})
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != http.StatusOK || len(r.Body) == 0 {
			return
		}
		fmt.Println(string(r.Body))
	})
	c.OnError(func(r *colly.Response, err error) {

		if err1 := Retry(r.Request, 3); err1 != nil {
			return
		}
	})

}

//重试
func Retry(r *colly.Request, count int) error {
	key := fmt.Sprintf("err_req_%s", r.URL.String())
	var et int
	if errCount := r.Ctx.GetAny(key); errCount != nil {
		et = errCount.(int)
		if et >= count {
			return fmt.Errorf("exceed %d counts,url:%s", count, r.URL)
		}
	}
	r.Ctx.Put(key, et+1)
	return r.Retry()
}
