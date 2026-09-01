CREATE TABLE users_tokens (
    token_hash TEXT PRIMARY KEY,

    user_id BIGINT NOT NULL,

    expires_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user_tokens_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_user_tokens_user
ON users_tokens(user_id);