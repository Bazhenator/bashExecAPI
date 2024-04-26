CREATE ROLE replicator WITH REPLICATION LOGIN PASSWORD 'postgres';

CREATE TABLE IF NOT EXISTS "Commands" (
        id integer PRIMARY KEY,
        command TEXT,
);