-- 建立yourdb库
create database yourdb;

-- 创建user表
USE yourdb;
CREATE TABLE user(
    username char(50) NULL,
    passwd char(50) NULL
)ENGINE=InnoDB;

-- 添加数据
INSERT INTO user(username, passwd) VALUES('name', 'passwd');

-- 查看数据库
show tables;

-- 查看数据
select * from user;



#指定账号在某个host才可以登录 
#CREATE USER 'username'@'host' IDENTIFIED BY 'password';
create user 'yzx'@'%' identified by '123456';
flush privileges;

