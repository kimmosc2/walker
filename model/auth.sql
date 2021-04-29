# authorization list
USE walker;

CREATE TABLE `walker_auths`
(
    `id`             int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`           varchar(20)         DEFAULT '' COMMENT '姓名',
    `teacher`        varchar(20)         DEFAULT '' COMMENT '辅导员',
    `tel`            varchar(11)         DEFAULT '' COMMENT '电话',
    `contact`        varchar(20)         DEFAULT '' COMMENT '紧急联系人',
    `img_url`        varchar(300)        DEFAULT '' COMMENT '图片URL',
    `contact_tel`    varchar(11)         DEFAULT '' COMMENT '紧急联系人电话',
    `direction`      varchar(100)        DEFAULT '' COMMENT '去向',
    `reason`         varchar(100)        DEFAULT '' COMMENT '请假事由',
    `school_number`  varchar(15) UNIQUE  DEFAULT '0' COMMENT '学号',
    `last_cookie`    varchar(300)        DEFAULT '0' COMMENT '最后一次的cookie',
    `start_time`     int(3) unsigned     DEFAULT '0' COMMENT '请假起始时间',
    `end_time`       int(3) unsigned     DEFAULT '0' COMMENT '请假结束时间',
    `pay_start_time` int(13) unsigned    DEFAULT '0' COMMENT '支付起始时间',
    `pay_end_time`   int(13) unsigned    DEFAULT '0' COMMENT '支付到期时间',
    `state`          tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
    `created_at`     int(13) unsigned    DEFAULT '0' COMMENT '创建时间',
    `modified_at`    int(15) unsigned    DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='权限列表';