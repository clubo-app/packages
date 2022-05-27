CREATE TABLE IF NOT EXISTS users (
    id text PRIMARY KEY,
    username text,
    email text,
    firstname text,
    lastname text,
    avatar text,
    friend_count bigint,
    provider text,
    email_verified boolean,
    email_code text,
    password_hash text,
    role text
) WITH transactions = { 'enabled' : true };

CREATE UNIQUE INDEX IF NOT EXISTS user_by_username ON users(username);
CREATE UNIQUE INDEX IF NOT EXISTS users_by_email ON users(email);
