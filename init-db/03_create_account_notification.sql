CREATE TABLE IF NOT EXISTS account_notification (
    account_id UUID REFERENCES account(id),
    notification_id UUID REFERENCES notification(id),
    PRIMARY KEY (account_id, notification_id)
);