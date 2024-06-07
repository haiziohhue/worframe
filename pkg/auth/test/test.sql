DROP TABLE IF EXISTS sys_dept;
DROP TABLE IF EXISTS sys_user;
DROP TABLE IF EXISTS sys_role;
DROP TABLE IF EXISTS sys_menu;
DROP TABLE IF EXISTS sys_user_role;
DROP TABLE IF EXISTS sys_role_menu;
DROP TABLE IF EXISTS sys_role_dept;
DROP TABLE IF EXISTS sys_user_post;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS sys_dept
(
    dept_id   BIGSERIAL PRIMARY KEY,
    parent_id BIGINT      DEFAULT 0,
    ancestors VARCHAR(50) DEFAULT '',
    dept_name VARCHAR(30) DEFAULT '',
    order_num INT         DEFAULT 0,
    leader    VARCHAR(20),
    phone     VARCHAR(11),
    email     VARCHAR(50),
    status    CHAR(1)     DEFAULT '0',
    del_flag  CHAR(1)     DEFAULT '0'
);

CREATE TABLE IF NOT EXISTS sys_user
(
    user_id   BIGSERIAL PRIMARY KEY,
    uuid      uuid         DEFAULT uuid_generate_v4(),
    dept_id   BIGINT,
    user_type VARCHAR(2)   DEFAULT '00',
    user_name VARCHAR(30)  DEFAULT '',
    nick_name VARCHAR(30) NOT NULL,
    email     VARCHAR(50)  DEFAULT '',
    phone     VARCHAR(30)  DEFAULT '',

    avatar    VARCHAR(100) DEFAULT '',
    password  VARCHAR(50)  DEFAULT '',

    salt      VARCHAR(20)  DEFAULT '',
    status    CHAR(1)      DEFAULT '0',
    del_flag  CHAR(1)      DEFAULT '0',
    remark    VARCHAR(500)
);

CREATE TABLE IF NOT EXISTS sys_role
(
    role_id   BIGSERIAL PRIMARY KEY,
    role_name VARCHAR(30) NOT NULL,
    role_key  VARCHAR(100) DEFAULT '',
    role_sort INT         NOT NULL,
    type      CHAR(1)      DEFAULT '1',
    status    CHAR(1)      DEFAULT '1',
    del_flag  CHAR(1)      DEFAULT '0',
    remark    VARCHAR(500)
);

CREATE TABLE IF NOT EXISTS sys_menu
(
    menu_id    BIGSERIAL PRIMARY KEY,
    menu_name  VARCHAR(50) NOT NULL,
    parent_id  BIGINT       DEFAULT 0,
    order_num  INT          DEFAULT 0,
    url        VARCHAR(200) DEFAULT '#',
    target     VARCHAR(20)  DEFAULT '',
    menu_type  CHAR(1)      DEFAULT '',
    visible    CHAR(1)      DEFAULT '0',
    is_refresh CHAR(1)      DEFAULT '1',
    perms      VARCHAR(100),
    icon       VARCHAR(100) DEFAULT '#',
    remark     VARCHAR(500)
);
CREATE TABLE IF NOT EXISTS sys_user_role
(
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id)
);
CREATE TABLE IF NOT EXISTS sys_role_menu
(
    role_id BIGINT NOT NULL,
    menu_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, menu_id)
);
CREATE TABLE IF NOT EXISTS sys_role_dept
(
    role_id BIGINT NOT NULL,
    dept_id BIGINT NOT NULL,
    PRIMARY KEY (role_id, dept_id)
);
CREATE TABLE IF NOT EXISTS sys_user_post
(
    user_id BIGINT NOT NULL,
    post_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, post_id)
);
insert into sys_dept (dept_id, parent_id, ancestors, dept_name, order_num, leader, phone, email)
VALUES (100, 0, '0', '若依科技', 0, '若依', '15888888888', 'ry@qq.com'),
       (101, 100, '0,100', '深圳总公司', 1, '若依', '15888888888', 'ry@qq.com'),
       (102, 100, '0,100', '长沙分公司', 2, '若依', '15888888888', 'ry@qq.com'),
       (103, 101, '0,100,101', '研发部门', 1, '若依', '15888888888', 'ry@qq.com'),
       (104, 101, '0,100,101', '市场部门', 2, '若依', '15888888888', 'ry@qq.com'),
       (105, 101, '0,100,101', '测试部门', 3, '若依', '15888888888', 'ry@qq.com'),
       (106, 101, '0,100,101', '财务部门', 4, '若依', '15888888888', 'ry@qq.com'),
       (107, 101, '0,100,101', '运维部门', 5, '若依', '15888888888', 'ry@qq.com'),
       (108, 102, '0,100,102', '市场部门', 1, '若依', '15888888888', 'ry@qq.com'),
       (109, 102, '0,100,102', '财务部门', 2, '若依', '15888888888', 'ry@qq.com');
insert into sys_user (user_id,
                      dept_id,
                      user_name,
                      nick_name,
                      user_type,
                      email, phone, password, salt)
