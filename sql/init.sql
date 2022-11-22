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
    "user_id"    UUID PRIMARY KEY NOT NULL,
    "role_id"    int              NOT NULL,
    "name"       VARCHAR          NOT NULL,
    "email"      VARCHAR UNIQUE   NOT NULL,
    "password"   VARCHAR          NOT NULL,
    "created_at" timestamp        NOT NULL DEFAULT (NOW()),
    "updated_at" timestamp
);

CREATE TABLE "roles"
(
    "role_id"    serial PRIMARY KEY NOT NULL,
    "name"       VARCHAR            NOT NULL,
    "created_at" timestamp          NOT NULL DEFAULT (NOW()),
    "updated_at" timestamp
);

CREATE TABLE "debtors"
(
    "debtor_id"            UUID PRIMARY KEY NOT NULL,
    "user_id"              UUID             NOT NULL,
    "credit_health_id"     int              NOT NULL,
    "contract_tracking_id" int              NOT NULL,
    "credit_limit"         FLOAT            NOT NULL DEFAULT 0,
    "credit_used"          FLOAT            NOT NULL DEFAULT 0,
    "created_at"           timestamp        NOT NULL DEFAULT (NOW()),
    "updated_at"           timestamp
);

CREATE TABLE "credit_health_types"
(
    "credit_health_id" serial PRIMARY KEY NOT NULL,
    "name"             VARCHAR            NOT NULL,
    "created_at"       timestamp          NOT NULL DEFAULT (NOW()),
    "updated_at"       timestamp
);

CREATE TABLE "contract_tracking_types"
(
    "contract_tracking_id" serial PRIMARY KEY NOT NULL,
    "name"                 VARCHAR            NOT NULL,
    "created_at"           timestamp          NOT NULL DEFAULT (NOW()),
    "updated_at"           timestamp
);

CREATE TABLE "lendings"
(
    "lending_id"        UUID PRIMARY KEY NOT NULL,
    "debtor_id"         UUID             NOT NULL,
    "loan_period_id"    int              NOT NULL,
    "lending_status_id" int              NOT NULL,
    "amount"            float            NOT NULL,
    "created_at"        timestamp        NOT NULL DEFAULT (NOW()),
    "updated_at"        timestamp
);

CREATE TABLE "loan_periods"
(
    "loan_period_id" serial PRIMARY KEY NOT NULL,
    "duration"       int                NOT NULL,
    "percentage"     int                NOT NULL,
    "created_at"     timestamp          NOT NULL DEFAULT (NOW()),
    "updated_at"     timestamp
);

CREATE TABLE "lending_status_types"
(
    "lending_status_id" serial PRIMARY KEY NOT NULL,
    "name"              VARCHAR            NOT NULL,
    "created_at"        timestamp          NOT NULL DEFAULT (NOW()),
    "updated_at"        timestamp
);

CREATE TABLE "installments"
(
    "installment_id"        UUID PRIMARY KEY NOT NULL,
    "lending_id"            UUID             NOT NULL,
    "installment_status_id" int              NOT NULL,
    "amount"                float            NOT NULL,
    "due_date"              timestamp        NOT NULL,
    "created_at"            timestamp        NOT NULL DEFAULT (NOW()),
    "updated_at"            timestamp
);

CREATE TABLE "installment_status_types"
(
    "installment_status_id" serial PRIMARY KEY NOT NULL,
    "name"                  VARCHAR            NOT NULL,
    "created_at"            timestamp          NOT NULL DEFAULT (NOW()),
    "updated_at"            timestamp
);

CREATE TABLE "payments"
(
    "payment_id"       UUID PRIMARY KEY NOT NULL,
    "installment_id"   UUID             NOT NULL,
    "voucher_id"       UUID,
    "payment_fine"     float            NOT NULL DEFAULT 0,
    "payment_discount" float            NOT NULL DEFAULT 0,
    "payment_amount"   float            NOT NULL,
    "payment_date"     timestamp        NOT NULL DEFAULT (NOW())
);

CREATE TABLE "vouchers"
(
    "voucher_id"       UUID PRIMARY KEY NOT NULL,
    "discount_payment" int              NOT NULL,
    "active_date"      timestamp        NOT NULL,
    "expired_at"       timestamp        NOT NULL,
    "discount_quota"   int              NOT NULL
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
