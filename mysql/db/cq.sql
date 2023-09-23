create database cq;
use cq;

create table Users(
    user_id     varchar(256) primary key,
    user_name   varchar(256),
    email       varchar(256) unique not null ,
    password    varchar(256) not null
);

create table Channels(
    channel_id varchar(256) primary key,
    owner varchar(256) not null,
    created_at  timestamp default current_timestamp,
    foreign key (owner) references Users(user_id)
);

create table Questions(
    channel_id  varchar(256) not null ,
    id          varchar(256) primary key,
    content     TEXT,
    created_at  timestamp default current_timestamp,
    foreign key (channel_id) references Channels(channel_id)
);

create table RefreshTokens(
    token       varchar(512) primary key,
    user_id     varchar(256) not null,
    expiry      timestamp    not null,
    foreign key (user_id) references Users(user_id)
);
