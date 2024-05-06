CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'postgres';

CREATE TABLE IF NOT EXISTS commands (
        id SERIAL PRIMARY KEY,
        command TEXT,
        result TEXT
);