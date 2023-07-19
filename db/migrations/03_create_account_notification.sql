CREATE TABLE IF NOT EXISTS account_notification (
    account_id UUID REFERENCES account(id) ON DELETE CASCADE,
    notification_id UUID REFERENCES notification(id) ON DELETE CASCADE,
    PRIMARY KEY (account_id, notification_id)
);