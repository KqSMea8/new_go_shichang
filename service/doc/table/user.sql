# drop table if exists user;
create table user
(
  id                   bigint not null auto_increment,
  user_name            varchar(100) not null comment '',
  password             varchar(100) not null comment '',
  is_valid             varchar(100) not null comment 'y:启用，n：未启用',
  primary key (id)
);

create unique index user_idx1 on user
(
 user_name
)