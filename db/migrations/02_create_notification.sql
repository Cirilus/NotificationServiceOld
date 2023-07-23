CREATE TABLE IF NOT EXISTS notification (
    id uuid primary key,
    title text not null,
    body text not null,
    Execution timestamp not null,
    telegram text null,
    email text null,
	is_sent_email bool default false,
	is_sent_telegram bool default false
);
