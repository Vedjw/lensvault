CREATE TABLE users (
    pk SERIAL PRIMARY KEY,
    id UUID DEFAULT uuidv7() NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INT
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);