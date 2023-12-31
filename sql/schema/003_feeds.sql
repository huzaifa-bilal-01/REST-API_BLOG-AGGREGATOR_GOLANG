-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url text unique not null,
    user_id UUID not null references users(id)
);

-- +goose Down
DROP TABLE feeds;