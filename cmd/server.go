package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/taoshihan1991/miaosha/controller"
	"github.com/taoshihan1991/miaosha/setting"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	port   string
	daemon bool
)
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "example:miaosha server port 8082",
	Example: "go-fly server -c config/",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	serverCmd.PersistentFlags().StringVarP(&port, "port", "p", "8082", "监听端口号")
	serverCmd.PersistentFlags().BoolVarP(&daemon, "daemon", "d", false, "是否为守护进程模式")
}
func run() {
	if daemon == true {
		if os.Getppid() != 1 {
			// 将命令行参数中执行文件路径转换成可用路径
			filePath, _ := filepath.Abs(os.Args[0])
			cmd := exec.Command(filePath, os.Args[1:]...)
			// 将其他命令传入生成出的进程
			cmd.Stdin = os.Stdin // 给新进程设置文件描述符，可以重定向到文件中
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Start() // 开始执行新进程，不等待新进程退出
			os.Exit(0)
		}
	}
	baseServer := "0.0.0.0:" + port
	log.Println("start api server...\r\nurl：http://" + baseServer)
	engine := gin.Default()
	engine.Static("/static", "./front/static")

	//配置文件
	setting.GetConfigIni("config.ini")
	setting.GetRedisConfig()

	//性能监控
	pprof.Register(engine)
	initRouter(engine)
	initBackendService()

	engine.Run(baseServer)
}
func initRouter(engine *gin.Engine) {
	engine.GET("/", controller.PageIndex)
	engine.GET("/product", controller.GetProduct)
	engine.GET("/buy", controller.GetKillUrl)
	engine.GET("/timestamp", controller.GetTimestamp)
	engine.GET("/orders", controller.GetOrders)
	engine.GET("/userinfo", controller.GetUserInfo)
	engine.POST("/userinfo", controller.PostUserInfo)
	engine.GET("/seckill/:token", controller.PostSale)
}
func initBackendService() {
	go controller.GetProductQueueToOrder()
}
