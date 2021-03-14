-- ./migrate -path ./database/migrations/ -database postgresql://beautyshop:beautyshop456498@localhost:5432/beautyshop?sslmode=disable up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Таблица типов услуг
CREATE TABLE service_types
(
    uuid UUID NOT NULL
        CONSTRAINT service_types_pk
            PRIMARY KEY,
    name VARCHAR
);
CREATE UNIQUE INDEX service_types_unique_index ON service_types (name);

-- Таблица салонов красоты
CREATE TABLE beautyshops
(
    uuid         UUID                        NOT NULL
        CONSTRAINT beautyshops_pk
            PRIMARY KEY,
    name         VARCHAR,
    city         VARCHAR,
    address      VARCHAR,
    coordinates  POINT,
    open_hour    SMALLINT,
    close_hour   SMALLINT,
    created_date TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE UNIQUE INDEX beautyshops_unique_index ON beautyshops (name, city);

-- Таблица мастеров салона красоты
CREATE TABLE workers
(
    uuid        UUID NOT NULL
        CONSTRAINT workers_pk
            PRIMARY KEY,
    full_name   VARCHAR,
    description VARCHAR
);

-- Таблица связей салон красоты - мастер
CREATE TABLE beautyshops_workers
(
    beautyshop_uuid UUID NOT NULL,
    worker_uuid     UUID NOT NULL
);
CREATE UNIQUE INDEX beautyshops_workers_unique_index ON beautyshops_workers (beautyshop_uuid, worker_uuid);

-- Таблица связей мастер - тип услуги
CREATE TABLE workers_service_types
(
    worker_uuid       UUID NOT NULL,
    service_type_uuid UUID NOT NULL,
    price             DECIMAL(8, 2),
    duration          SMALLINT
);
CREATE UNIQUE INDEX workers_service_types_unique_index ON workers_service_types (worker_uuid, service_type_uuid);

-- Таблица записей клиентов на услуги
CREATE TABLE checkin_list
(
    uuid              UUID                        NOT NULL
        CONSTRAINT checkin_list_pk
            PRIMARY KEY,
    beautyshop_uuid   UUID                        NOT NULL,
    client_uuid       UUID                        NOT NULL,
    worker_uuid       UUID                        NOT NULL,
    service_type_uuid UUID                        NOT NULL,
    start_date        TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    end_date          TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted           BOOLEAN                     NOT NULL,
    created_date      TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- Таблица записей клиентов на стрижку (@TODO перенести в отдельный сервис)
CREATE TABLE clients
(
    uuid               UUID NOT NULL,
    full_name          VARCHAR,
    phone              VARCHAR,
    session_id         VARCHAR,
    session_private_id VARCHAR,
    salt               VARCHAR

);
CREATE UNIQUE INDEX clients_unique_index ON clients (phone);


-- Ниже - сидинг БД. @TODO Выпилить его в дальнейшем из миграций


INSERT INTO service_types (uuid, name)
VALUES ('2e0668af-ef32-4702-9bfb-16876957431d', 'Женская стрижка'),
       ('4cf6b154-e3e8-4135-915e-407943fff873', 'Мужская стрижка'),
       ('f98342b8-18bd-4634-ba7d-ab0075827fd8', 'Мужская модельная стрижка');

INSERT INTO beautyshops (uuid, name, city, address, coordinates, open_hour, close_hour, created_date)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', 'Сахар', 'Москва', 'Лубянский проезд, 19', POINT(55.755723, 37.633670), 9, 20, '2021-02-03 10:00:00'),
       ('9fbec264-1655-4ccf-a368-da30b9019c0b', 'Место красоты', 'Москва', 'Лубянский проезд, 17', NULL, 11, 18, '2021-02-03 10:05:00'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', 'Оранжевое небо', 'Москва', 'Солянка, 1/2 ст. 2', POINT(55.753857, 37.639433), 8, 22, '2021-02-03 10:10:00');

INSERT INTO workers (uuid, full_name, description)
VALUES ('42c9f442-203b-4deb-b8e7-ef2bee010494', 'Тестовая Марина Вячеславовна',
        'Очень быстрый мастер, если вы хотите быстро - то это к ней'),
       ('c380b673-ffc6-4a48-9618-9ce997a42476', 'Тожетестовая Василиса',
        'Очень нежный мастер, если вы хотите нежно - то это к ней'),
       ('900376a1-17a6-4364-bbe4-2d03b9dfe976', 'Тестовая Мария Львовна',
        'Очень дешевый мастер, если вы хотите дешево - то это к ней'),
       ('e5f22585-b722-4b15-b552-2d0243625a9d', 'Какаятотестовая Анжелика', ''),
       ('13ab06a3-3cfa-4b56-8fde-97905fc4c78f', 'Ещёоднатестоваяфамилия Тестовоеимя', 'Просто тестовое описание');

INSERT INTO workers_service_types (worker_uuid, service_type_uuid, price, duration)
VALUES ('42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', 1000, 60),
       ('42c9f442-203b-4deb-b8e7-ef2bee010494', '4cf6b154-e3e8-4135-915e-407943fff873', 1000, 30),
       ('42c9f442-203b-4deb-b8e7-ef2bee010494', 'f98342b8-18bd-4634-ba7d-ab0075827fd8', 1000, 30),
       ('c380b673-ffc6-4a48-9618-9ce997a42476', '2e0668af-ef32-4702-9bfb-16876957431d', 1000, 45);

INSERT INTO beautyshops_workers (beautyshop_uuid, worker_uuid)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', '42c9f442-203b-4deb-b8e7-ef2bee010494'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', 'c380b673-ffc6-4a48-9618-9ce997a42476'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', '900376a1-17a6-4364-bbe4-2d03b9dfe976'),
       ('9fbec264-1655-4ccf-a368-da30b9019c0b', 'e5f22585-b722-4b15-b552-2d0243625a9d'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', '13ab06a3-3cfa-4b56-8fde-97905fc4c78f');

INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, deleted, created_date)
VALUES ('0c43d9de-998f-4373-8529-f3622f8e371b',
        '73b00c6d-a503-46b2-ae50-2bf609a82973',
        '66c937fe-f857-45a6-8ed2-d6fcb88216ff',
        '42c9f442-203b-4deb-b8e7-ef2bee010494',
        '2e0668af-ef32-4702-9bfb-16876957431d',
        '2021-02-03 10:00:00',
        '2021-02-03 10:30:00',
        FALSE,
        NOW());

INSERT INTO clients (uuid, full_name, phone, session_id, session_private_id, salt)
VALUES ('66c937fe-f857-45a6-8ed2-d6fcb88216ff',
        'Тестовый чувак',
        '79991112233',
        '084d412c63f873244dd1b73edf11d1b953996d8493fd428de1714928b74914ea',
        'df3ab2de8ed5d68e4866a9c346e9a7efbd1b7b812ef3124a64333f72bf9f34f5',
        '4dde0a47095a425e755d55ac2d8ec41fb8b410a1da796c13d084b4ef66d3e875');
