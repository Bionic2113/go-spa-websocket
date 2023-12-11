create database sushkof;

use sushkof;

DROP TABLE IF EXISTS users;

drop table if exists tasks; 


create table users(
    email varchar(255) unique not null,
    age int,
    active bool default true,
    admin bool not null,
    name text,
    password varchar(255),
    primary key (email)
);


create table tasks(
  task_id int not null auto_increment primary key,
  title varchar(255),
  description text,
  status varchar(40) not null,
  time_start BIGINT UNSIGNED,
  time_finish BIGINT UNSIGNED,
  user_email varchar(255) not null,
  FOREIGN KEY(user_email) REFERENCES users(email)
);


insert into users(email, age, admin, name, password) values 
('kekok@mail.ru', 30, false, 'Nikita', '$2a$10$o8lWXGHd5Xw2JKV8IDIV8.OXTDjqk.qpAn1R2DLPWvtNnKh4tMXRy'), --'qwerty'
('admin@mail.ru', 11, true, 'Misha', '$2a$10$xoZpp.Iis0S1aRtj8HkIBu0Cz6WNsqUbVzJWR/Sep0lJqxvOEhNwu'), --'00werty11'
('aaaa@aaa.ro', 22, false, 'Masha', '$2a$10$oFKT4rVIHKUAmf1zHuyiJeEwx1jMfujfbbVGsj/ksDAGkGa2S1tgO'), --'zima'
('kdawlkwd@aa.ri', 22, false, 'OOOad', '$2a$10$GRnf4aht7Y9orALfc1TZ0.Gxy3nB6PJgG/jbRXaveKKtRWz8VuPBC'), --'daow'
('lafwia@dlaad.fak', 22, false, 'adwkda', '$2a$10$4rtbLjLo6XTHMn.aLTYU.OKsI1dHB4BUU5E0Q/W4uNxES9gMyIBMe'), -- 'dadwaa'

insert into tasks(title, description, status, time_start, time_finish, user_email) values
('Gnom','dadadad','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kekok@mail.ru'),
('Azbika','fLNA:KWmf','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kekok@mail.ru'),
('KLDA','giewajs;flkm,aad','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kekok@mail.ru'),
('AADf','flakm','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kekok@mail.ru'),
('Ford','faw','overdue', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kekok@mail.ru'),
('Azbika','gae','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'admin@mail.ru'),
('Plot','gae','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'admin@mail.ru'),
('Kasj','gakla;','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'admin@mail.ru'),
('Cash','alwdkm.','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'admin@mail.ru'),
('Paper','lfkanwd.','completed', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'admin@mail.ru'),
('Piter','al;dwdaw;l','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'aaaa@aaa.ro'),
('Azbika','falwndk','overdue', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'aaaa@aaa.ro'),
('Parker','g;al.f','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'aaaa@aaa.ro'),
('Spider','adwlkdma','completed', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'aaaa@aaa.ro'),
('Azbika','lnasf.','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'aaaa@aaa.ro'),
('Azbika','flnakwd','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kdawlkwd@aa.ri'),
('Man','fawkmld','overdue', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kdawlkwd@aa.ri'),
('Azino99','fkwalm.','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kdawlkwd@aa.ri'),
('Marvel','f;kaf.w','pending', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kdawlkwd@aa.ri'),
('DC','afwlnkd','completed', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'kdawlkwd@aa.ri'),
('Azbika','awmd.aa','completed', UNIX_TIMESTAMP(), UNIX_TIMESTAMP() + 432000, 'lafwia@dlaad.fak');
