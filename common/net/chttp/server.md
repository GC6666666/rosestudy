# 1.听课笔记

```go
package chttp
// 基于gin框架进行封装
// 启动http server依赖我们的http框架，是一种通用能力，需要和业务进行解耦
// 放在common目录下主要是为了和业务进行解耦
import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"sync/atomic"
)

type Server struct {
	// 内嵌，继承RouterGroup的所有方法
	*gin.RouterGroup
    
	conf   *Config
	// 用来存储http server，方便获取
	server atomic.Value
	// gin的Engine
	engine *gin.Engine
}
// NewServer
// 使用http框架的时候，先调用NewServer创建Server对象
// 同时初始化conf和engine
// 嵌套的RouterGroup被初始化为gin的RouterGroup
func NewServer(conf *Config) *Server {
	s := &Server{
		conf:   conf,
		engine: gin.New(),
	}

	s.RouterGroup = &s.engine.RouterGroup

	return s
}

func (s *Server) Start() error {
	lis, err := net.Listen(s.conf.Network, s.conf.Address)
	if err != nil {
		return err
	}

	hs := &http.Server{
		Handler:      s.engine,
		ReadTimeout:  s.conf.ReadTimeOut,
		WriteTimeout: s.conf.WriteTimeOut,
	}
	s.server.Store(hs)

	return hs.Serve(lis)
}

func (s *Server) getServer() *http.Server {
	// atomic.Value存的是interface，进行类型断言，判断是不是http.Server，是的话返回true
	server, ok := s.server.Load().(*http.Server)
	if !ok {
		return nil
	}

	return server
}

func (s *Server) ShutDown(ctx context.Context) error {
	server := s.getServer()
	if server == nil {
		return fmt.Errorf("[chttp] server is nil")
	}

	return server.Shutdown(ctx)
}

```

# 2.代码知识点补充讲解

## 2.1 Server和 NewServer

### 2.1.1 Server

定义了一个简单的 HTTP 服务器，使用了 Gin 框架来处理路由和请求，并利用 Go 的 `atomic` 包来安全地处理并发操作。

1. `*gin.RouterGroup`:
    - 这个字段实际上是一个嵌入字段，代表 `Server` 结构体继承了 `gin.RouterGroup` 的所有方法。`gin.RouterGroup` 是 Gin 框架中用于路由分组的类型，允许你在路由中创建分组和中间件。
    - `gin.RouterGroup` 是 Gin 框架中用于分组路由的功能，可以帮助你组织和管理具有共同路径前缀的路由，常用于构建具有层次结构的 API。此外，`RouterGroup` 允许为一组路由添加共享的中间件，非常适合处理类似的路由功能，如认证、日志记录等。
    - 具体使用见3.扩展知识-gin

2. `conf *Config`:
    - 这个字段存储了服务器的配置信息，比如监听的网络和地址。是一个指向 `Config` 类型的指针。
    - `conf *Config` 为什么要加 `*`：

    在 Go 语言中，`*` 符号用于创建指针。在这个例子中，`conf *Config` 表示 `conf` 是一个指向 `Config` 类型的指针。使用指针的原因通常是为了节省内存和提高效率（如何提高的效率），特别是在传递大型结构或对象时。通过传递一个指向原始数据的指针，而不是复制整个数据结构，可以减少内存使用并避免不必要的复制成本。

3. `server atomic.Value`:
    - 这个字段使用 `atomic.Value` 来存储和管理 `http.Server` 的实例。
    - `atomic.Value` 提供了一个能够存储任意类型值的线程安全的容器。在多线程环境中，它用于保持对值的读写操作的原子性，确保并发操作的安全。（？）
    - **代码示例**：
```go
package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func main() {
    var val atomic.Value
    val.Store(42)

    go func() {
        time.Sleep(time.Second)
        val.Store(100)
    }()

    fmt.Println(val.Load()) // 输出 42
    time.Sleep(2 * time.Second)
    fmt.Println(val.Load()) // 输出 100
}
```
(自认为对这一部分的理解还有待加强)

