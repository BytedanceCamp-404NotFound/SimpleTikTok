# nginx脚本说明

## 安装nginx
目前nginx没有装载于docker中，具体原因开会的时候再细说
没有安装nginx环境的可以用
```Shell
./Nginx.sh setup
```
执行Nginx安装

## 配置nginx
nginx配置文件即文件目录下的default.conf（nginx默认配置）
快速配置该文件
```Shell
./Nginx.sh use
```

## 对漏桶限流的一些补充说明
目前在网上查找到的资料只有nginx配置漏桶限流，具体参数含义我在配置文件中注释了

以目前参数为例，假设1s内user接受30个req
其中15个req被处理，5个req背放入桶中。剩下的10个被drop

这里具体每个参数的值设为多少最好待调试