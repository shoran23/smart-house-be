CREATE TABLE IF NOT EXISTS sessions (
    token VARCHAR(48) NOT NULL,
    username VARCHAR(24) NOT NULL,
    created TEXT NOT NULL,
    expiry TEXT NOT NULL,

    FOREIGN KEY (username) REFERENCES users(username),
    PRIMARY KEY(token)
);