4. `engine *gin.Engine`:
    - `*gin.Engine` 是 Gin 框架的核心，负责路由的调度和中间件的处理。它是一个 HTTP 请求的处理引擎。
    - `*` 表示这是一个指向 `gin.Engine` 类型的指针，通常这样使用是为了能够通过方法共享和修改它的状态。
    - 代码中，`*gin.Engine` 作为一个字段存在而不是嵌入，我认为这是为了更好的初始化嵌套的GroupRouter,也是为了更好的调用gin框架的API

### 2.1.2 NewSrever
这部分就是一个初始化，没有什么特别需要讲的

# 3.扩展知识-gin

## 3.1 gin框架最基本使用

下面是一个简化的 Gin 框架示例代码：创建一个服务器，设置几个简单的路由，并启动服务器。

我把以下程序放在了mgin.go文件中

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建一个Gin路由器实例
    router := gin.Default()

    // 设置一个GET请求的路由
    router.GET("/welcome", func(c *gin.Context) {
        // 向客户端发送字符串响应
        c.String(http.StatusOK, "Welcome to Gin!")
    })

    // 启动服务器，在8080端口监听
    router.Run(":8080")
}
```

1. **导入 Gin 包和 net/http 包：**
   - `import "github.com/gin-gonic/gin"` 导入 Gin 框架，允许我们使用 Gin 的各种功能。
   - `import "net/http"` 导入 Go 的 HTTP 包，以使用 HTTP 状态码等。

2. **创建一个路由器实例：**
   - `router := gin.Default()` 通过调用 `gin.Default()` 创建一个带有默认中间件（日志和恢复中间件）的路由器。

3. **设置路由：**
   - `router.GET("/welcome", func(c *gin.Context) {...})` 创建一个处理 GET 请求的路由 `/welcome`。当访问这个路由时，服务器会执行括号中的匿名函数，使用 `c.String()` 方法向请求者发送一个欢迎消息。

4. **启动服务器：**
   - `router.Run(":8080")` 启动服务器并设置监听端口为 8080。这意味着服务器会处理发送到这个端口的所有网络请求。

5. **运行并查看：**
   - 在终端中输入go run .\mgin.go
   - 在网页中输入网址：http://localhost:8080/welcome， 可以看到发送的消息"Welcome to Gin!"

## 3.2 加入RouterGroup
使用 `RouterGroup` 是一种组织路由的好方法，特别是在构建有多个相关路由的复杂应用时。它可以帮助你保持代码的清洁和结构化。

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建一个Gin路由器实例
    router := gin.Default()

    // 创建一个路由组
    apiGroup := router.Group("/api")
    {
        // 在/api组中设置一个GET请求的路由
        apiGroup.GET("/welcome", func(c *gin.Context) {
            // 向客户端发送字符串响应
            c.String(http.StatusOK, "Welcome to the API!")
        })

        // 在/api组中设置另一个GET请求的路由
        apiGroup.GET("/greet", func(c *gin.Context) {
            // 接收一个名为"name"的查询参数
            name := c.DefaultQuery("name", "Guest")
            // 使用查询参数在响应中发送个性化的问候
            c.String(http.StatusOK, "Hello %s!", name)
        })
    }

    // 启动服务器，在8080端口监听
    router.Run(":8080")
}
```

创建了一个名为 `apiGroup` 的 `RouterGroup`。这个组包括两个路由：`/api/welcome` 和 `/api/greet`。

- **创建路由组：**
```go
apiGroup := router.Group("/api")
```
  这行代码创建了一个新的路由组 `apiGroup`，所有属于这个组的路由都会自动带有 `/api` 前缀。这有助于你将路由逻辑按功能或 API 版本进行分组，从而使路由结构更清晰。

- **定义路由组内的路由：**
   - `/api/welcome` 路由返回一个欢迎消息。
   - `/api/greet` 路由提供一个简单的问候功能，客户可以通过查询参数 `name` 来接收个性化的问候。如果没有提供 `name` 参数，它会默认为 "Guest"。

- **运行和测试:**
通过以下 URL 访问新的路由：
  - 访问 `http://localhost:8080/api/welcome` 应该返回 "Welcome to the API!"。
  - 访问 `http://localhost:8080/api/greet` 将返回 "Hello Guest!"。
  - 访问 `http://localhost:8080/api/greet?name=John` 将返回 "Hello John!"。

这种方式可以根据功能模块或版本号组织 API，提高了代码的可维护性和扩展性。

