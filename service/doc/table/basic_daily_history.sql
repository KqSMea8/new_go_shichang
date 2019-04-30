# drop table if exists basic_daily_history;
create table basic_daily_history
(
id                   bigint not null auto_increment,
driver_name          varchar(100) comment '驾驶员',
number_flag          varchar(100) not null  comment '编号',
car_number           varchar(100) not null comment '车号',
shop_number	        varchar(60) not null comment '货号',
gross_weight	        varchar(60) comment '毛重',
tare_weight          varchar(60) comment '皮重',
net_weight           varchar(60) not null comment '净重',
vendor_net_weight    varchar(60) comment '供应商净重',
vendor               varchar(60) not null comment '供应商,例如成协',
vendor_sub_name      varchar(60) comment '供应商，例如成协(坝仔村)',
group_name           varchar(60) not null comment '石粉、石仔',
batch_id             varchar(60) not null comment '导入批次号',
row_index            int comment '行号',
import_date          datetime not null comment '导入时间',
day_str              varchar(20) not null comment '称磅日期yyyymmdd',
time_str             varchar(20) not null comment '称磅时间 hh:mm:ss',
weight_date          datetime not null comment '称磅时间yyyymmdd hh:mm:ss',
price_batch_id       varchar(100) not null comment '采用的价格批次号',
primary key (id)
);

create index basic_daily_history_idx1 on basic_daily_history
(
number_flag
);