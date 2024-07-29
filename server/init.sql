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
	user_id INTEGER REFERENCES users(user_id),
	todo_id INTEGER,
	title TEXT,
	content TEXT,
	PRIMARY KEY(user_id, todo_id)
);

INSERT INTO users (username, password)
VALUES ('username', 'password');

INSERT INTO todos (user_id, todo_id, title, content)
VALUES (1, 1, 'title', 'content');

COMMIT;