package main
import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"
	"mix-proxy/log"
	 "io"
	"runtime"

)


///获取程序的内存 协程数 以及各种数据
func  init() {



http.HandleFunc("/", getRoot)

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Panic(err)
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
		defer log.DRecover("http server ")
		http.HandleFunc("/runtime", func(w http.ResponseWriter, r *http.Request){
			//runtime.NumCPU()
			//runtime.NumGoroutine()
			//memStatus := runtime.MemStats{}
	    //runtime.ReadMemStats(&memStatus)
			resultString:=fmt
			io.WriteString(w, "This is my website!\n")
			 
		})
		panic(http.Serve(ln, nil))
	 
	}()
}

