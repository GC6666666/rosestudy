## 1. Go语言中如何管理数据库连接

### 1.1 包导入

```go
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
```
- **`"database/sql"`**: 这个包提供了与SQL数据库交互的接口。它定义了一些用于处理数据库连接和执行SQL命令的基础工具。
- **`_ "github.com/go-sql-driver/mysql"`**: 这是MySQL的驱动程序，用来实现`database/sql`定义的接口。`_`的使用表示我们只是想引入该包，从而调用其初始化函数，该函数会自动注册MySQL驱动到`database/sql`包中。这样，当我们使用`sql.Open("mysql", ...)`时，`database/sql`包能够找到并使用MySQL驱动。
- **`"time"`**: 这个包提供了时间相关的功能，如计时、时间差计算等。在这里，它主要用于设置数据库连接的超时时间。

### 1.2 配置结构体（Config）

```go
type Config struct {
	DSN         string        `yaml:"dsn"`
	Active      int           `yaml:"active"`
	Idle        int           `yaml:"idle"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}
```
这个结构体定义了数据库连接的配置项：
- **`DSN`**: 数据源名称，这是一个字符串，包含连接数据库所需的全部信息（如用户名、密码、服务器地址、数据库名等）。
- **`Active`**: 最大的活动（打开的）连接数。这限制了同时打开的数据库连接的数量。
- **`Idle`**: 最大的空闲连接数。这是连接池中可以保持打开但未使用的连接的数量。
- **`IdleTimeout`**: 连接可以保持空闲状态的最长时间，之后应被关闭释放资源。

### 1.3 数据库结构体（DB）和构造函数

```go
type DB struct {
	*sql.DB
}

func NewDB(conf *Config) *DB {
	if conf == nil {
		panic("conf cannot be nil")
	}

	d, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(conf.Active)
	d.SetMaxIdleConns(conf.Idle)
	d.SetConnMaxIdleTime(conf.IdleTimeout)

	return &DB{
		DB: d,
	}
}
```
- **`DB` 结构体**: 这是对 `sql.DB` 的一个简单封装。`sql.DB` 代表了一个数据库的长连接池。它并不是一个单一的连接，而是可能有多个连接的集合。
- **`NewDB` 函数**: 这是一个构造函数，用于创建`DB`实例。
    - **参数检查**: 如果传入的配置是`nil`，使用`panic`立即终止程序，因为没有配置就无法连接数据库。
    - **打开连接**: `sql.Open("mysql", conf.DSN)`尝试使用给定的DSN字符串（已在`Config`中定义）打开数据库连接。`"mysql"`指定使用MySQL驱动。
    - **设置连接池参数**: 使用`SetMaxOpenConns`, `SetMaxIdleConns`, 和 `SetConnMaxIdleTime`方法设置连接池的相关属性，这些属性控制了连接的数量和生命周期。
    - **返回**: 创建一个`DB`实例，包装了`sql.DB`对象。


### 1.4 扩展知识

关于：
```go
return &DB{
	DB: d,
}
```

#### 1.4.1 结构体封装和返回

定义一个结构体来封装一些数据或功能时，通常会这样写：

```go
type DB struct {
	*sql.DB
}
```

- 这里，`DB` 是一个新的结构体类型，它内嵌了 `*sql.DB`。
- `sql.DB` 本身是一个结构体类型，代表着一个持久的数据库连接池，而不是单个的数据库连接。这种内嵌方式使得 `DB` 类型的实例自动继承了 `*sql.DB` 的所有方法，因此可以直接调用，如 `db.Query()`、`db.Close()` 等。

#### 1.4.2 构造新实例

```go
return &DB{
	DB: d,
}
```

- `DB: d`：这是一个结构体字面量的初始化表达式。`DB` 是 `DB` 结构体中嵌入的 `*sql.DB` 类型字段的名称，而 `d` 是一个 `*sql.DB` 实例。这个赋值表明，新创建的 `DB` 结构体中的 `DB` 字段（即嵌入的 `*sql.DB`）将指向 `d`。
- `&DB{...}`：`&` 操作符创建一个新的结构体实例，并返回这个实例的地址，即一个指向 `DB` 的指针。在Go中，使用指针可以避免数据的复制，提高效率，并允许直接修改结构体内部的数据。

#### 1.4.3 总结

这种写法的目的是：
1. **封装**：通过自定义的 `DB` 结构体封装 `*sql.DB`，可以在这个封装层添加额外的功能或管理逻辑，例如错误处理、日志记录、性能监控等，而不影响原有的数据库操作逻辑。
2. **扩展性**：在未来，如果需要添加更多的字段或方法到 `DB` 结构体，可以轻松实现，而不需要修改大量的代码。

通过返回指向这个新实例的指针，调用者可以直接使用这个新的 `DB` 实例，利用其提供的所有方法和功能，这些方法和功能都是安全地封装在 `DB` 类型中的。这样的设计模式在Go语言中非常常见，用于提供清晰、可维护的代码结构。