create database blogger character set utf8;

use blogger;

create table article (
id bigint(20) primary key auto_increment not null,/*文章id*/
category_id bigint(20) not null,/*分类id*/
content longtext not null,/*文章内容*/
title varchar(1024) not null,/*文章标题*/
view_count int(255) not null,/*阅读次数*/
comment_count int(255) not null,/*评论次数*/
username varchar(128) not null,/*作者*/
status int(10) not null default 1,/*状态，正常为1*/
summary varchar(256) not null,/*文章摘要*/
create_time timestamp default current_timestamp,/*发布时间*/
update_time timestamp/*更新时间*/
);

create table category(
id bigint(20) primary key auto_increment not null,
category_name varchar(255) not null,
category_no bigint(20) not null,
create_time timestamp default current_timestamp,/*发布时间*/
update_time timestamp/*更新时间*/
);

create table comment(
id bigint(20) primary key auto_increment not null,/*评论id*/
content text not null,/*评论内容*/
username varchar(64) not null,/*评论作者*/
create_time timestamp default current_timestamp not null,/*发布时间*/
status int(255) default 1 not null,/*评论状态*/
article_id bigint(20) not null/*关联article表id*/
);

create table `leave`(
id bigint(20) primary key auto_increment not null,
username varchar(255) not null,
email varchar(255) not null,
content text not null,
create_time timestamp default current_timestamp,
update_time timestamp
);

insert into category (category_name,category_no,update_time)values("css/html",1,now());
insert into category (category_name,category_no,update_time)values("后端开发",2,now());
insert into category (category_name,category_no,update_time)values("Java开发",3,now());
insert into category (category_name,category_no,update_time)values("C++开发",4,now());
insert into category (category_name,category_no,update_time)values("架构剖析",5,now());
insert into category (category_name,category_no,update_time)values("Golang开发",6,now());