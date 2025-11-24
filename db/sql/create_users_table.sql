CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(24) NOT NULL,
    password VARCHAR(24) NOT NULL,
    privilege INTEGER NOT NULL,

    PRIMARY KEY(username)
);