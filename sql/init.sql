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
values ('101401ce-4a0f-11ed-9772-acde48001122', 1, 'Admin', '911', 'USA', 'admin@seafund.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('b4959bd6-dbbc-4871-9bc7-bbfd4d97f1ae', 2, 'Emir', '083187115996', 'Jakarta', 'emir@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('8c0c6601-dfc0-465a-a2ac-6628410d4639', 2, 'Angela', '083187115996', 'Jakarta', 'angela@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('94bba6de-564c-40a4-ba79-1945150c74b9', 2, 'Ratu', '083187115996', 'Jakarta', 'ratu@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('816cf89e-3315-4489-8c0a-706ad75188c6', 2, 'Tafia', '083187115996', 'Jakarta', 'tafia@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('4181cea0-fcd5-4d04-8c3d-adcadacfdb2d', 2, 'Ryo', '083187115996', 'Jakarta', 'ryo@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('468dcb83-9716-4921-b31b-4ef5dadf26cb', 2, 'Richard', '083187115996', 'Jakarta', 'richard@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('957a69d7-35c8-40f9-827d-52c97b17a34b', 2, 'Ryan', '083187115996', 'Jakarta', 'ryan@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('db3558b5-2540-4bb5-8b4a-ab07675b930d', 2, 'Alvin', '083187115996', 'Jakarta', 'alvin@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('cdd64907-27de-4918-abad-f44300f9d2d0', 2, 'Aldo', '083187115996', 'Jakarta', 'aldo@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('bdb55276-5f20-4d4d-9306-d9d293b3f9fb', 2, 'Arva', '083187115996', 'Jakarta', 'arva@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('c6747a38-0463-4ac6-8e9d-f42a2cb8af72', 2, 'Hanif', '083187115996', 'Jakarta', 'hanif@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('cb89f4d8-6321-4cb7-91df-2a9ed6f22164', 2, 'Yudhis', '083187115996', 'Jakarta', 'yudhis@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('5e39e0ab-e400-4d63-b57c-fc2f87742649', 2, 'Fikri', '083187115996', 'Jakarta', 'fikri@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('627f4fc1-91f7-4d4c-a46c-8111aee9cf0c', 2, 'Julius', '083187115996', 'Jakarta', 'julius@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('3096bc4e-a8fb-4170-b08e-fcba34858de7', 2, 'Aldi', '083187115996', 'Jakarta', 'aldi@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('09772339-a9cb-45dd-9b24-7af72fd94adb', 2, 'Anang', '083187115996', 'Jakarta', 'anang@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('3580498c-7c9a-4fb8-8c08-ee07e7cda565', 2, 'Ayyub', '083187115996', 'Jakarta', 'ayyub@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('4cbda84b-9d34-494e-a294-25e728732dde', 2, 'Adit', '083187115996', 'Jakarta', 'adit@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('7735e399-54b7-4b90-9f4d-62f65ef8ceec', 2, 'Nanda', '083187115996', 'Jakarta', 'nanda@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC'),
       ('c4d46062-a4a9-4aeb-97e5-b5e4c4cf24c0', 2, 'Daniel', '083187115996', 'Jakarta', 'daniel@gmail.com',
        '$2a$10$ne0VPTKWnzVsdX7zfg1I1.MVK8RiNJDrXRf3JzoXqjaFdA3jaAGCC');

insert into "debtors" (debtor_id, user_id, credit_health_id, contract_tracking_id, credit_limit, credit_used,
                       total_delay)
values ('f8d54756-37ca-4fc4-8fa4-a822248daf59', 'b4959bd6-dbbc-4871-9bc7-bbfd4d97f1ae', 1, 5, 1000000, 0, 0),
       ('81a05d49-678f-4c11-bc8b-c3dc92d4a346', '8c0c6601-dfc0-465a-a2ac-6628410d4639', 1, 5, 1000000, 0, 0),
       ('b77059d0-e60d-4d31-b7c4-da34a42126b8', '94bba6de-564c-40a4-ba79-1945150c74b9', 1, 5, 1000000, 0, 0),
       ('5f640496-eab2-4897-b22c-e283e645dae7', '816cf89e-3315-4489-8c0a-706ad75188c6', 1, 5, 1000000, 0, 0),
       ('5abf0fb9-19a2-452a-8457-8ddfb17a8ebd', '4181cea0-fcd5-4d04-8c3d-adcadacfdb2d', 1, 5, 1000000, 0, 0),
       ('e6ecf20f-dea3-4a87-94e8-64407dd388d2', '468dcb83-9716-4921-b31b-4ef5dadf26cb', 1, 5, 1000000, 0, 0),
       ('406ba062-512b-4c7b-8e5f-e1eb3d94cb6e', '957a69d7-35c8-40f9-827d-52c97b17a34b', 1, 5, 1000000, 0, 0),
       ('4e196206-63ec-45c1-9d94-433774f78a92', 'db3558b5-2540-4bb5-8b4a-ab07675b930d', 1, 5, 1000000, 0, 0),
       ('8e95b61a-72e5-11ed-a1eb-0242ac120002', 'cdd64907-27de-4918-abad-f44300f9d2d0', 1, 5, 1000000, 0, 0),
       ('498c09ac-960d-452f-b67d-76b89e7efdac', 'c6747a38-0463-4ac6-8e9d-f42a2cb8af72', 1, 5, 1000000, 0, 0),
       ('b09cc98e-4ec9-42a0-8cfb-f423e6913728', 'cb89f4d8-6321-4cb7-91df-2a9ed6f22164', 1, 5, 1000000, 0, 0),
       ('23fa7af5-ac10-4237-8058-f7a1e8709fa8', 'cb89f4d8-6321-4cb7-91df-2a9ed6f22164', 1, 5, 1000000, 0, 0),
       ('5b1b3cc9-79fb-422b-b781-ad952f4a5f23', '5e39e0ab-e400-4d63-b57c-fc2f87742649', 1, 5, 1000000, 0, 0),
       ('a1373364-817e-41c3-b833-a9354dd3274f', '627f4fc1-91f7-4d4c-a46c-8111aee9cf0c', 1, 5, 1000000, 0, 0),
       ('9ce1c7a6-1f8e-4337-9272-453eaf60276f', '3096bc4e-a8fb-4170-b08e-fcba34858de7', 1, 5, 1000000, 0, 0),
       ('13dee731-03d9-40bc-b48f-59cdfc8545d4', '09772339-a9cb-45dd-9b24-7af72fd94adb', 1, 5, 1000000, 0, 0),
       ('51caf4f1-2095-47fb-b66d-99a7db2418de', '3580498c-7c9a-4fb8-8c08-ee07e7cda565', 1, 5, 1000000, 0, 0),
       ('78e5d4b3-2c19-47d9-9d73-8a662ca52d99', '4cbda84b-9d34-494e-a294-25e728732dde', 1, 5, 1000000, 0, 0),
       ('920b5857-7e55-4a2f-930d-98e8105930eb', 'c4d46062-a4a9-4aeb-97e5-b5e4c4cf24c0', 1, 5, 1000000, 0, 0),
       ('ec55a8af-b02c-4550-96ef-44d7b8bffea0', '7735e399-54b7-4b90-9f4d-62f65ef8ceec', 1, 5, 1000000, 0, 0);

insert into "lendings" (lending_id, debtor_id, loan_period_id, lending_status_id, name, amount)
VALUES ('fb5ab385-20a1-4c87-8e4b-900507838d86', 'f8d54756-37ca-4fc4-8fa4-a822248daf59', 1, 4, 'sbux', 1000000),
       ('1a028314-72e5-11ed-a1eb-0242ac120002', '81a05d49-678f-4c11-bc8b-c3dc92d4a346', 1, 4, 'sbux', 1000000),
       ('35e76898-e54e-4eae-ae61-8a768442d42f', 'b77059d0-e60d-4d31-b7c4-da34a42126b8', 1, 4, 'sbux', 1000000),
       ('28314064-6130-4764-8e26-358a8c8004b6', '5f640496-eab2-4897-b22c-e283e645dae7', 1, 4, 'sbux', 1000000),
       ('21989865-f009-4077-9cf8-bb611273d258', '5abf0fb9-19a2-452a-8457-8ddfb17a8ebd', 1, 4, 'sbux', 1000000),
       ('368b0f27-4ea5-42a5-9106-1c5f5a465f15', 'e6ecf20f-dea3-4a87-94e8-64407dd388d2', 1, 4, 'sbux', 1000000),
       ('a5011f09-4fa6-4945-9dbd-4ec42b4ba275', '406ba062-512b-4c7b-8e5f-e1eb3d94cb6e', 1, 4, 'sbux', 1000000),
       ('982b53c6-bd7e-4a73-bede-d275f0665f7e', '4e196206-63ec-45c1-9d94-433774f78a92', 1, 4, 'sbux', 1000000),
       ('66302fb4-2302-4d92-9c8f-4770c8a95258', '8e95b61a-72e5-11ed-a1eb-0242ac120002', 1, 4, 'sbux', 1000000),
       ('bea872f8-c548-47dc-aa4b-4bb5308442d7', '498c09ac-960d-452f-b67d-76b89e7efdac', 1, 4, 'sbux', 1000000),
       ('3e056754-dfc5-47cb-b98e-e7d986e87341', 'b09cc98e-4ec9-42a0-8cfb-f423e6913728', 1, 4, 'sbux', 1000000),
       ('2280abf5-058e-4ba1-9672-c8f2eafb026c', '23fa7af5-ac10-4237-8058-f7a1e8709fa8', 1, 4, 'sbux', 1000000),
       ('6eece99e-e356-4392-be25-de7cdeba6434', '5b1b3cc9-79fb-422b-b781-ad952f4a5f23', 1, 4, 'sbux', 1000000),
       ('1d70d7db-aea6-4c87-b00d-a033a6741855', 'a1373364-817e-41c3-b833-a9354dd3274f', 1, 4, 'sbux', 1000000),
       ('c58180b7-fe28-4e49-bd77-9af7bb844e7b', '9ce1c7a6-1f8e-4337-9272-453eaf60276f', 1, 4, 'sbux', 1000000),
       ('d5800b4d-2fa8-4cf9-80cc-47ea5420bc24', '13dee731-03d9-40bc-b48f-59cdfc8545d4', 1, 4, 'sbux', 1000000),
       ('7876fe2e-ec53-4362-9d5c-0b923af4466c', '51caf4f1-2095-47fb-b66d-99a7db2418de', 1, 4, 'sbux', 1000000),
       ('09895fc0-72e1-11ed-a1eb-0242ac120002', '78e5d4b3-2c19-47d9-9d73-8a662ca52d99', 1, 4, 'sbux', 1000000),
       ('0e3162f2-72e1-11ed-a1eb-0242ac120002', '920b5857-7e55-4a2f-930d-98e8105930eb', 1, 4, 'sbux', 1000000),
       ('1260284a-72e1-11ed-a1eb-0242ac120002', 'ec55a8af-b02c-4550-96ef-44d7b8bffea0', 1, 4, 'sbux', 1000000);

insert into "installments" (installment_id, lending_id, installment_status_id, amount, due_date)
VALUES ('897fa8c4-72e1-11ed-a1eb-0242ac120002', 'fb5ab385-20a1-4c87-8e4b-900507838d86', 2, 1000000, current_timestamp),
       ('b31f012e-72e2-11ed-a1eb-0242ac120002', '1a028314-72e5-11ed-a1eb-0242ac120002', 2, 1000000, current_timestamp),
       ('bd099cbc-72e2-11ed-a1eb-0242ac120002', '35e76898-e54e-4eae-ae61-8a768442d42f', 2, 1000000, current_timestamp),
       ('c1283100-72e2-11ed-a1eb-0242ac120002', '28314064-6130-4764-8e26-358a8c8004b6', 2, 1000000, current_timestamp),
       ('c75a75f6-72e2-11ed-a1eb-0242ac120002', '21989865-f009-4077-9cf8-bb611273d258', 2, 1000000, current_timestamp),
       ('caaa096a-72e2-11ed-a1eb-0242ac120002', '368b0f27-4ea5-42a5-9106-1c5f5a465f15', 2, 1000000, current_timestamp),
       ('cde5cf4c-72e2-11ed-a1eb-0242ac120002', 'a5011f09-4fa6-4945-9dbd-4ec42b4ba275', 2, 1000000, current_timestamp),
       ('d1431320-72e2-11ed-a1eb-0242ac120002', '982b53c6-bd7e-4a73-bede-d275f0665f7e', 2, 1000000, current_timestamp),
       ('d4ac7592-72e2-11ed-a1eb-0242ac120002', '66302fb4-2302-4d92-9c8f-4770c8a95258', 2, 1000000, current_timestamp),
       ('d7ed2116-72e2-11ed-a1eb-0242ac120002', 'bea872f8-c548-47dc-aa4b-4bb5308442d7', 2, 1000000, current_timestamp),
       ('db2a84d6-72e2-11ed-a1eb-0242ac120002', '3e056754-dfc5-47cb-b98e-e7d986e87341', 2, 1000000, current_timestamp),
       ('dfc8f7fc-72e2-11ed-a1eb-0242ac120002', '2280abf5-058e-4ba1-9672-c8f2eafb026c', 2, 1000000, current_timestamp),
       ('e3bd7284-72e2-11ed-a1eb-0242ac120002', '6eece99e-e356-4392-be25-de7cdeba6434', 2, 1000000, current_timestamp),
       ('e8829808-72e2-11ed-a1eb-0242ac120002', '1d70d7db-aea6-4c87-b00d-a033a6741855', 2, 1000000, current_timestamp),
       ('f2364200-72e2-11ed-a1eb-0242ac120002', 'c58180b7-fe28-4e49-bd77-9af7bb844e7b', 2, 1000000, current_timestamp),
       ('f8615d40-72e2-11ed-a1eb-0242ac120002', 'd5800b4d-2fa8-4cf9-80cc-47ea5420bc24', 2, 1000000, current_timestamp),
       ('fdd19c54-72e2-11ed-a1eb-0242ac120002', '7876fe2e-ec53-4362-9d5c-0b923af4466c', 2, 1000000, current_timestamp),
       ('02ad42be-72e3-11ed-a1eb-0242ac120002', '09895fc0-72e1-11ed-a1eb-0242ac120002', 2, 1000000, current_timestamp),
       ('066c6d26-72e3-11ed-a1eb-0242ac120002', '0e3162f2-72e1-11ed-a1eb-0242ac120002', 2, 1000000, current_timestamp),
       ('0a86abce-72e3-11ed-a1eb-0242ac120002', '1260284a-72e1-11ed-a1eb-0242ac120002', 2, 1000000, current_timestamp);

insert into "payments" (payment_id, installment_id, payment_fine, payment_discount, payment_amount,
                        payment_date)
VALUES ('24367fc2-72e3-11ed-a1eb-0242ac120002', '897fa8c4-72e1-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('bb6f47fc-72e3-11ed-a1eb-0242ac120002', 'b31f012e-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('bf4ecf1e-72e3-11ed-a1eb-0242ac120002', 'bd099cbc-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('c468cc52-72e3-11ed-a1eb-0242ac120002', 'c1283100-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('cc463d6a-72e3-11ed-a1eb-0242ac120002', 'c75a75f6-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('d0370080-72e3-11ed-a1eb-0242ac120002', 'caaa096a-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('d3aadf02-72e3-11ed-a1eb-0242ac120002', 'cde5cf4c-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('d733a384-72e3-11ed-a1eb-0242ac120002', 'd1431320-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('da6cfb9a-72e3-11ed-a1eb-0242ac120002', 'd4ac7592-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('de112c76-72e3-11ed-a1eb-0242ac120002', 'd7ed2116-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('e1722a8c-72e3-11ed-a1eb-0242ac120002', 'db2a84d6-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('e488c4ba-72e3-11ed-a1eb-0242ac120002', 'dfc8f7fc-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('e81c6b04-72e3-11ed-a1eb-0242ac120002', 'e3bd7284-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('ed614e18-72e3-11ed-a1eb-0242ac120002', 'e8829808-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('f1050a78-72e3-11ed-a1eb-0242ac120002', 'f2364200-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('f4793ac6-72e3-11ed-a1eb-0242ac120002', 'f8615d40-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('f791374a-72e3-11ed-a1eb-0242ac120002', 'fdd19c54-72e2-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('fa9164c4-72e3-11ed-a1eb-0242ac120002', '02ad42be-72e3-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('095b7292-72e4-11ed-a1eb-0242ac120002', '066c6d26-72e3-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp),
       ('119e058c-72e4-11ed-a1eb-0242ac120002', '0a86abce-72e3-11ed-a1eb-0242ac120002', 0, 0, 1000000,
        current_timestamp);

insert into vouchers (voucher_id, name, discount_payment, discount_quota, active_date, expire_date)
VALUES ('3f286a92-72e4-11ed-a1eb-0242ac120002', 'end year 1%', 1, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('ac1e2678-72e4-11ed-a1eb-0242ac120002', 'end year 2%', 2, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('b0258162-72e4-11ed-a1eb-0242ac120002', 'end year 3%', 3, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('b4459dea-72e4-11ed-a1eb-0242ac120002', 'end year 4%', 4, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('b8fcdfba-72e4-11ed-a1eb-0242ac120002', 'end year 5%', 5, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('be3b737e-72e4-11ed-a1eb-0242ac120002', 'end year 6%', 6, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('c299711e-72e4-11ed-a1eb-0242ac120002', 'end year 7%', 7, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('c6f073c0-72e4-11ed-a1eb-0242ac120002', 'end year 8%', 8, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('cc27f836-72e4-11ed-a1eb-0242ac120002', 'end year 9%', 9, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('cf68b4e0-72e4-11ed-a1eb-0242ac120002', 'end year 10%', 10, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('d2b6ca10-72e4-11ed-a1eb-0242ac120002', 'end year 11%', 11, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('d5ae4360-72e4-11ed-a1eb-0242ac120002', 'end year 12%', 12, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('d94e9e7a-72e4-11ed-a1eb-0242ac120002', 'end year 13%', 13, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('dcb47116-72e4-11ed-a1eb-0242ac120002', 'end year 14%', 14, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('74697bf0-72e5-11ed-a1eb-0242ac120002', 'end year 15%', 15, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('e3757360-72e4-11ed-a1eb-0242ac120002', 'end year 16%', 16, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('e66e3e26-72e4-11ed-a1eb-0242ac120002', 'end year 17%', 17, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('ecfb5dc8-72e4-11ed-a1eb-0242ac120002', 'end year 18%', 18, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('f12b17ee-72e4-11ed-a1eb-0242ac120002', 'end year 19%', 19, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700'),
       ('f45ff1a0-72e4-11ed-a1eb-0242ac120002', 'end year 20%', 20, 1, current_timestamp,
        '2022-12-31 23:59:59.999 +0700');