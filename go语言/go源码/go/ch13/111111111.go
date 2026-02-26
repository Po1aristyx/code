package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings" // 引入 strings 包用于 TrimPrefix
)

// --- 用于任务 2 的结构体 ---
type Server struct {
	ServerName string `json:"serverName"` // 添加 json tag 是好习惯
	ServerIP   string `json:"serverIP"`
}

type ServerSlice struct {
	Servers []Server `json:"servers"`
}

// --- 用于任务 1 的处理函数 ---
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling /hello/ request")
	// 更健壮地提取名字，去除前缀 /hello/
	name := strings.TrimPrefix(req.URL.Path, "/hello/")
	if name == "" {
		name = "Guest" // 如果没有提供名字，给个默认值
	}
	fmt.Fprintf(w, "hello %s", name) // 响应 hello Name
}

// --- 用于任务 3 的处理函数 ---
func loginHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling /login request")

	// 检查是否为 GET 请求
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析 GET 请求的查询参数
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")

	// 计算字符串长度
	lenUsername := len(username)
	lenPassword := len(password)

	// 设置响应内容类型为纯文本（如果需要 HTML 可以设置为 "text/html"）
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// 格式化输出到响应
	fmt.Fprintf(w, "您的登录信息：\n") // 使用 \n 换行
	fmt.Fprintf(w, "用户名： %s, 字符串长度为%d\n", username, lenUsername)
	fmt.Fprintf(w, "密码： %s，字符串长度为%d\n", password, lenPassword)
}

func main() {
	// --- 执行任务 2：解析 JSON ---
	fmt.Println("--- Task 2: Parsing JSON ---")
	var s ServerSlice
	// 使用反引号定义多行原始字符串，避免转义
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	} else {
		// %+v 会打印字段名和值，更清晰
		fmt.Printf("Parsed JSON data: %+v\n", s)
		// 也可以遍历打印
		// for _, server := range s.Servers {
		// 	fmt.Printf("Server Name: %s, Server IP: %s\n", server.ServerName, server.ServerIP)
		// }
	}
	fmt.Println("----------------------------")
	fmt.Println("Starting web server on http://localhost:9999")

	// --- 注册任务 1 和 任务 3 的路由处理函数 ---
	http.HandleFunc("/hello/", helloHandler) // 处理 /hello/ 开头的请求
	http.HandleFunc("/login", loginHandler)  // 处理 /login 请求

	// --- 启动 Web 服务器 (任务 1 和 3) ---
	// 设置监听的端口为 9999
	err = http.ListenAndServe("localhost:9999", nil) // 使用 nil 表示使用默认的 DefaultServeMux
	// 错误处理
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
