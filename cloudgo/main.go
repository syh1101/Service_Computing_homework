package main
import (
	"os"
	"github.com/syh1101/cloudgo/service"
	flag"github.com/spf13/pflag"
)

const(//默认端口
	PORT string = "8080"
)

func main(){
	port := os.Getenv("PORT")
	if len(port) == 0 {//监听不到端口，默认8080
		port = PORT
	}
	//解析端口
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening") 
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	//启动服务器
	server.NewServer(port) 
}
