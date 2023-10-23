/* 一个简单的http自定义路由demo
 * 测试地址
 * http://127.0.0.1/
 * http://127.0.0.1/api
 */
package httptest

import (
	"log"
	"net/http"
	"testing"
)

type MyServer struct {
	handlers map[string]http.HandlerFunc
}

func TestHttp(t *testing.T) {

	ms := &MyServer{
		handlers: make(map[string]http.HandlerFunc),
	}
	hServer := &http.Server{
		Addr:    ":80",
		Handler: ms,
	}
	ms.handleFunc("/", myRoureIndex)
	ms.handleFunc("/api", myRoureApi)
	// 80端口提供http自定义函数的服务
	log.Println("Serving on http://0.0.0.0:80")
	log.Fatalln(hServer.ListenAndServe())
}
func (h *MyServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	p := request.URL.Path
	if handler, ok := h.handlers[p]; ok {
		handler(writer, request)
	}
}

func (h *MyServer) handleFunc(path string, handlerFunc http.HandlerFunc) {
	h.handlers[path] = handlerFunc
}
func myRoureIndex(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("ok find index"))
}
func myRoureApi(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("ok find api"))
}
