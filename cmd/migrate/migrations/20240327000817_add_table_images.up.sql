CREATE TABLE IF NOT EXISTS images (
    id serial primary key,
    user_id uuid references auth.users,
    prompt text not null,
    batch_id uuid not null,
    status int not null default 1,
    image_location text,
    created_at timestamp not null default now(),
    deleted boolean not null default 'false',
    deleted_at timestamp
)