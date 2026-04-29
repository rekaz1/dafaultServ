CREATE table if not exists panels (
    id serial primary key,
    title text not null,
    description text not null
);
