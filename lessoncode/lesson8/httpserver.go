package main

import (
	"encoding/json"
	"net/http"
)

// User 表示用户信息
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Resp 表示响应信息
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// MyRoute 定义了路由的方法、路径、处理函数
type MyRoute struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// routes 存储所有的路由组
var routes []MyRoute

// MyHandler 实现http.Handler接口的ServeHTTP方法
type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range routes {
		if r.Method == route.Method && r.RequestURI == route.Path {
			route.Handler(w, r)
			return
		}
	}
}

func main() {
	// 1. 初始化路由列表
	routes = append(routes, []MyRoute{
		// 1.1 Handler 处理器函数，*http.Request：接收HTTP请求，http.ResponseWriter：发送HTTP响应
		{
			Method: "GET", Path: "/hello", Handler: func(w http.ResponseWriter, r *http.Request) {
				// 实例化User结构体
				user := &User{Name: "John", Age: 30}
				// 序列化为JSON格式
				userBytes, err := json.Marshal(user)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				// 通过http.ResponseWriter发送给客户端
				w.Write(userBytes)
			},
		},
		// 1.2 Handler:处理器函数
		{
			Method: "POST", Path: "/world", Handler: func(w http.ResponseWriter, r *http.Request) {
				// 实例化Resp结构体
				resp := &Resp{Code: 200, Msg: "success", Data: nil}
				respBytes, err := json.Marshal(resp)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(respBytes)
			},
		},
	}...)
	// 1.2 创建HTTP服务器
	srv := &http.Server{
		Addr:    ":9090",
		Handler: &MyHandler{},
	}
	// 1.3 启动HTTP服务器
	srv.ListenAndServe()
}

/*
	func main() {
		srv := &http.Server{
			Addr:    ":9090",
			Handler: &MyHandler{},
		}
		srv.ListenAndServe()
	}
	func main() {
		h := func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		}
		http.HandleFunc("/", h)
		http.ListenAndServe(":9090", nil)
	}
*/
