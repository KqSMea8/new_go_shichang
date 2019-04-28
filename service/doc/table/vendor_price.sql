# drop table if exists vendor_price;
create table vendor_price
(
  id                   bigint not null auto_increment,
  vendor               varchar(100) not null comment '供应商',
  goods_name           varchar(100) not null  comment '货物名称',
  unit_price           float(10,2) not null comment '价格',
  update_date          datetime not null comment 'yyyymmdd hh:mm:ss',
  primary key (id)
);

create unique index vendor_price_idx1 on vendor_price
(
 vendor,goods_name
);