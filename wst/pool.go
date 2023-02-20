package main

import (
	"context"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"time"
	"github.com/koding/websocketproxy"
)

type ServerLink struct {
	Addr string
	Header http.Header
}

type Proxy struct {
	URL *url.URL
	WSProxy *websocketproxy.WebsocketProxy
	Score float64
}

type WebsocketProxyPool struct {
	ProxyPool []Proxy
}

func NewProxyPool() *WebsocketProxyPool {
	ctx := context.Background()
	w := &WebsocketProxyPool{}
	go w.score(ctx)
	return w
}

func (w *WebsocketProxyPool) PoolAddHandler(ll []ServerLink) {
	var proxyPool []Proxy
	for _, lotusLink := range ll {
		_url, err := url.Parse(lotusLink.Addr)
		if err != nil {
			Logger.Error("url parse err:", err)
		}
		proxyPool = append(proxyPool, Proxy{
			URL: _url,
			WSProxy: websocketproxy.NewProxy(_url),
			Score:   0,
		})
	}
	w.ProxyPool = proxyPool
}

func (w *WebsocketProxyPool) score(ctx context.Context) {
	timeTicker := time.NewTicker(time.Second * 2)
	for{
		Logger.Info("ProxyPool info:", w.ProxyPool)
		select {
		case <-timeTicker.C:
			for index, _ := range w.ProxyPool {
				w.ProxyPool[index].Score = rand.Float64()
			}
			sort.Slice(w.ProxyPool, func(i, j int) bool { return w.ProxyPool[i].Score > w.ProxyPool[j].Score })
		case <-ctx.Done():
			return
		}
	}
}

func (w *WebsocketProxyPool) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if len(w.ProxyPool) == 0 {
		Logger.Error("have no proxy to use")
	}
	w.ProxyPool[0].WSProxy.ServeHTTP(rw, req)
}
