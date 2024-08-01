BEGIN;

DROP TABLE IF EXISTS todos CASCADE;
DROP TABLE IF EXISTS users CASCADE;


CREATE TABLE users
(
	user_id SERIAL PRIMARY KEY,
	username TEXT,
	password TEXT
);

CREATE TABLE todos
(
	todo_id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES users(user_id),
	title TEXT,
	content TEXT
);

INSERT INTO users (username, password)
VALUES ('username', 'password');

INSERT INTO todos (user_id, title, content)
VALUES (1, 'title', 'content');

INSERT INTO todos (user_id, title, content)
VALUES (1, 'title2', 'content2');

INSERT INTO todos (user_id, title, content)
VALUES (1, 'title3', 'content3');

COMMIT;
