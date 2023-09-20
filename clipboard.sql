DROP TABLE IF EXISTS clipboard;
DROP TABLE IF EXISTS log;

CREATE TABLE clipboard (
    id UUID PRIMARY KEY,
    port INT NOT NULL,
    message TEXT,
    password VARCHAR(72),
    ip_owner VARCHAR(15) NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    last_update TIMESTAMP NOT NULL,
    last_see TIMESTAMP NOT NULL
);

CREATE TABLE log (
    id UUID PRIMARY KEY,
    type_request INT NOT NULL,
    ip_user VARCHAR(21),
    clipboard UUID NOT NULL,
    date_request TIMESTAMP NOT NULL,
    active BOOLEAN DEFAULT FALSE
);