CREATE TABLE IF NOT EXISTS panels (
    id serial primary key,
    title text not null,
    description text not null
);

CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    email text not null unique,
    password_hash text not null,
    role text not null default 'user'
);

CREATE TABLE IF NOT EXISTS sessions (
    token text primary key,
    user_id int not null references users(id) on delete cascade,
    expires_at timestamptz not null
);
