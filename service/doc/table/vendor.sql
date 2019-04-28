# drop table if exists vendor;
create table vendor
(
  id                   bigint not null auto_increment,
  vendor_name          varchar(100) not null comment '供应商名称,例如成协',
  tel                  varchar(50) comment '',
  contact_name         varchar(100) comment '联系人',
  cost_advance         bigint not null comment '预付金额元',
  cost_alert           bigint not null comment '预警线',
  need_cost_advance    varchar(20) not null comment 'Y:需要预付；N:不需要预付',
  is_valid             varchar(20) not null comment 'Y:启用，N：未启用',
  primary key (id)
);

create unique index vendor_idx1 on vendor
(
 vendor_name
)