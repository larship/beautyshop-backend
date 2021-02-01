-- ./migrate -path /home/larship/projects/go/barbershop/database/migrations/ -database postgresql://barbershop:barbershop456498@localhost:5432/barbershop?sslmode=disable up

-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Таблица типов стрижек
CREATE TABLE haircut_types
(
    uuid UUID NOT NULL
        CONSTRAINT haircut_types_pk
            PRIMARY KEY,
    name VARCHAR
);

ALTER TABLE haircut_types
    OWNER TO barbershop;

CREATE UNIQUE INDEX haircut_types_unique_index
    ON haircut_types (name);

INSERT INTO haircut_types (uuid, name)
VALUES ('2e0668af-ef32-4702-9bfb-16876957431d', 'Женская стрижка'),
       ('4cf6b154-e3e8-4135-915e-407943fff873', 'Мужская стрижка'),
       ('f98342b8-18bd-4634-ba7d-ab0075827fd8', 'Мужская модельная стрижка');

-- Таблица парикмахерских
CREATE TABLE barbershops
(
    uuid    UUID NOT NULL
        CONSTRAINT barbershops_pk
            PRIMARY KEY,
    name    VARCHAR,
    city    VARCHAR,
    address VARCHAR
);

ALTER TABLE barbershops
    OWNER TO barbershop;

CREATE UNIQUE INDEX barbershops_unique_index
    ON barbershops (name, city);

INSERT INTO barbershops (uuid, name, city, address)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', 'Сахар', 'Москва', 'Лубянский проезд, 19'),
       ('9fbec264-1655-4ccf-a368-da30b9019c0b', 'Место красоты', 'Москва', 'Лубянский проезд, 17'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', 'Оранжевое небо', 'Москва', 'Солянка, 1/2 ст. 2');

-- Таблица связей парикмахерская - тип стрижки
CREATE TABLE barbershops_haircut_types
(
    barbershop_uuid   UUID NOT NULL,
    haircut_type_uuid UUID NOT NULL
);

CREATE UNIQUE INDEX barbershops_haircut_types_unique_index
    ON barbershops_haircut_types (barbershop_uuid, haircut_type_uuid);

INSERT INTO barbershops_haircut_types (barbershop_uuid, haircut_type_uuid)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', '2e0668af-ef32-4702-9bfb-16876957431d'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', '4cf6b154-e3e8-4135-915e-407943fff873'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', 'f98342b8-18bd-4634-ba7d-ab0075827fd8'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', '2e0668af-ef32-4702-9bfb-16876957431d');

-- Таблица мастеров-парикхамеров
CREATE TABLE hairdressers
(
    uuid    UUID NOT NULL
        CONSTRAINT hairdressers_pk
            PRIMARY KEY,
    full_name    VARCHAR
);

ALTER TABLE barbershops
    OWNER TO barbershop;

INSERT INTO hairdressers (uuid, full_name)
VALUES ('42c9f442-203b-4deb-b8e7-ef2bee010494', 'Тестовая Марина Вячеславовна'),
       ('c380b673-ffc6-4a48-9618-9ce997a42476', 'Тожетестовая Василиса'),
       ('900376a1-17a6-4364-bbe4-2d03b9dfe976', 'Тестовая Мария Львовна'),
       ('e5f22585-b722-4b15-b552-2d0243625a9d', 'Какаятотестовая Анжелика'),
       ('13ab06a3-3cfa-4b56-8fde-97905fc4c78f', 'Ещёоднатестоваяфамилия Тестовоеимя');

-- Таблица связей парикмахерская - мастер
CREATE TABLE barbershops_hairdressers
(
    barbershop_uuid   UUID NOT NULL,
    hairdresser_uuid UUID NOT NULL
);

ALTER TABLE barbershops_hairdressers
    OWNER TO barbershop;

CREATE UNIQUE INDEX barbershops_hairdressers_unique_index
    ON barbershops_hairdressers (barbershop_uuid, hairdresser_uuid);

INSERT INTO barbershops_hairdressers (barbershop_uuid, hairdresser_uuid)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', '42c9f442-203b-4deb-b8e7-ef2bee010494'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', 'c380b673-ffc6-4a48-9618-9ce997a42476'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', '900376a1-17a6-4364-bbe4-2d03b9dfe976'),
       ('9fbec264-1655-4ccf-a368-da30b9019c0b', 'e5f22585-b722-4b15-b552-2d0243625a9d'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', '13ab06a3-3cfa-4b56-8fde-97905fc4c78f');
