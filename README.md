# SimpleTikTok

### 创建go-zero
```shell
goctl api new BaseInterface
goctl api new great   # 多了一个整个great文件夹
go mod tidy
```

### 修改API
./GoZeroUse.sh create
### 编译
./GoZeroUse.sh build
### 启动go-zero
./GoZeroUse.sh run


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

# 参考资料

https://go-zero.dev/cn/docs/quick-start/monolithic-service
https://go-zero.dev/cn/docs/goctl/goctl/
https://go-zero.dev/cn/docs/goctl/goctl/