# drop table if exists report_file;
create table report_file
(
  id                   bigint not null auto_increment,
  report_type          varchar(100) not null comment 'WX_DAILY/DAILY/WEEKLY/MONTHLY',
  file_name            varchar(200) comment '',
  batch_id             varchar(100) comment '批次号',
  day_str              varchar(20) not null comment '报表日期yyyymmdd',
  report_date          datetime not null comment '报表日期yyyymmdd hh:mm:ss',
  primary key (id)
);

create unique index report_file_idx1 on report_file
(
 report_date
);

create unique index report_file_idx2 on report_file
(
 report_type
)