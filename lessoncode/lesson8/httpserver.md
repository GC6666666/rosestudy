### 1. 关于routes的创建流程，已经写在了代码的备注中

### 2. 关于创建并启动HTTP服务器

也就是以下代码的逻辑讲解：
```go
srv := &http.Server{
	Addr:    ":9090",
	Handler: &MyHandler{},
}
srv.ListenAndServe()
```
创建并启动一个HTTP服务器的逻辑可以分解为以下几个步骤：

1. **定义服务器地址和处理器**：
   ```go
   srv := &http.Server{
       Addr:    ":9090",
       Handler: &MyHandler{},
   }
   ```
    - `Addr` 字段指定了服务器监听的网络地址。在这个例子中，它被设置为 `":9090"`，意味着服务器将在本地的9090端口上监听所有网络接口（因为地址为空字符串）。

    - `Handler` 字段是 `http.Handler` 接口的一个实现，它定义了如何处理传入的HTTP请求。在这个例子中，`Handler` 被设置为 `&MyHandler{}`，即 `MyHandler` 类型的实例的地址。

2. **启动服务器**：
   ```go
   srv.ListenAndServe()
   ```
    - `ListenAndServe` 方法首先会调用 `Listen` 方法，该方法根据 `Addr` 字段的值绑定到指定的地址和端口，然后开始监听传入的TCP连接。

    - 一旦 `Listen` 成功，`Serve` 方法会启动。这个方法会一直运行，直到服务器因为某些原因关闭。它接受传入的连接，并将每个连接的HTTP请求交给 `Handler` 来处理。

3. **处理请求**：
    - 当服务器接受到一个HTTP请求时，它会为这个请求创建一个 `http.ResponseWriter` 对象和一个 `*http.Request` 对象。

    - 然后，它会调用之前设置的 `Handler` 的 `ServeHTTP` 方法，将 `ResponseWriter` 和 `Request` 对象作为参数传递。

4. **自定义路由和响应**：
    - 在 `MyHandler` 的 `ServeHTTP` 方法中，服务器会遍历所有的自定义路由（`routes`），寻找匹配当前请求方法和路径的路由。

    - 一旦找到匹配的路由，就会调用该路由的处理器函数（`Handler`），该函数负责生成响应并写入 `ResponseWriter`。

5. **响应客户端**：
    - 处理器函数通过 `ResponseWriter` 对象向客户端发送响应。这包括设置HTTP状态码（如200 OK）、编写响应头和发送响应体。

6. **服务器运行**：
    - `ListenAndServe` 方法会阻塞调用它的goroutine，直到服务器关闭。服务器关闭可以是因为接收到了关闭信号、发生了错误、或者调用了 `http.Server` 的 `Shutdown` 方法。

7. **优雅关闭**：
    - 如果需要优雅地关闭服务器（例如，在接收到中断信号时），可以调用 `srv.Shutdown(ctx)` 方法，其中 `ctx` 是一个 `context.Context` 对象。这会停止接收新的HTTP请求，并且允许已经接收的请求完成处理。

通过以上步骤，Go语言中的HTTP服务器就可以处理传入的HTTP请求，并且可以根据自定义的路由逻辑来生成和发送响应。