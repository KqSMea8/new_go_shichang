# drop table if exists groups;
create table groups
(
  id                   bigint not null auto_increment,
  group_name           varchar(100) not null comment '',
  nodes                varchar(300) comment '',
  primary key (id)
);

create unique index groups_idx1 on groups
(
 group_name
)