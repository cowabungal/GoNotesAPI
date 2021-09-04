CREATE TABLE users
(
    id       serial       not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null
);

CREATE TABLE notes
(
    id          serial                                      not null unique,
    title       varchar(255)                                not null,
    description varchar(25555)                                not null,
    user_id     int references users (id) on delete cascade not null
);