## 3.3 关于engine

在 Gin 框架中，`gin.Engine` 是一个定义的结构体，它实际上是整个 Gin 框架的核心。`gin.Engine` 承载了路由管理、中间件管理、模板渲染、组设置等多种功能，是处理 HTTP 请求和响应的主要接口。

### 3.3.1 `gin.Engine` 结构体的主要功能包括：

1. **路由管理：**
   `gin.Engine` 包含一个路由器，可以注册到不同的 URL 路径和对应的处理函数。它支持各种 HTTP 方法，如 GET、POST、PUT、DELETE 等。

2. **中间件支持：**
   `gin.Engine` 允许你在全局或特定路由上添加中间件。中间件是可以处理请求或响应，或者在请求处理前后执行某些操作的函数。

3. **群组路由：**
   RouterGroup:使用 `gin.Engine`，你可以创建路由群组，这有助于组织和管理具有共同URL前缀的路由集合。路由群组可以继承或拥有它们自己的中间件。

4. **错误管理：**
   `gin.Engine` 提供了错误管理功能，可以捕获和处理请求过程中的错误。

5. **数据绑定和验证：**
   `gin.Engine` 支持将请求中的数据（如 JSON、表单数据等）绑定到 Go 的结构体中，并支持数据验证。

6. **渲染支持：**
   它还提供了响应数据的渲染支持，如 JSON、XML 和 HTML 模板渲染。

### 3.3.2 代码示例 1：
在创建一个 `gin.Engine` 实例时，通常会这样做：

```go
router := gin.Default()
```

或者如果你想自己管理中间件，可以直接使用：

```go
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
```
gin.Default()里面调用了gin.New();在调用完gin.New()得到Engine 实例后,还调用了engine.Use(Logger(), Recovery());gin.Default()获取到的Engine 实例集成了Logger 和 Recovery 中间件

以上如何使用 `gin.Engine` 来创建一个 HTTP 服务器，并通过添加中间件和路由来配置其行为。

总之，`gin.Engine` 是 Gin 框架的核心，提供了创建和管理 Web 应用所需的大部分功能和工具。它通过一个简洁的 API 使得创建快速且高效的 HTTP 服务器变得容易。

