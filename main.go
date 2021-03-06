package main

import (
	"git.zgwit.com/zgwit/MyDTU/conf"
	"git.zgwit.com/zgwit/MyDTU/core"
	"git.zgwit.com/zgwit/MyDTU/flag"
	_ "git.zgwit.com/zgwit/MyDTU/protocol/modbus" //默认支持Modbus协议
	"git.zgwit.com/zgwit/MyDTU/web"
	"github.com/denisbrodbeck/machineid"
	"log"
)

// @title DTU manager API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.zgwit.com/support
// @contact.email jason@zgwit.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath /open
func main() {
	//解析参数
	if !flag.Parse() {
		return
	}

	id, err := machineid.ID()
	if err != nil {
		log.Println("获取ID错误：", err)
		return
	}
	log.Println("机器码：", id)

	//加载配置
	err = conf.Load()
	if err != nil {
		log.Println(err)
		return
	}

	//打印用户名密码，方便用户登录
	if !conf.Config.SysAdmin.Enable {
		for k, v := range conf.Config.BaseAuth.Users {
			log.Println("用户名：", k, "，密码：", v)
		}
		log.Println("系统使用HTTP简单认证，有泄露密码的风险，请采用安全的SysAdmin方式！！！")
	}

	//err = db.Open()
	//if err != nil {
	//	log.Println("数据库错误：", err)
	//	return
	//}

	//启动总线
	err = core.StartDBus(conf.Config.DBus.Addr)
	if err != nil {
		log.Println(err)
		return
	}

	//恢复之前的链接
	err = core.Recovery()
	if err != nil {
		log.Println(err)
		return
	}

	//启动Web服务
	web.Serve()
}
