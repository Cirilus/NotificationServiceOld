CREATE TABLE IF NOT EXISTS notification (
    id uuid primary key,
    title text not null,
    body text not null,
    Execution date not null,
    telegram text null,
    email text null,
	is_sent bool default false
);
