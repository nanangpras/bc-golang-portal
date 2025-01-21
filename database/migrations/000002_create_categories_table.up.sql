CREATE TABLE IF NOT EXISTS "categories"(
   id serial PRIMARY KEY,
   title varchar(200) NOT NULL,
   slug VARCHAR (200) NOT NULL,
   created_by_id INT REFERENCES users (id) ON DELETE CASCADE,
   created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_categories_created_by_id ON categories(created_by_id);