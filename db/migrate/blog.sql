CREATE TABLE blog (
                       id serial PRIMARY KEY unique,
                       title VARCHAR(255) not null ,
                       number_of_comments integer default 0,
                        content VARCHAR(255) not null,
                        author_id integer not null,
                        FOREIGN KEY (author_id) REFERENCES users (id)

);