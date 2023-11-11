-- +goose Up
create table posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title text not null,
    description text,
    published_at timestamp not null,
    url text not null unique,
    feed_id UUID not null references feeds(id) on delete cascade
);

-- +goose Down
drop table posts;
