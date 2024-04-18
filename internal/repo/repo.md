按照我现在的理解，我觉得repo属于数据访问层：

"repo" 指的是一个封装了数据访问逻辑的类或结构体。在这种用法下，"repo" 作为一个抽象层存在，用来隔离业务逻辑和数据存储/检索操作。这种设计模式常见于应用软件架构中，特别是在遵循领域驱动设计（Domain-Driven Design, DDD）或其他需要清晰分层的架构风格中。这样的 "repository" 提供了一个或多个模型的CRUD（创建、读取、更新、删除）操作接口。

语法上没有多说的，就是目前只有一个db成员，通过internal/repo/db封装的DB来初始化数据库（因为现在只有数据库被定义了）