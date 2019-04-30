# drop table if exists report_notify;
create table report_notify
(
  id                   bigint not null auto_increment,
  date_str             varchar(100) not null comment '',
  weight_date          datetime not null comment '',
  handle_flag          varchar(10) not null comment 'Y/N',
  primary key (id)
);

create index report_notify_idx1 on report_notify
(
   handle_flag
);