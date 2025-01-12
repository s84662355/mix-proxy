package proxyServer
import(
	"mix-proxy/log"
		 "sync"
		 "fmt"
		 	 "sync/atomic"
		 	 "net"
)
	 
func (s * Server) initTcp() []net.Listener {
	ports:= s.com.ListenPort()
	if len(ports) ==0{
		log.Error("监听端口数量为0")
		return
	}
 	var wg sync.WaitGroup
 	count:=len(ports)
  	wg.Add(count)
  	errChan := make(chan err,1)
  	defer close(errChan)
  	var isStop atomic.Bool
  	isStop.Store(false)
    listenerAtt:= make([]net.Listener,count)
	for i, port := range ports {
  		go func(i,port int){
  			defer log.Recover("initTcp Listen ")
  			defer wg.Done()
  			if !isStop.Load(){
  				return
  			}
  			listen,err := net.Listen("tcp",fmt.Sprint("0.0.0.0:",port))
  			if err!=nil{
  				if isStop.CompareAndSwap(false,true){
					errChan<-err
					return
  				}
  				return
  			}  
  			listenerAtt[i]=listen
  		}(i,port)
  		
  	}
  	wg.Wait()
  	select {
  	case err:=<-errChan:
  		//抛出异常
  		log.Panic("监听端口失败",err)
  		return
  	default:
  	}
	s.listeners = listenerAtt
	s.connBytePool = sync.Pool{
   		New: func() interface{} {
        	return make([]byte,1024)
    	},
	}

}


func (s * Server)  accept( ){
	///每个listen 开启5个Accept 可以加快获取在监听队列获取连接的速度，减少出现队列溢出的情况
	for i := 0; i < 5; i++ {
	  for _, listen := range s.listeners  {
	  	  go func(listen net.Listener){
	  	  	defer log.Recover("listen recover")
			for !s.isStop.Load() {         
				conn,err := listen.Accept()                    
				if err != nil{
					log.Error("listen fail ",listen,err)
					continue	                                    
				}
				go s.handleTcpConn(conn)
			}
	  	  }(listen)
	   
	  }
	}
} 

func (s * Server)  handleTcpConn(conn net.Conn) {
  	defer log.DRecover("handleTcpConn ")
  	defer conn.Close()
 	readByte:=s.connBytePool.Get().([]byte)
    
    ///这里要设置读取的超时
 	n,err:=conn.Read(readByte)
 	if err!=nil{
 		log.Error("handleTcpConn first read ",err)
 		return
 	}

 	if n == 0{
 		log.Error("handleTcpConn first read  byte == 0")
 		return
 	}
    
    
 	if n < 10 {
 	 //     socks5 proxy

 	}else {
			// http proxy
 	}
    


  





 	 
}
