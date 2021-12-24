package main

import (
	"github.com/zbh255/ss5-simple/handler"
	snet "github.com/zbh255/ss5-simple/net"
	"log"
	"net"
)

func SimpleServer() {
	listener, err := net.Listen("tcp", "localhost:1080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	socks5Server := snet.NewNoAuthServer(listener)
	for {
		ssc, err := socks5Server.Connection()
		if err != nil {
			log.Print(err.Error())
			continue
		}
		go simpleHandlerConnection(ssc)
	}
}

func simpleHandlerConnection(conn snet.SSConn) {
	defer conn.Close()
	conn.RegisterConnectHandler(handler.Comment, func(request []byte) ([]byte, error) {
		rep := []byte("<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <title>安全入口校验失败</title>\n</head>\n<body>\n    <h1>请使用正确的入口登录面板</h1>\n    <p><b>错误原因：</b>当前新安装的已经开启了安全入口登录，新装机器都会随机一个8位字符的安全入口名称，亦可以在面板设置处修改，如您没记录或不记得了，可以使用以下方式解决</p>\n    <p><b>解决方法：</b>在SSH终端输入以下一种命令来解决</p>\n    <p>1.查看面板入口：/etc/init.d/bt default</p>\n    <p>2.关闭安全入口：rm -f /www/server/panel/data/admin_path.pl</p>\n    <p style=\"color:red;\">注意：【关闭安全入口】将使您的面板登录地址被直接暴露在互联网上，非常危险，请谨慎操作</p>\n</body>\n</html>")
		return rep, nil
	})
	err := conn.Handler()
	if err != nil {
		log.Print(err.Error())
	}
}
