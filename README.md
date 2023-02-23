# SimpleTikTok
go-zero微服务

sudo apt-get install ffmpeg # 安装
sudo apt-get purge ffmpeg # 卸载

ffmpeg -i /yzx/src/SimpleTikTok/source/video/video_test2.mp4 -filter_complex [0]select=gte(n\,1)[s0] -map [s0] -f image2 -vcodec mjpeg -vframes 1 pipe:
## 0. 获取帮助
```shell
useage: ./GoZeroUse.sh 
```

## 1. goctl生成
什么都不填，两个都会生成
```shell
useage: ./GoZeroUse.sh create api
        ./GoZeroUse.sh create proto
        ./GoZeroUse.sh create 
```
### 2. 编译
现在编译会输出参数
```shell
useage: ./GoZeroUse.sh build api
        ./GoZeroUse.sh build proto
        ./GoZeroUse.sh build proto -a # 重新编译
        ./GoZeroUse.sh build  # 全部编译
```
### 3. 启动
发现运行失败，尝试使用ps查看进程是否已经启动了
如果有启动的进程，使用```./GoZeroUse.sh.sh kill all```清除
```shell
useage: ./GoZeroUse.sh run api
        ./GoZeroUse.sh run proto
```

### 4. 停止进程
```shell
useage: ./GoZeroUse.sh kill # 显示帮助信息，一定要填写参数
        ./GoZeroUse.sh kill api
        ./GoZeroUse.sh kill etcd
        ./GoZeroUse.sh kill all
```

### 5. 清理日志文件
```shell
useage: ./GoZeroUse.sh clear 
```


# nginx脚本说明

## 安装nginx
目前nginx没有装载于docker中，具体原因开会的时候再细说
没有安装nginx环境的可以用
```Shell
./NginxUse.sh install
```
执行Nginx安装

## 配置nginx
nginx配置文件即文件目录下的default.conf（nginx默认配置）
快速配置该文件
```Shell
./NginxUse.sh use
```

## 对漏桶限流的一些补充说明
目前在网上查找到的资料只有nginx配置漏桶限流，具体参数含义我在配置文件中注释了

以目前参数为例，假设1s内user接受30个req
其中15个req被处理，5个req背放入桶中。剩下的10个被drop

这里具体每个参数的值设为多少最好待调试


# redis相关

```shell
sudo apt-get update 
sudo apt-get install redis # 安装
redis-server -v # 查看redis版本
redis-cli # 登录redis
```
```shell
set 123 test
get 123 
del 123
ping

keys * # 查一下有什么键

SMEMBERS FollowAndFollowerList:follower_id:323    # 我们的数据类型是set，查询一下关注列表
SMEMBERS FollowAndFollowerList:user_id:1    # 查询一下粉丝列表

# 获取set值
smembers follower_id:300
# 删除当前数据库中的所有Key
flushdb
# 删除所有数据库中的key
flushall
```
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
[Goctl 简介]https://go-zero.dev/cn/docs/goctl/goctl/
[go-redis]https://redis.uptrace.dev/zh/guide/
[使用redis实现关注和粉丝列表 ｜ 青训营笔记]https://juejin.cn/post/7103869225260810277
[go-redis文档]https://juejin.cn/post/7027347979065360392#heading-20
[redis的五种数据结构和应用场景【如微博微信点赞/共同关注/加购物车】]https://juejin.cn/post/6895185457110319118#heading-21