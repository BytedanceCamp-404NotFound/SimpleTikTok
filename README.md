# SimpleTikTok
测试传代码


变量命名使用大驼峰
### 编译
```shell
cd greet
go build greet.go 
```

### 调试
在任何位置启动F5都可以使用，但是需要先编译

BasicInterface
### 创建go-zero
```shell
goctl api new BaseInterface
goctl api new great   # 多了一个整个great文件夹
go mod tidy
```

### 修改API
goctl api go -api api/BaseInterface.api -dir .. -style GoZero

### 启动go-zero
```shell
cd BasicInterface
go run BasicInterface-Api.go -f etc/BasicInterface-Api.yaml
```
### 测试go-zero
```shell
curl -i -X GET http://localhost:8888/from/you
```
浏览器中打开
http://[自己IP地址]:8080/from/you




# 临时放在这里的环境配置

### 安装mysql
sudo apt install mysql-server -y
检查安装的是否正确
sudo systemctl status mysql.service

# 未来计划以及已完成

+ [ ] vscode可以F5一键编译和调试
+ [X] 增加Action的CI检测


# commit类型

用于说明 commit 的类别，只允许使用下面7个标识。
feat：新功能（feature）</br>
fix/to：修补bug </br>
  - fix：产生 diff 并自动修复此问题。适合于一次提交直接修复问题 </br>
  - to：只产生 diff不 自动修复此问题。适合于多次提交。最终修复问题提交时使用 fix </br>
docs：仅仅修改了文档（documentation） </br>
style： 仅仅修改了空格、格式缩进、逗号等等，不改变代码逻辑 </br>
refactor：代码重构，没有加新功能或者修复 bug（即不是新增功能，也不是修改bug的代码变动） </br>
test：增加测试 </br>
chore：改变构建流程、或者增加依赖库、工具等 </br>
revert：回滚到上一个版本 </br>
merge：代码合并 </br>
sync：同步主线或分支的Bug </br>


# vscode 写go代码，没法转跳函数

ctrl + shift + p 
Go:Install/Update Tools
全选
OK

setting
 "go.toolsEnvVars": {
        "GOFLAGS": "-mod=vendor"
}


ctrl + shift + p 
reload

Go: Use Language Serve


/usr/local/bin/  # 放可执行文件的位置

go-outline之类的，放在/usr/local/bin/下面，就可以转跳了
# 参考资料

https://go-zero.dev/cn/docs/quick-start/monolithic-service
https://go-zero.dev/cn/docs/goctl/goctl/
https://go-zero.dev/cn/docs/goctl/goctl/