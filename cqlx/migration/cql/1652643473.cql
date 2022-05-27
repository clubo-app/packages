CREATE TABLE IF NOT EXISTS friend_relations (
    user_id text,
    friend_id text,
    accepted boolean,
    requested_at timestamp,
    accepted_at timestamp,
    PRIMARY KEY (user_id, accepted, friend_id)
);
