CREATE TABLE users (
  id serial PRIMARY KEY,
  name VARCHAR (50) NOT NULL,
  age INT NOT NULL,
  active BOOLEAN NOT NULL
);

INSERT INTO users VALUES
  (1, 'User One', 10, true),
  (2, 'User Two', 20, false),
  (3, 'User Three', 30, true),
  (4, 'User Four', 40, false),
  (5, 'User Five', 50, true);
