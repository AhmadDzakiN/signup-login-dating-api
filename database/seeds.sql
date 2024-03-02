CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username varchar(255) not NULL,
    password varchar(255) not NULL,
    fullname varchar(125) not NULL,
    location varchar(100) not null,
    gender varchar(50) not null,
    age int not null
);
