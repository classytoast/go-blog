CREATE TABLE categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE posts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    author_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,

    slug VARCHAR(255) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,

    cover_image_ref TEXT,

    status VARCHAR(20) NOT NULL DEFAULT 'draft',

    views_count BIGINT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CHECK (status IN ('draft', 'published', 'archived')),

    CONSTRAINT fk_posts_author
        FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_posts_category
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
);

CREATE TABLE comments (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    post_id BIGINT NOT NULL,
    parent_id BIGINT,
    user_id BIGINT NOT NULL,

    content TEXT NOT NULL,

    likes_count BIGINT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_comments_post
        FOREIGN KEY (post_id)
        REFERENCES posts(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_comments_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_comments_parent
        FOREIGN KEY (parent_id)
        REFERENCES comments(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_posts_author
ON posts(author_id);

CREATE INDEX idx_posts_category
ON posts(category_id);

CREATE INDEX idx_comments_post
ON comments(post_id);

CREATE INDEX idx_comments_user
ON comments(user_id);

CREATE INDEX idx_comments_parent
ON comments(parent_id);