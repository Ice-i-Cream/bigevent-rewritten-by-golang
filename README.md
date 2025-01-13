Mysql(DDL)
==============
Article
------------
```
-- auto-generated definition
create table article
(
    id          int unsigned auto_increment comment 'ID'
        primary key,
    title       varchar(30)               not null comment '文章标题',
    content     varchar(10000)            not null comment '文章内容',
    cover_img   varchar(128)              not null comment '文章封面',
    state       varchar(3) default '草稿' null comment '文章状态：只能是[已发布]或者[草稿]',
    category_id int unsigned              null comment '文章分类ID',
    create_user int unsigned              not null comment '创建人ID',
    create_time datetime                  not null comment '创建时间',
    update_time datetime                  not null comment '修改时间',
    constraint fk_article_category
        foreign key (category_id) references category (id),
    constraint fk_article_user
        foreign key (create_user) references user (id)
);
```
