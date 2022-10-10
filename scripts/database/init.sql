#创建数据库

create database if not exists lydb default charset utf8mb4;


#创建数据表
create table if not exists user(
    id bigint primary key auto_increment,
    employee_id varchar(32) not null default '',
    username varchar(64) not null default '',
    nickname varchar(64) not null default '',
    avatar varchar(1024) not null  default '',
    password varchar(1024) not null  default '',
    gender int not null default 0 comment '0女 1男',
    phone_num varchar(32) not null default '',
    address varchar(128) not null default '',
    email varchar(64) not null default '',
    department varchar(64) not null default '',
    post varchar(64) not null default '',
    role int not null default 2 comment '0系统管理员 1部门管理员 2普通员工',
    status int not null default 1 comment '0已禁用 1使用中',
    created_at datetime not null,
    updated_at datetime not null


)engine=innodb default charset utf8mb4;

#插入数据
insert into `user`(employee_id,username,nickname,avatar,password,gender,phone_num,address,email,department,post,role,status,created_at,updated_at) VALUES("ly0003","kangkang","康%康","",md5("kangkang"),1,"18384235350","成都市","kangkang@kang.com","测试运维部","运维开发",0,1,now(),now())