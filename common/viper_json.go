package  common
import(	
	"mix-proxy/log"
	 "strings"
	 "strconv"
)
type Config struct { 
    AppId string 
    Secret string
    Host Host 
 	ListenPort string

}
 
func (c *Common) GetConfig() *Config{
 	return  c.viperConfig
}

func (c *Common) initConfig(){
	c.viperConfig = &Config{}
	c.viperTool = viper.New()
	c.viperTool.AddConfigPath("./config")     //设置读取的文件路径 
	c.viperTool.SetConfigName("config.json")　　　　  //设置读取的文件名 
	c.viperTool.SetConfigType("json")　　　　  //设置文件的类型　  
    //配置读取　　
    if err := config.ReadInConfig(); err != nil {
        panic(err)
    }　　 

    //直接反序列化为Struct
    var configjson Config
    if err :=config.Unmarshal(&c.viperConfig);err !=nil{
    	log.Panic("启动读取配置失败",err)
        panic(err)
    }
 
	c.viperTool.WatchConfig()  //启动监听
	c.viperTool.OnConfigChange(func(e fsnotify.Event) {   //注册文件监控的回调函数
	    if err :=config.Unmarshal(&c.viperConfig);err !=nil{
	    	log.Error("配置发生变化 读取失败",err)
	    }  
	})  
}

func (c *Common) ListenPort() []int {

	if  strings.Contains(	c.viperConfig.ListenPort, "-") {
 		result := strings.Split(c.viperConfig.ListenPort, "-")
 		resultArr:=[]int{}
 		if len(result) !=2 {
 		 	return resultArr
 		}

	    num1, err := strconv.Atoi(		result[0])
	    if err != nil{
	    	return resultArr
	    }

		num2, err := strconv.Atoi(		result[1])
    	if err != nil{
		   return resultArr
		}

	    for num1<=num2 {
	    	resultArr=append(resultArr,num1)
	    	num1++
	    	
	    }
 		return resultArr

	}else if  strings.Contains(	c.viperConfig.ListenPort, ","){
		result := strings.Split(c.viperConfig.ListenPort, ",")
		resultArr:=[]int{}
		if len(result) == 0 {
			return resultArr
		}

		for _, vvv := range result {
			port, _ := strconv.Atoi(vvv)
			if port >0 {
				resultArr=append(resultArr,port)
			}
		}

		return  resultArr

	}

	return []int{}


}