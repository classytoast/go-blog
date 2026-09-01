CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,

    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,

    password_hash TEXT NOT NULL,

    avatar_ref TEXT,
    bio TEXT,

    role VARCHAR(20) NOT NULL DEFAULT 'user',

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    last_login_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);