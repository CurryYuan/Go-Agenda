# Go-Agenda

## 接口约束

cmd 用来创建命令行接口接口，entity 用来实现内部逻辑，cmd 调用 entity 函数

### entity接口

Register(username,password,mail,phone) error
