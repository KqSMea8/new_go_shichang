# drop table if exists goods;
create table goods
(
  id                   bigint not null auto_increment,
  goods_name           varchar(100) not null comment '货物名称',
  goods_alias          varchar(100)  comment '货物别名',
  group_name           varchar(100) not null comment '大类:石仔、石粉',
  order_index          int not null comment '排序',
  primary key (id)
);

create unique index goods_idx1 on goods
(
 goods_name
);