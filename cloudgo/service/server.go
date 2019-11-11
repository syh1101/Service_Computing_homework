package server
import (
    "github.com/go-martini/martini" 
)

func NewServer(port string) {   
    //创建一个martini实例
    m := martini.Classic()
    //请求处理器
    m.Get("/", func() string {
        return "hello world！"
    })
    //创建请求端口
    m.RunOnAddr(":"+port)   
}
