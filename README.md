# Logger
日志模块，支持日志按天|小时切割

##example
package main

import "github.com/dongzerun/logger"

func main(){
    //创建全局logger
    logger.InitLooger("INFO,WARNING,ERROR,DEBUG","/tmp","test.log",7,"20060102",0,0)
    logger.Warning("just test logger ok")
}
