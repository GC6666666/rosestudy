# common/net/chttp/config.go

internal/conf/conf.go就是嵌套了这里的Config

# 将YAML文件中的键映射到结构体的对应字段

定义了一个名为`Config`的结构体，它被设计来存储网络配置信息。这种结构体在处理配置文件（如YAML）时很常见，尤其是在网络应用或服务中。这里详细解释每个部分：

```go
type Config struct {
    Network      string        `yaml:"network"`
    Address      string        `yaml:"address"`
    ReadTimeOut  time.Duration `yaml:"readTimeOut"`
    WriteTimeOut time.Duration `yaml:"writeTimeOut"`
}
```

1. **字段**：
    - `Network`：一个类型为`string`的字段，用于存储网络类型，比如`tcp`、`udp`等。
    - `Address`：一个类型为`string`的字段，通常用于存储IP地址和端口号，例如`"127.0.0.1:8080"`。
    - `ReadTimeOut` 和 `WriteTimeOut`：这两个字段的类型都是`time.Duration`，这是Go标准库中定义的类型，用于表示时间长度。这里，它们分别用于配置读取和写入操作的超时时间。

2. **标签（Tag）**：
    - 每个字段后面的``` `yaml:"network"` ```部分是一个标签（Tag），这在Go中用于为结构体字段提供元数据。在这个例子中，这些标签用于指示`yaml`库如何将YAML文件中的键映射到结构体的对应字段。例如，YAML文件中的`network`键的值将被解析并存储到`Config`结构体的`Network`字段中。
    - 这些标签对于使用`yaml.Unmarshal`函数解析YAML配置文件并自动填充这些字段非常重要。

