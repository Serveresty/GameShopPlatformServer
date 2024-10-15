CREATE TABLE IF NOT EXISTS "users_data" (
    "user_id" SERIAL PRIMARY KEY,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "roles" (
    "role_id" SERIAL PRIMARY KEY,
    "role_name" VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "users_roles" (
    "user_id" BIGINT,
    "role_id" INT,
    CONSTRAINT fk_user
        FOREIGN KEY ("user_id") 
        REFERENCES users_data("user_id") 
        ON DELETE CASCADE,
    CONSTRAINT fk_role
        FOREIGN KEY ("role_id") 
        REFERENCES roles("role_id") 
        ON DELETE CASCADE,
    PRIMARY KEY ("user_id", "role_id")
);