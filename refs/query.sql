CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
);

CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    first_name  VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    role_id     INT REFERENCES roles (id),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE directories
(
    id          SERIAL PRIMARY KEY,
    name  	VARCHAR(255) NOT NULL,
    user_id     INT REFERENCES users (id),
    parent_id 	INT REFERENCES directories (id) ON DELETE CASCADE,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE files
(
    id          SERIAL PRIMARY KEY,
    fili_name  VARCHAR(255) NOT NULL,
    directory_id     INT REFERENCES directories (id),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
