# drop table if exists role_group;
create table role_group
(
  id                   bigint not null auto_increment,
  group_name           varchar(100) not null comment '',
  role_name            varchar(100) not null comment '',
  primary key (id)
);

create unique index role_group_idx1 on role_group
(
   group_name,role_name
)