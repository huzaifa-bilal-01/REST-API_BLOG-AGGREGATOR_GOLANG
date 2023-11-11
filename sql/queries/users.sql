-- name: CreateUser :one
Insert into users(id,created_at,updated_at,name,apikey)
values($1,$2,$3,$4,
    encode(sha256(random()::text::bytea), 'hex')
)
returning *;

-- name: GetUserByApiKey :one
select * from users where apikey = $1;