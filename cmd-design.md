---
date: 2018-11-03 20:06
status: public
title: cmd-design
---

# Agenda 命令设计
>agenda help [option]

列出命令说明(可选是否列出具体功能的说明)。
>agenda register -u username -p password -m email -t phone

用户注册，如果用户名已被使用，返回错误信息；如果注册成功，返回成功信息。
>agenda login -u username -p password

用户登录，登录失败返回失败原因;登录成功，返回成功信息，并列出可选操作。

>agenda logout

用户退出登录，返回成功信息并列出可选操作。

>agenda list

已登录用户查询已注册用户信息

>agenda delUser

已登录用户注销帐号，操作成功返回成功信息；否则，返回失败信息。若成功，删除一切与该用户的信息。

>agenda createMeeting -t title -p participator -s start -e end

创建会议

>agenda addPar -p participator

添加会议参与者

>agenda removePar -p participator

删除会议参与者

>agenda listUsers

列出所有用户

>agenda list -s start -e end

列出用户再一个时间段内的所有会议

>agenda cancel -t title

取消自己创建的会议

>agenda quit -t title

退出自己参与的命令

>agenda clear

清空自己发起的所有会议