VALUES (1, 103, 'admin', '若依', '00', 'ry@163.com', '15888888888', '29c67a30398638269fe600f73a054934', '111111'),
       (2, 105, 'ry', '若依', '00', 'ry@qq.com', '15666666666', '8e6d98b90472783cc73c17047ddccf36', '222222');
insert into sys_role (role_id, role_name, role_sort, type)
VALUES ('1', '超级管理员', 1, 1),
       ('2', '普通角色', 2, 2);
INSERT into sys_menu(menu_id, menu_name, parent_id, order_num, url, target, menu_type, visible, is_refresh, perms, icon,
                     remark)
VALUES ('1', '系统管理', '0', '1', '#', '', 'M', '0', '1', '', 'fa fa-gear', 'admin'),
       ('2', '系统监控', '0', '2', '#', '', 'M', '0', '1', '', 'fa fa-video-camera', 'admin'),
       ('3', '系统工具', '0', '3', '#', '', 'M', '0', '1', '', 'fa fa-bars', 'admin'),
       ('4', '若依官网', '0', '4', 'https://ruoyi.vip', 'menuBlank', 'C', '0', '1', '', 'fa fa-location-arrow',
        'admin'),
       ('100', '用户管理', '1', '1', '/system/user', '', 'C', '0', '1', 'system:user:view', 'fa fa-user-o', 'admin'),
       ('101', '角色管理', '1', '2', '/system/role', '', 'C', '0', '1', 'system:role:view', 'fa fa-user-secret',
        'admin'),
       ('102', '菜单管理', '1', '3', '/system/menu', '', 'C', '0', '1', 'system:menu:view', 'fa fa-th-list', 'admin'),
       ('103', '部门管理', '1', '4', '/system/dept', '', 'C', '0', '1', 'system:dept:view', 'fa fa-outdent', 'admin'),
       ('104', '岗位管理', '1', '5', '/system/post', '', 'C', '0', '1', 'system:post:view', 'fa fa-address-card-o',
        'admin'),
       ('105', '字典管理', '1', '6', '/system/dict', '', 'C', '0', '1', 'system:dict:view', 'fa fa-bookmark-o',
        'admin'),
       ('106', '参数设置', '1', '7', '/system/config', '', 'C', '0', '1', 'system:config:view', 'fa fa-sun-o', 'admin'),
       ('107', '通知公告', '1', '8', '/system/notice', '', 'C', '0', '1', 'system:notice:view', 'fa fa-bullhorn',
        'admin'),
       ('108', '日志管理', '1', '9', '#', '', 'M', '0', '1', '', 'fa fa-pencil-square-o', 'admin'),
       ('109', '在线用户', '2', '1', '/monitor/online', '', 'C', '0', '1', 'monitor:online:view', 'fa fa-user-circle',
        'admin'),
       ('110', '定时任务', '2', '2', '/monitor/job', '', 'C', '0', '1', 'monitor:job:view', 'fa fa-tasks', 'admin'),
       ('111', '数据监控', '2', '3', '/monitor/data', '', 'C', '0', '1', 'monitor:data:view', 'fa fa-bug', 'admin'),
       ('112', '服务监控', '2', '4', '/monitor/server', '', 'C', '0', '1', 'monitor:server:view', 'fa fa-server',
        'admin'),
       ('113', '缓存监控', '2', '5', '/monitor/cache', '', 'C', '0', '1', 'monitor:cache:view', 'fa fa-cube', 'admin'),
       ('114', '表单构建', '3', '1', '/tool/build', '', 'C', '0', '1', 'tool:build:view', 'fa fa-wpforms', 'admin'),
       ('115', '代码生成', '3', '2', '/tool/gen', '', 'C', '0', '1', 'tool:gen:view', 'fa fa-code', 'admin'),
       ('116', '系统接口', '3', '3', '/tool/swagger', '', 'C', '0', '1', 'tool:swagger:view', 'fa fa-gg', 'admin'),
       ('500', '操作日志', '108', '1', '/monitor/operlog', '', 'C', '0', '1', 'monitor:operlog:view',
        'fa fa-address-book', 'admin'),
       ('501', '登录日志', '108', '2', '/monitor/logininfor', '', 'C', '0', '1', 'monitor:logininfor:view',
        'fa fa-file-image-o', 'admin');

insert into sys_user_role(user_id, role_id)
values ('1', '1'),
       ('2', '2');
-- ----------------------------
INSERT INTO sys_role_menu
values ('2', '1'),
       ('2', '2'),
       ('2', '3'),
       ('2', '4'),
       ('2', '100'),
       ('2', '101'),
       ('2', '102'),
       ('2', '103'),
       ('2', '104'),
       ('2', '105'),
       ('2', '106'),
       ('2', '107'),
       ('2', '108'),
       ('2', '109'),
       ('2', '110'),
       ('2', '111'),
       ('2', '112'),
       ('2', '113'),
       ('2', '114'),
       ('2', '115'),
       ('2', '116'),
       ('2', '500'),
       ('2', '501');
insert into sys_role_dept(role_id, dept_id)
values ('2', '100'),
       ('2', '101'),
       ('2', '105');