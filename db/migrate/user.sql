CREATE TABLE users (
                         id serial PRIMARY KEY unique,
                         name VARCHAR(255),
                         age integer,
                         profession VARCHAR(255),
                         friendly BOOLEAN
  );