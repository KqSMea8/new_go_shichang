# drop table if exists user;
# roles control
# https://studygolang.com/topics/6999
# https://github.com/casbin/casbin
# demo
#https://github.com/gin-contrib/authz
create table user
(
  id                   bigint not null auto_increment,
  user_name            varchar(100) not null comment '',
  password             varchar(100) not null comment '',
  is_valid             varchar(100) not null comment 'Y:启用，N：未启用',
  is_supper_user       varchar(20) not null comment 'Y/N',
  primary key (id)
);

create unique index user_idx1 on user
(
 user_name
)