CREATE TABLE IF NOT EXISTS "users"(
   id serial PRIMARY KEY,
   name varchar(255) NOT NULL,
   password VARCHAR (255) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);