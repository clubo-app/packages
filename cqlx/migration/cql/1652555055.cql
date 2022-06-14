CREATE TABLE IF NOT EXISTS party_favorites (
    user_id text,
    party_id text,
    favorited_at timestamp,
    PRIMARY KEY (party_id, user_id)
);

CREATE MATERIALIZED VIEW IF NOT EXISTS favorite_parties_by_user AS
    SELECT * FROM party_favorites
    WHERE party_id IS NOT NULL AND user_id IS NOT NULL AND favorited_at IS NOT NULL
    PRIMARY KEY (user_id, favorited_at, party_id)
    WITH CLUSTERING ORDER BY (favorited_at DESC, party_id DESC);