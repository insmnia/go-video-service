CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS Category(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_at = NOW();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS Video(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    author_id uuid NOT NULL,
    category_id uuid NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    FOREIGN KEY (category_id) REFERENCES Category(id)
);

CREATE TRIGGER update_category_updated_at_column_trg BEFORE UPDATE ON Category FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_video_updated_at_column_trg BEFORE UPDATE ON Video FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
