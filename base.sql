create table if not exists iron_categories
(
	id int auto_increment
		primary key,
	parent_id int default 0 not null comment '父类id',
	name varchar(50) default '' not null comment '类别名称',
	description text not null comment '类别描述',
	status tinyint default 0 not null comment '类别状态',
	user_id int default 0 not null comment '创建人',
	created_at timestamp not null comment '创建时间',
	updated_at timestamp not null comment '修改时间',
	constraint icron_categories_name_uindex
		unique (name)
)
comment '训练类别';

create table if not exists iron_movements
(
	id int auto_increment
		primary key,
	cat_id int default 0 not null comment '类别 eg：胸,肩,背,腿,手臂,核心，有氧等',
	name varchar(50) default '' not null comment '动作名称 eg: 上斜哑铃卧推 ',
	description text not null comment '动作简介',
	status tinyint default 0 not null comment '状态',
	user_id int default 0 not null comment '创建人',
	created_at timestamp not null comment '创建时间',
	updated_at timestamp not null
)
comment '训练动作';

create table if not exists iron_plan_details
(
	id int auto_increment
		primary key,
	plan_id int default 0 not null comment '属于哪个计划',
	movement_id int default 0 not null comment '使用哪个动作',
	weight int default 0 not null comment '重量',
	count int default 0 not null comment '次数',
	break int default 0 not null comment '间歇',
	status tinyint default 0 not null comment '状态',
	user_id int default 0 not null comment '用户',
	created_at timestamp not null,
	updated_at timestamp not null
)
comment '计划详情表';

create table if not exists iron_plans
(
	id int auto_increment
		primary key,
	plan_name varchar(100) default '' not null comment '计划名称',
	status tinyint default 0 not null comment '状态',
	user_id int default 0 not null comment '用户id',
	created_at timestamp not null comment '创建时间',
	updated_at timestamp not null comment '修改时间'
)
comment '训练计划表';

create table if not exists iron_training_logs
(
	id int auto_increment
		primary key,
	plan_detail_id int default 0 not null comment '计划详情id',
	done tinyint default 0 not null comment '是否完成    done (0 为完成，2, 部分完成  3, 完成 4,失败)',
	status tinyint default 0 not null comment '状态',
	user_id int default 0 not null comment '用户',
	created_at timestamp not null comment '创建时间',
	updated_at timestamp not null comment '修改时间'
)
comment '训练内容';

create table if not exists iron_trainings
(
	id int auto_increment
		primary key,
	training_date date not null comment '训练日期',
	plan_id int default 0 not null comment '使用计划',
	start_time datetime not null comment '开始时间',
	end_time datetime not null comment '结束时间',
	description text not null comment '训练小记',
	status tinyint default 0 not null,
	user_id int default 0 not null comment '用户id',
	created_at timestamp not null,
	updated_at timestamp not null
)
comment '训练计划';

create table if not exists iron_users
(
	id int auto_increment
		primary key,
	user_name varchar(100) default '' not null comment '用户名',
	password varchar(100) not null,
	nick_name varchar(100) default '' not null,
	mobile varchar(100) default '' not null comment '手机号',
	wx_id varchar(100) default '' not null comment '微信号',
	status tinyint default 0 null comment '状态',
	created_at timestamp default CURRENT_TIMESTAMP not null,
	updated_at timestamp default CURRENT_TIMESTAMP null,
	constraint users_name_uindex
		unique (user_name)
)
comment '用户表';

