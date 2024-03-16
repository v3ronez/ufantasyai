CREATE TABLE IF NOT EXISTS accounts (
    id serial primary key,
    user_id uuid references auth.users,
    username  TEXT NOT NULL,
    created_at TIMESTAMP  NOT NULL DEFAULT NOW()
)