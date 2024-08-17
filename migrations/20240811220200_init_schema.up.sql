-- Enable UUID generation extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create roles table
CREATE TABLE "roles" (
                         "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                         "name" varchar NOT NULL,
                         "description" text NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT now(),
                         "updated_at" timestamptz NOT NULL DEFAULT now()
);

-- Create privileges table
CREATE TABLE "privileges" (
                              "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                              "name" varchar NOT NULL UNIQUE,
                              "description" text NOT NULL,
                              "created_at" timestamptz NOT NULL DEFAULT now(),
                              "updated_at" timestamptz NOT NULL DEFAULT now()
);

-- Create role_privileges table
CREATE TABLE "role_privileges" (
                                   "role_id" uuid NOT NULL,
                                   "privilege_id" uuid NOT NULL,
                                   PRIMARY KEY ("role_id", "privilege_id"),
                                   FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE CASCADE,
                                   FOREIGN KEY ("privilege_id") REFERENCES "privileges" ("id") ON DELETE CASCADE
);

-- Create users table
CREATE TABLE "users" (
                         "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                         "name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "avatar" text,
                         "phone_number" varchar UNIQUE NOT NULL,
                         "password" varchar NOT NULL,
                         "is_verified" boolean NOT NULL DEFAULT false,
                         "role_id" uuid REFERENCES "roles" ("id") ON DELETE CASCADE NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT now(),
                         "updated_at" timestamptz NOT NULL DEFAULT now()
);

-- Create index on role_id in users table
CREATE INDEX idx_users_role_id ON users(role_id);

-- Create sessions table
CREATE TABLE "sessions" (
                            "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                            "access_token" VARCHAR(255) UNIQUE NOT NULL, -- Unique token for the session
                            "refresh_token" VARCHAR(255) UNIQUE NOT NULL, -- Unique refresh token
                            "user_id" uuid NOT NULL, -- Allow multiple sessions per user
                            "expiry_date" TIMESTAMPTZ NOT NULL, -- Use TIMESTAMPTZ for timezone support
                            "user_agent" TEXT, -- Store the user agent of the client
                            "ip_address" VARCHAR(45), -- Store the IP address of the client
                            "last_accessed_at" TIMESTAMPTZ, -- Timestamp of the last activity
                            "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
                            "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
                            FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);
