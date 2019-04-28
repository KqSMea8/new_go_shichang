# drop table if exists user_group;
create table user_group
(
  id                   bigint not null auto_increment,
  group_name           varchar(100) not null comment '',
  user_name            varchar(100) not null comment '',
  primary key (id)
);

create unique index user_group_idx1 on user_group
(
   group_name,user_name
)