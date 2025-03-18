
CREATE TABLE IF NOT EXISTS apps (
id INTEGER PRIMARY KEY,
name TEXT NOT NULL,
secret TEXT NOT NULL
);


INSERT OR IGNORE INTO apps (id, name, secret)
VALUES (1, 'test', 'test-secret');