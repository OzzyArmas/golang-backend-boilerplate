-- name: GetUser :one
SELECT * FROM users
    WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users (user_name, phone_number)
    VALUES ($1, $2) 
    RETURNING user_id;