create table users(
    user_id int auto_increment,
    user_information_id int,
    user_name text,
    email text,
    primary key (user_id)
) ;

create table user_information(
    id int auto_increment,
    last_name text,
    first_name text,
    sex text,
    age int,
    post_code text,
    address text,
    remarks text,
    primary key (id)
) ;

insert into users(user_information_id user_name email)values(1,'山田太郎', 'email.com1');
insert into users(user_information_id,user_name,email)values(2,'田中太郎', 'email.com2');
insert into users(user_information_id,user_name,email)values(3,'村田花子', 'email.com3');
insert into users(user_information_id,user_name,email)values(4,'高橋姫子', 'email.com4');

insert into user_information(last_name,first_name,sex,age,post_code,address,remarks)values('山田', '太郎','男',25,'111-1111','東京都港区1-1-1','');
insert into user_information(last_name,first_name,sex,age,post_code,address,remarks)values('田中', '太郎','男',23,'222-2222','東京都港区2-2-2','');
insert into user_information(last_name,first_name,sex,age,post_code,address,remarks)values('村田', '花子','女',30,'333-3333','東京都港区3-3-3','');
insert into user_information(last_name,first_name,sex,age,post_code,address,remarks)values('高橋', '姫子','女',20,'444-4444','東京都港区4-4-4','');