DROP TABLE IF EXISTS clipboard;

-- Clipboard table
CREATE TABLE clipboard(
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ip_author INET NOT NULL,
    id_access INT NOT NULL UNIQUE,
    password VARCHAR(64),
    date_creation TIMESTAMP NOT NULL,
    date_last_update TIMESTAMP NOT NULL,
    message TEXT
);