CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'pg_pwd';

CREATE TABLE IF NOT EXISTS commands (
        id SERIAL PRIMARY KEY,
        command TEXT,
        result TEXT
);