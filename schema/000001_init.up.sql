CREATE TABLE users (
    id SERIAL NOT NULL,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    count INTEGER NOT NULL DEFAULT 0,
    CONSTRAINT pk_users_id PRIMARY KEY (id)
);

CREATE INDEX idx_users_name ON users(name);