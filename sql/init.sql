CREATE
DATABASE p2p_db_emir;

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS debtors CASCADE;
DROP TABLE IF EXISTS credit_health_types CASCADE;
DROP TABLE IF EXISTS contract_tracking_types CASCADE;
DROP TABLE IF EXISTS lendings CASCADE;
DROP TABLE IF EXISTS loan_periods CASCADE;
DROP TABLE IF EXISTS lending_status_types CASCADE;
DROP TABLE IF EXISTS installments CASCADE;
DROP TABLE IF EXISTS installment_status_types CASCADE;
DROP TABLE IF EXISTS payments CASCADE;
DROP TABLE IF EXISTS vouchers CASCADE;

CREATE TABLE "users"
(
    "user_id"      UUID PRIMARY KEY NOT NULL,
    "role_id"      int              NOT NULL,
    "name"         VARCHAR          NOT NULL,
    "phone_number" VARCHAR          NOT NULL,
    "address"      TEXT             NOT NULL,
    "email"        VARCHAR UNIQUE   NOT NULL,
    "password"     VARCHAR          NOT NULL,
    "created_at"   timestamptz      NOT NULL DEFAULT (NOW()),
    "updated_at"   timestamptz
);

CREATE TABLE "roles"
(
    "role_id"    serial PRIMARY KEY NOT NULL,
    "name"       VARCHAR            NOT NULL,
    "created_at" timestamptz        NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz
);

CREATE TABLE "debtors"
(
    "debtor_id"            UUID PRIMARY KEY NOT NULL,
    "user_id"              UUID             NOT NULL,
    "credit_health_id"     int              NOT NULL,
    "contract_tracking_id" int              NOT NULL,
    "credit_limit"         FLOAT            NOT NULL DEFAULT 0,
    "credit_used"          FLOAT            NOT NULL DEFAULT 0,
    "total_delay"          int              NOT NULL DEFAULT 0,
    "created_at"           timestamptz      NOT NULL DEFAULT (NOW()),
    "updated_at"           timestamptz
);

CREATE TABLE "credit_health_types"
(
    "credit_health_id" serial PRIMARY KEY NOT NULL,
    "name"             VARCHAR            NOT NULL,
    "created_at"       timestamptz        NOT NULL DEFAULT (NOW()),
    "updated_at"       timestamptz
);

CREATE TABLE "contract_tracking_types"
(
    "contract_tracking_id" serial PRIMARY KEY NOT NULL,
    "name"                 VARCHAR            NOT NULL,
    "created_at"           timestamptz        NOT NULL DEFAULT (NOW()),
    "updated_at"           timestamptz
);

CREATE TABLE "lendings"
(
    "lending_id"        UUID PRIMARY KEY NOT NULL,
    "debtor_id"         UUID             NOT NULL,
    "loan_period_id"    int              NOT NULL,
    "lending_status_id" int              NOT NULL,
    "name"              VARCHAR          NOT NULL,
    "amount"            float            NOT NULL,
    "created_at"        timestamptz      NOT NULL DEFAULT (NOW()),
    "updated_at"        timestamptz
);

CREATE TABLE "loan_periods"
(
    "loan_period_id" serial PRIMARY KEY NOT NULL,
    "duration"       int                NOT NULL,
    "percentage"     int                NOT NULL,
    "created_at"     timestamptz        NOT NULL DEFAULT (NOW()),
    "updated_at"     timestamptz
);

CREATE TABLE "lending_status_types"
(
    "lending_status_id" serial PRIMARY KEY NOT NULL,
    "name"              VARCHAR            NOT NULL,
    "created_at"        timestamptz        NOT NULL DEFAULT (NOW()),
    "updated_at"        timestamptz
);

