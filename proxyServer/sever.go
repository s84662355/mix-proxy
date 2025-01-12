package proxyServer
import(
	  "mix-proxy/common"
	  	 "sync/atomic"
	  	 "context"
	  	 "net"
)

type Server struct{
	com 	*common.Common
    isStop atomic.Bool
    ctx context.Context
    listeners  []net.Listener 
    connBytePool sync.Pool
}