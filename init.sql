CREATE TABLE users(
                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                      username TEXT UNIQUE NOT NULL,
                      password TEXT NOT NULL,
                      created_at TIMESTAMP NOT NULL,
                      updated_at TIMESTAMP NOT NULL
);

CREATE TABLE bookings(
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         username TEXT REFERENCES users(username),
                         start_time TIMESTAMP NOT NULL,
                         end_time TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION update_field_updatedAt()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = (CURRENT_TIMESTAMP AT TIME ZONE 'Europe/Moscow');
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
    BEFORE UPDATE OF username, password ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_field_updatedAt();
