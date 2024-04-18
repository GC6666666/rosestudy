**`当前目录定义了所有参数和返回值`**

定义了几个用于数据传输的结构体（DTOs，即Data Transfer Objects）。这些结构体通常用于封装客户端和服务器之间传输的数据。


#### UserDetailReq
```go
UserDetailReq struct {
	UserId int64 `form:"userId"`
}
```
这是一个请求结构体，用于用户详情请求。它包含一个字段 `UserId`，类型为 `int64`。`form:"userId"` 是一个结构体标签，用于表单解析时映射请求中的数据到结构体字段。

#### UserDetailResp
```go
UserDetailResp struct {
	*User
}
```
这个响应结构体用于返回用户详情。它嵌入了一个指向 `User` 结构体的指针。在Go中，可以通过嵌入结构体来实现类似继承的功能。这里 `User` 的所有字段都会被包含在 `UserDetailResp` 中。

#### User
```go
User struct {
	UserId     int64  `json:"userId"`
	UserName   string `json:"userName"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}
```
`User` 结构体用于表示用户的基本信息，包括用户ID、用户名、创建时间和更新时间。这里的JSON标签（例如 `json:"userId"`）指定了在序列化和反序列化JSON时，字段对应的JSON键名称。

#### UserCreateReq
```go
UserCreateReq struct {
	UserName string `json:"userName"`
}
```
这是创建用户请求的结构体，包含一个字段 `UserName`。这里使用了JSON标签以确保在处理JSON数据时使用正确的键。

#### UserCreateResp
```go
UserCreateResp struct {
	UserId int64 `json:"userId"`
}
```
用户创建响应的结构体，包含一个 `UserId` 字段，用来返回新创建的用户ID。

#### UserListReq
```go
UserListReq struct {
	UserIds string `form:"userIds"`
}
```
这个请求结构体用于批量获取用户信息，其中 `UserIds` 可能是一个用逗号分隔的ID字符串。`form:"userIds"` 标签用于从表单数据中解析此字段。

#### UserListResp
```go
UserListResp struct {
	Users []*User `json:"users"`
}
```
这是用户列表响应的结构体，包含一个 `Users` 字段，它是指向 `User` 结构体指针的切片，用于批量返回用户信息。

#### 总结
以上定义的结构体被用在API请求和响应中，以标准的格式封装数据，使得前端和服务端的数据交换更加规范和安全。结构体中使用的标签（如`json`和`form`）帮助Go语言在解析JSON或表单数据时正确映射字段。这种方式提高了代码的可读性和维护性，是处理HTTP请求数据的常见做法。