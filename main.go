package main

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"
	"mix-proxy/log"
	 


)

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	go func(){
		defer recover()
		for {
			log.Info("端口:%d", ln.Addr().(*net.TCPAddr).Port)
			select{
			case <-time.After(1*time.Hour):
			}
		}
	}()
	go func() {
		panic(http.Serve(ln, nil))
	}()
}


