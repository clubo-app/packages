CREATE TABLE IF NOT EXISTS favorite_parties (
    user_id text,
    party_id text,
    favorited_at timestamp,
    PRIMARY KEY (user_id, party_id, favorited_at)
) WITH CLUSTERING ORDER BY (party_id DESC, favorited_at DESC);

CREATE INDEX favorite_parties_by_party ON favorite_parties(party_id);
