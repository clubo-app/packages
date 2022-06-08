CREATE TABLE IF NOT EXISTS friend_relations (
    user_id text,
    friend_id text,
    accepted boolean,
    requested_at timestamp,
    accepted_at timestamp,
    PRIMARY KEY (user_id, accepted, friend_id)
);

CREATE MATERIALIZED VIEW friend_relations_by_friend AS
        SELECT * FROM friend_relations
        WHERE user_id IS NOT NULL AND friend_id IS NOT NULL AND accepted IS NOT NULL 
        PRIMARY KEY(friend_id, accepted, user_id);
