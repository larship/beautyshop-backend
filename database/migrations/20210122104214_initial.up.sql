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

-- Таблица связей салон красоты - клиент-администратор
CREATE TABLE beautyshops_admins
(
    beautyshop_uuid UUID NOT NULL,
    client_uuid     UUID NOT NULL
);
CREATE UNIQUE INDEX beautyshops_admins_unique_index ON beautyshops_admins (beautyshop_uuid, client_uuid);

-- Таблица связей мастер - тип услуги
CREATE TABLE workers_service_types
(
    worker_uuid       UUID          NOT NULL,
    service_type_uuid UUID          NOT NULL,
    price             DECIMAL(8, 2) NOT NULL,
    duration          SMALLINT      NOT NULL
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
    price             DECIMAL(8, 2)               NOT NULL,
    deleted           BOOLEAN                     NOT NULL,
    created_date      TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- Таблица записей клиентов на стрижку (@TODO перенести в отдельный сервис)
CREATE TABLE clients
(
    uuid               UUID NOT NULL
        CONSTRAINT clients_pk
            PRIMARY KEY,
    full_name          VARCHAR,
    phone              VARCHAR,
    session_id         VARCHAR,
    session_private_id VARCHAR,
    salt               VARCHAR

);
CREATE UNIQUE INDEX clients_name_unique_index ON clients (phone);

-- Таблица проверочных кодов (@TODO перенести в отдельный сервис)
CREATE TABLE security_codes
(
    uuid       UUID                        NOT NULL
        CONSTRAINT security_codes_pk
            PRIMARY KEY,
    phone      VARCHAR                     NOT NULL,
    code       VARCHAR                     NOT NULL,
    status     VARCHAR                     NOT NULL,
    error_text VARCHAR                     NOT NULL,
    send_time  TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX security_codes_phone_index ON security_codes (phone);
CREATE INDEX security_codes_status_index ON security_codes (status);
CREATE INDEX security_codes_send_time_index ON security_codes (send_time);