CREATE TABLE "installments"
(
    "installment_id"        UUID PRIMARY KEY NOT NULL,
    "lending_id"            UUID             NOT NULL,
    "installment_status_id" int              NOT NULL,
    "amount"                float            NOT NULL,
    "due_date"              timestamptz      NOT NULL,
    "created_at"            timestamptz      NOT NULL DEFAULT (NOW()),
    "updated_at"            timestamptz
);

CREATE TABLE "installment_status_types"
(
    "installment_status_id" serial PRIMARY KEY NOT NULL,
    "name"                  VARCHAR            NOT NULL,
    "created_at"            timestamptz        NOT NULL DEFAULT (NOW()),
    "updated_at"            timestamptz
);

CREATE TABLE "payments"
(
    "payment_id"       UUID PRIMARY KEY NOT NULL,
    "installment_id"   UUID             NOT NULL,
    "voucher_id"       UUID,
    "payment_fine"     float            NOT NULL DEFAULT 0,
    "payment_discount" float            NOT NULL DEFAULT 0,
    "payment_amount"   float            NOT NULL,
    "payment_date"     timestamptz      NOT NULL DEFAULT (NOW())
);

CREATE TABLE "vouchers"
(
    "voucher_id"       UUID PRIMARY KEY NOT NULL,
    "name"             VARCHAR          NOT NULL,
    "discount_payment" int              NOT NULL,
    "discount_quota"   int              NOT NULL,
    "active_date"      timestamptz      NOT NULL,
    "expire_date"      timestamptz      NOT NULL,
    "created_at"       timestamptz      NOT NULL DEFAULT (NOW()),
    "updated_at"       timestamptz,
    "deleted_at"       timestamptz
);

ALTER TABLE "debtors"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "users"
    ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("role_id");

ALTER TABLE "lendings"
    ADD FOREIGN KEY ("debtor_id") REFERENCES "debtors" ("debtor_id");

ALTER TABLE "debtors"
    ADD FOREIGN KEY ("credit_health_id") REFERENCES "credit_health_types" ("credit_health_id");

ALTER TABLE "debtors"
    ADD FOREIGN KEY ("contract_tracking_id") REFERENCES "contract_tracking_types" ("contract_tracking_id");

ALTER TABLE "installments"
    ADD FOREIGN KEY ("lending_id") REFERENCES "lendings" ("lending_id");

ALTER TABLE "lendings"
    ADD FOREIGN KEY ("lending_status_id") REFERENCES "lending_status_types" ("lending_status_id");

ALTER TABLE "lendings"
    ADD FOREIGN KEY ("loan_period_id") REFERENCES "loan_periods" ("loan_period_id");

ALTER TABLE "payments"
    ADD FOREIGN KEY ("installment_id") REFERENCES "installments" ("installment_id");

ALTER TABLE "installments"
    ADD FOREIGN KEY ("installment_status_id") REFERENCES "installment_status_types" ("installment_status_id");

ALTER TABLE "payments"
    ADD FOREIGN KEY ("voucher_id") REFERENCES "vouchers" ("voucher_id");

INSERT INTO "roles" (name)
VALUES ('admin'),
       ('user');

INSERT INTO "credit_health_types" (name)
VALUES ('good'),
       ('warning'),
       ('blocked');

INSERT INTO "contract_tracking_types" (name)
VALUES ('no contract yet'),
       ('the contract was given to the expedition partner'),
       ('contract in delivery'),
       ('contract accepted by user'),
       ('confirmed contract');

insert into "loan_periods" (duration, percentage)
values (1, 100),
       (3, 105),
       (6, 110),
       (12, 120),
       (18, 130),
       (24, 150);

insert into "lending_status_types" (name)
values ('new'),
       ('approved'),
       ('on progress'),
       ('paid'),
       ('reject');

insert into "installment_status_types" (name)
values ('on progress'),
       ('paid');

-- // password: Tested8*
insert into "users" (user_id, role_id, name, phone_number, address, email, password)
values ('101401ce-4a0f-11ed-9772-acde48001122', 1, 'admin', '911', 'USA', 'admin@seafund.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC');