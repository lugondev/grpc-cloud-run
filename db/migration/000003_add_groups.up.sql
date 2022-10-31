CREATE TABLE "groups"
(
    "id"          bigserial PRIMARY KEY,
    "group_name"  varchar     NOT NULL,
    "owner"       varchar     NOT NULL,
    "status"      boolean     NOT NULL DEFAULT true,
    "deactivated" boolean     NOT NULL DEFAULT false,
    "updated_at"  timestamptz NOT NULL DEFAULT (now()),
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT group_name_unique UNIQUE (group_name)
);

CREATE TYPE status_user_group AS ENUM ('accepted', 'rejected', 'pending');

CREATE TABLE "users_groups"
(
    "user_id"   varchar           NOT NULL,
    "group_id"  bigserial         NOT NULL,
    "weight"    numeric default 0,
    "threshold" varchar(5)        NOT NULL,
    "status"    status_user_group NOT NULL,
    PRIMARY KEY ("user_id", "group_id")
);

ALTER TABLE "users_groups"
    ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");
