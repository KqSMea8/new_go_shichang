# drop table if exists role;
create table role
(
  id                   bigint not null auto_increment,
  role_name           varchar(100) not null comment '',
  nodes                varchar(300) comment '',
  primary key (id)
);

create unique index role_idx1 on role
(
 role_name
)