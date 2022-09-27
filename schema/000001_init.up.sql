CREATE TABLE users (
    id TEXT PRIMARY KEY UNIQUE NOT NULL,
    balance DECIMAL CHECK ( balance >= 0 )
);

INSERT INTO users (id) VALUES ('first');
INSERT INTO users (id) VALUES ('second');
INSERT INTO users (id) VALUES ('third');
INSERT INTO users (id) VALUES ('forty');