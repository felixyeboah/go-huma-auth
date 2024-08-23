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
                            "access_token" TEXT UNIQUE NOT NULL, -- Unique token for the session
                            "refresh_token" TEXT UNIQUE NOT NULL, -- Unique refresh token
                            "user_id" uuid NOT NULL, -- Allow multiple sessions per user
                            "expiry_date" TIMESTAMPTZ NOT NULL, -- Use TIMESTAMPTZ for timezone support
                            "user_agent" TEXT, -- Store the user agent of the client
                            "ip_address" VARCHAR(45), -- Store the IP address of the client
                            "last_accessed_at" TIMESTAMPTZ, -- Timestamp of the last activity
                            "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
                            "updated_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
                            FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

-- Insert roles with descriptions
INSERT INTO roles (name, description)
VALUES
    ('user', 'Regular user with limited access to application resources.'),
    ('admin', 'Administrator with full access to all resources and management features.');


-- Insert privileges with descriptions
INSERT INTO privileges (name, description)
VALUES
    ('create_user', 'Ability to create new users.'),
    ('read_user', 'Ability to view user details.'),
    ('update_user', 'Ability to update user details.'),
    ('delete_user', 'Ability to delete users.'),
    ('create_role', 'Ability to create new roles.'),
    ('read_role', 'Ability to view role details.'),
    ('update_role', 'Ability to update role details.'),
    ('delete_role', 'Ability to delete roles.'),
    ('create_content', 'Ability to create new content or resources.'),
    ('read_content', 'Ability to view content or resources.'),
    ('update_content', 'Ability to update content or resources.'),
    ('delete_content', 'Ability to delete content or resources.'),
    ('manage_permissions', 'Ability to manage user permissions and roles.'),
    ('view_logs', 'Ability to view system logs or audit trails.'),
    ('system_health', 'Ability to view system health or status.'),
    ('manage_settings', 'Ability to manage application settings or configurations.'),
    ('generate_reports', 'Ability to generate reports based on system data.'),
    ('view_reports', 'Ability to view generated reports.');


-- Assign privileges to roles
INSERT INTO role_privileges (role_id, privilege_id)
VALUES
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'create_user')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'read_user')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'update_user')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'delete_user')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'create_role')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'read_role')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'update_role')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'delete_role')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'create_content')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'read_content')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'update_content')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'delete_content')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'manage_permissions')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'view_logs')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'system_health')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'manage_settings')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'generate_reports')),
    ((SELECT id FROM roles WHERE name = 'admin'), (SELECT id FROM privileges WHERE name = 'view_reports'));

-- Assuming 'user' role has a subset of these privileges
INSERT INTO role_privileges (role_id, privilege_id)
VALUES
    ((SELECT id FROM roles WHERE name = 'user'), (SELECT id FROM privileges WHERE name = 'read_user')),
    ((SELECT id FROM roles WHERE name = 'user'), (SELECT id FROM privileges WHERE name = 'update_user')),
    ((SELECT id FROM roles WHERE name = 'user'), (SELECT id FROM privileges WHERE name = 'read_content')),
    ((SELECT id FROM roles WHERE name = 'user'), (SELECT id FROM privileges WHERE name = 'update_content'));