### 3.3.2 代码示例 2

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // 创建一个没有任何默认中间件的新引擎
    engine := gin.New()

    // 为引擎添加中间件
    engine.Use(gin.Logger())  // 添加日志中间件，记录所有请求
    engine.Use(gin.Recovery())  // 添加恢复中间件，处理所有panic，防止崩溃

    // 创建一个路由组
    apiGroup := engine.Group("/api")
    {
        // 在/api组中设置一个GET请求的路由
        apiGroup.GET("/welcome", func(c *gin.Context) {
            // 向客户端发送字符串响应
            c.String(http.StatusOK, "Welcome to the API!")
        })

        // 在/api组中设置另一个GET请求的路由
        apiGroup.GET("/greet", func(c *gin.Context) {
            // 接收一个名为"name"的查询参数
            name := c.DefaultQuery("name", "Guest")
            // 使用查询参数在响应中发送个性化的问候
            c.String(http.StatusOK, "Hello %s!", name)
        })
    }

    // 启动服务器，在8080端口监听
    engine.Run(":8080")
}
```

以上版本中，`*gin.Engine` 被显式创建和配置了：
- **创建引擎：**
```go
engine := gin.New()
```
这行代码创建了一个新的 Gin 引擎实例，但没有包含任何默认的中间件。这使得初始化过程更加透明，你可以自定义添加所需的中间件。

- **添加中间件：**
```go
engine.Use(gin.Logger())
engine.Use(gin.Recovery())
```
这两行分别添加了日志和恢复中间件。日志中间件用于记录每个请求的详细信息，而恢复中间件可以捕获处理过程中的任何 panic，保证服务不会因为异常而中断。

通过这种方式，你可以看到 `*gin.Engine` 的创建和配置过程，以及如何手动添加中间件。这种做法可以在需要细粒度控制时提供更大的灵活性，例如在需要精确控制哪些中间件被包含在你的应用中时。

## 3.4 关于RouterGroup

`RouterGroup`也是 Gin 框架中定义的一个结构体,主要用于组织和管理具有相同前缀的一组路由。这使得路由管理更加模块化和可维护。`RouterGroup` 提供了一种方便的方式来分组相关的路由，并且可以应用特定的中间件到这个组中的所有路由。

`gin.Engine` 实际上内嵌了RouterGroup。在 Gin 框架中，gin.Engine 本身扩展自 RouterGroup，这意味着 gin.Engine 继承了 RouterGroup 的所有方法和属性

### `RouterGroup` 的主要特性包括：

1. **路由组织：**
   `RouterGroup` 允许你将相关的路由聚集在一起，这些路由共享相同的URL前缀。例如，所有关于用户的路由可能都在一个名为 `/users` 的 `RouterGroup` 下。

2. **中间件应用：**
   你可以为特定的 `RouterGroup` 添加中间件，这些中间件只会影响该组内的路由。这是一种非常有效的方式来处理那些只需要在特定路由上运行的逻辑。

3. **路由嵌套：**
   `RouterGroup` 支持嵌套，这意味着你可以在一个 `RouterGroup` 内部创建另一个 `RouterGroup`。这提供了更精细的组织结构，允许构建复杂的路由层级结构。

4. **继承属性：**
   子 `RouterGroup` 会继承父 `RouterGroup` 的中间件和路径前缀，同时也可以添加自己的中间件或覆盖继承来的中间件。

# 4.扩展知识- Start 和 ShutDown 方法的 API 解析

## 4.1 `net.Listen`

`net.Listen` 是 Go 语言标准库中 `net` 包的一个函数，用于在指定的网络地址上监听来自客户端的连接请求。这是网络编程中一个基础且关键的功能，允许服务器端应用程序开始接收客户端发起的连接。

**函数签名：**
```go
func Listen(network, address string) (Listener, error)
```

- `network`: 指定网络类型，常用的类型有 "tcp", "tcp4" (只用IPv4), "tcp6" (只用IPv6), "unix" (UNIX域套接字), "unixpacket" 等。
- `address`: 指定监听地址。对于 TCP/UDP 网络，这通常是 IP 地址和端口号（例如 "192.168.1.1:8080" 或 ":8080" 表示监听所有接口的 8080 端口）。

**用法示例：**
```go
listener, err := net.Listen("tcp", ":8080")
if err != nil {
    log.Fatalf("Failed to listen: %v", err)
}
defer listener.Close()
```

## 4.2 `http.Server`

`http.Server` 结构体定义了一个 HTTP 服务器的所有配置。服务器的配置包括如何处理请求、监听的端口、超时设置等。

**重要字段：**
- `Handler`: `http.Handler` 接口的实例，负责处理所有的 HTTP 请求。？？
- `ReadTimeout` 和 `WriteTimeout`: 控制服务器读写操作的超时时间。
- `MaxHeaderBytes`: 控制最大的请求头大小。

**用法示例：**
```go
server := &http.Server{
    Handler:      myHandler, // 实现了 http.Handler 接口的实例
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second, // 新增空闲连接的超时时间
}
```

## 4.3 `hs.Serve`

此方法是 `http.Server` 的一个成员方法，负责将服务器绑定到一个监听器并开始处理请求。这是一个阻塞调用，服务器将持续运行，直到发生错误或被停止。

**函数签名：**
```go
func (srv *Server) Serve(l net.Listener) error
```

**用法示例：**
```go
err := server.Serve(listener)
if err != http.ErrServerClosed {
    log.Fatalf("Server failed: %v", err)
}
```

## 4.4 `server.Shutdown`

`server.Shutdown` 方法用于优雅地关闭服务器。这意味着服务器将停止接受新的请求，但会等待已经建立的请求处理完毕。

**函数签名：**
```go
func (srv *Server) Shutdown(ctx context.Context) error
```

- `ctx`: `context.Context` 实例，可以用于设置超时或取消关闭过程。

**用法示例：**
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := server.Shutdown(ctx); err != nil {
    log.Fatalf("Server shutdown failed: %v", err)
}
```

## 4.5 结合使用

将上述组件结合起来，可以创建一个健壮的 HTTP 服务器，它能够接受连接、处理请求，并优雅地处理关闭事件。这不仅增强了服务的可用性，还提高了服务的稳定性，能够在维护或升级时减少对客户端的影响。
