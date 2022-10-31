CREATE TABLE "approvals"
(
    "id"         uuid PRIMARY KEY,
    "group_id"   bigserial   NOT NULL,
    "blockchain" varchar     NOT NULL,
    "network"    varchar     NOT NULL,
    "amount"     bigint      NOT NULL,
    "status"     boolean     NOT NULL DEFAULT false,
    "expires_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "approvals"
    ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");
