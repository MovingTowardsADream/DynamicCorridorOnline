CREATE OR REPLACE FUNCTION make_uid() RETURNS text AS $$
DECLARE
new_uid text;
    done bool;
BEGIN
    done := false;
    WHILE NOT done LOOP
        new_uid := md5(''||now()::text||random()::text);
        done := NOT exists(SELECT 1 FROM users WHERE id=new_uid);
END LOOP;
RETURN new_uid;
END;
$$ LANGUAGE PLPGSQL VOLATILE;

CREATE TABLE IF NOT EXISTS users (
    id TEXT DEFAULT make_uid()::text NOT NULL UNIQUE,
    username TEXT NOT NULL UNIQUE,
    pass_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS statistic (
    id BIGSERIAL PRIMARY KEY,
    exp_value INT NOT NULL,
    user_id TEXT NOT NULL UNIQUE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

