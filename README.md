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

```
./agenda -h
A program for meeting manage

Usage:
  agenda [command]

Available Commands:
  addPar        Add your own meetings' participators.
  cancel        Cancel your own meeting by specifying title name.
  createMeeting create meeting
  delUser       delete user
  help          Help about any command
  list          List all of your own meetings during a time interval.
  listUsers     list all users
  login         user login
  logout        user logout
  quit          Quit others meeting by specifying title name
  register      Register user
  removePar     Remove your own meetings' participators.

Flags:
      --config string   config file (default is $HOME/.agenda.yaml)
  -h, --help            help for agenda
  -t, --toggle          Help message for toggle

Use "agenda [command] --help" for more information about a command.

```


```
./agenda register -u a -p b -m c -t d
register success
```

```
./agenda login -u a -p b
login success
```

```
./agenda login -u a -p b
login success
```
```
./agenda listUsers
[list all users]
---------------------------------------------------
[username]	yhz
[email]		1
[phone]		123
---------------------------------------------------
[username]	yh
[email]		1
[phone]		123
---------------------------------------------------
[username]	yzz
[email]		111
[phone]		123
---------------------------------------------------
[username]	a
[email]		c
[phone]		d
---------------------------------------------------
```

```
./agenda createMeeting -t t -p yhz a -s 2018-01-01/00:00 -e 2018-01-02/00:00
create meeting success
```

```
./agenda list -s 2018-01-01/00:00 -e 2018-02-01/00:00
[list meetings in 2018-01-01/00:00 —— 2018-02-01/00:00]
---------------------------------------------------
[sponsor]	a
[title]		t
[start]   	2018-1-1/0:0
[end]		2018-1-2/0:0
[participator]
	
  yhz---------------------------------------------------

```

```
./agenda addPar -t t  -p yzz yz
Add meeting participator successfully!
```

```
./agenda removePar -t t -p yhz yz
Remove meeting participator successfully!
```

```
./agenda list -s 2018-01-01/00:00 -e 2018-02-01/00:00
[list meetings in 2018-01-01/00:00 —— 2018-02-01/00:00]
---------------------------------------------------
[sponsor]	a
[title]		t
[start]   	2018-1-1/0:0
[end]		2018-1-2/0:0
[participator]
	
  yzz---------------------------------------------------
Listing meeting operation completed successfully!

```
