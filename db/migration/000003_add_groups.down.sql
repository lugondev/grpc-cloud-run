ALTER TABLE IF EXISTS "groups"
    DROP CONSTRAINT IF EXISTS "group_name_unique";

DROP TYPE IF EXISTS "status_user_group" CASCADE;
DROP TABLE IF EXISTS "users_groups";
DROP TABLE IF EXISTS "groups";
