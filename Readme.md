# 一个简单的Socks5服务端

#### 获得这个库
> go get github.com/zhb255/ss5-simple 

#### 简单的用法

> `./hello.html`的定义，在`example/hello.html`中

```go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Socks5 Server response</title>
</head>
<body>
hello world
</body>
</html>
```

`main.go`

```go
package main

import (
	"github.com/zbh255/ss5-simple/handler"
	snet "github.com/zbh255/ss5-simple/net"
	"io/ioutil"
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
		go func() {
			err := simpleHandlerConnection(ssc)
			if err != nil {
				log.Printf("[Error]: %v", err)
			}
		}()
	}
}

func simpleHandlerConnection(conn snet.SSConn) error {
	rep,err := ioutil.ReadFile("./hello.html")
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.RegisterConnectHandler(handler.Comment, func(request []byte) ([]byte, error) {
		return rep, nil
	})
	return conn.Handler()
}

func main() {
    SimpleServer()
}
```
#### 测试

```shell
go run main.go
export all_proxy=socks5://127.0.0.1:1080
curl google.com
```

`OutPut`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Socks5 Server response</title>
</head>
<body>
hello world
</body>
</html>
```

---

> 更多的例子可以在`example`中找到

