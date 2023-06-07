CREATE TABLE IF NOT EXISTS notification (
    id uuid primary key,
    title text,
    body text,
    Execution date,
    telegram text null,
    email text null,
    assignTo uuid REFERENCES account(id) ON DELETE CASCADE null
);