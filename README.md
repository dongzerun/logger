# Logger
日志模块，支持日志按天|小时切割

##example
```
package main

import "github.com/dongzerun/logger"

func main(){
    //创建全局logger
    logger.InitLooger("INFO,WARNING,ERROR,DEBUG","/tmp","test.log",7,"20060102",0,0)
    logger.Warning("just test logger ok")
}
```

一般都用toml做配置文件，使用如下配置
##toml config
```
[golog]
level       = "INFO,WARNING,ERROR,DEBUG"
console     = 0
dir         = "/home/logs"
filename    = "test.log"
reserve_num = 7
suffix      = "20060102"
colorfull   = 0
```

##log struct
```
type LogConfig struct {
    Level      string `toml:"level"`
    Console    int    `toml:"console"`
    Dir        string `toml:"dir"`
    Filename   string `toml:"filename"`
    ReserveNum int    `toml:"reserve_num"`
    Suffix     string `toml:"suffix"`
    Colorfull  int    `toml:"colorfull"`
}
```