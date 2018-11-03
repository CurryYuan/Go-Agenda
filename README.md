# Go-Agenda

## 命令设计

见cmd-design.md

## 接口约束

cmd 用来创建命令行接口接口，entity 用来实现内部逻辑，cmd 调用 entity 函数

### entity接口

用户注册

func Register(username,password,mail,phone) error

用户登录

func Login(user, password string) error

用户登出

func Logout() error

列出所有用户

func ListUsers() error

删除用户

func DelUser() error

创建会议

func CreateMeeting(title string, participators []string, start string, end string) error

添加参会者

func AddPar(title string, participators []string) error 

删除参会者

func RemovePar(title string, participators []string) error 

列出用户再某个时间段的会议

func ListMeetings(start, end string) error

取消自己发起的会议

func CancelMeeting(title string) error 

退出自己参与的会议

func QuitMeeting(title string) error 

清空自己发起的会议

func ClearMeeting() error 

## 测试

