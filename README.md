# 【逼上梁山系列】07-《Go语言从入门到实战》操作Mysql & Slice讲解 & 项目分层 & 小程序调用本地接口

## 1. 第一次更新 20240419 23:20

## 2. 第二次更新 20240422 18:55：
今天记录一个BUG以及DEBUG的过程：
1.BUG的原因是：
common/database/db.go下的一个方法的参数写错了：
func (db *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
    return db.ExecContext(ctx, query, args)
}
最开始args这里没有加...
A : 就是这里的原因！！！！下次细心点吧球球了！浪费了一小时 干！
2.DEBUG的过程：
首先是根据报错：
PS C:\NewWorker\projectforbingmeishi\rose> go run .\cmd\main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
- using env:   export GIN_MODE=release
- using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /v1/user/create           --> rose/internal/handler.userCreate (1 handlers)
[GIN-debug] GET    /v1/user/detail           --> rose/internal/handler.userDetail (1 handlers)
[GIN-debug] GET    /v1/user/list             --> rose/internal/handler.userList (1 handlers)
[GIN-debug] [WARNING] Headers were already written. Wanted to override status code 500 with 200

去找/v1/user/create路由下对应的方法rose/internal/handler.userCreate，然后发现这里关联了数据库操作相关的方法，于是加了一个JSON返回的内容：

ret, err := svc.UserCreate(c.Request.Context(), r.UserName)
if err != nil {
    // panic(err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "database insert error!"})
}

发现果然返回了"error": "database insert error!"

那就说明 ret, err := svc.UserCreate(c.Request.Context(), r.UserName) 这里UserCreate可能出了问题

然后加入panic(err)就能发现是有一个参数设置错误，然后一路找下去方法实现，就找到错误了


2.curl我在powershell中用不了，不知道为啥。然后就换了postman用，postman有界面，更适合新手吧我觉得
## 3.第三次更新 20240423 18:44
lesson8学习完成，代码更新在lessoncode/lesson8路径下

lesson7的新增代码（数据库操作相关）的内容还没有做解析
