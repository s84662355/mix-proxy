package  common

import (
	"context"
	 "sync/atomic"
	 "sync"
	 	"mix-proxy/uitl/Map"
	
)
type ConnContext struct{
	ctx context.Context
	cancel func()
	count atomic.Int64 ///连接数量
}

func (c *Common) initConnMap() {
   c.connConcurrentMap = Map.New[*ConnContext]()
}


///增加连接 并且返回上下文
func (c *Common) AddConn(string k) context.Context  {
   value := c.connConcurrentMap.Upsert(k,nil, func(exist bool, valueInMap *ConnContext, newValue *ConnContext) context.Context {
   	  var value   *ConnContext = nil
      if !exist {
      		ctx, cancel := context.WithCancel(c.ctx)
            valueInMap = &ConnContext{
            	ctx:ctx,
            	cancel:cancel,
            } 
      }
      valueInMap.count.Add(1)
      return valueInMap
   })

   return value.ctx
}

//根据key和上下文 减少连接 
func (c *Common) ReduceConn(string k,ctx context.Context){
	c.connConcurrentMap.RemoveCb(k, func(key string, valueInMap *ConnContext, exists bool) bool {
      if exist  {
      	 if ctx ==  valueInMap.ctx {
      		if valueInMap.count.Add(-1) == 0 {
		      	valueInMap.cancel()
		        return true
      		}
      	 }
 
      }
      return false
   })
}


//根据key 删除所有连接
func (c *Common) RemoveAllConn(string k)  {
    valueInMap,exist:=  c.connConcurrentMap.	Pop(k)
  	if exist  {
	  	valueInMap.cancel()
	    return true
  	}
}

//根据key获取连接数量
func (c *Common) GetConnAmountByKey(string k) (int64,bool) {
	valueInMap ,ok := c.connConcurrentMap.Get(k)
	if ok {
		return  valueInMap.count.Load(),ok
	}
	return 0,ok
}


///并发获取所有连接数
func (c *Common)  GetConnAmount() int64 {
  var result   atomic.Int64  
  valueInMaps:= c.connConcurrentMap.Items() 
  count:=10
  connContextChan := make(chan *ConnContext,count) 
  var wg sync.WaitGroup
  wg.Add(count)
  for i := 0; i < count; i++ {
  	go func(){
  		defer wg.Done()
		defer recover()
		for  valueInMap := range connContextChan {
			result.Add(valueInMap.count.Load())
		}
  	}()
  } 
  for _, valueInMap := range valueInMaps {
  	connContextChan<-valueInMap
  }
  close(connContextChan)
  wg.Wait()
  return result.Load()
}
 
 