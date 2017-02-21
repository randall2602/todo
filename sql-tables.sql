create database todo;

use todo;

create table users
(
    id varchar(255),
    primary key (id)
);

create table lists
(
    id int not null auto_increment,
    name varchar(255),
    owner varchar(255),
    primary key (id),
    foreign key (owner) references todo.users(id)
);

create table tasks
(
    id int not null auto_increment,
    parent_list int,
    description varchar(255),
    completed boolean,
    primary key (id),
    foreign key (parent_list) references todo.lists(id)
);