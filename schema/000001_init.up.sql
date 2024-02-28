CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) NOT NULL,
    username      varchar(255) NOT NULL unique,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) NOT NULL,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) NOT NULL,
    description varchar(255),
    done        boolean      NOT NULL default false
);

CREATE TABLE list_items
(
    id      serial                                           not null unique,
    item_id int references todo_items (id) on DELETE cascade not null,
    list_id int references todo_lists (id) on DELETE cascade not null
);