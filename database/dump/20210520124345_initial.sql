INSERT INTO service_types (uuid, name)
VALUES ('2e0668af-ef32-4702-9bfb-16876957431d', 'Женская стрижка'),
       ('4cf6b154-e3e8-4135-915e-407943fff873', 'Мужская стрижка'),
       ('f98342b8-18bd-4634-ba7d-ab0075827fd8', 'Мужская модельная стрижка');

INSERT INTO beautyshops (uuid, name, city, address, coordinates, open_hour, close_hour, created_date)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', 'Сахар', 'Москва', 'Лубянский проезд, 19', POINT(55.755723, 37.633670), 9, 20, '2021-02-03 10:00:00'),
       ('9fbec264-1655-4ccf-a368-da30b9019c0b', 'Место красоты', 'Москва', 'Лубянский проезд, 17', NULL, 11, 18, '2021-02-03 10:05:00'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', 'Оранжевое небо', 'Москва', 'Солянка, 1/2 ст. 2', POINT(55.753857, 37.639433), 8, 22, '2021-02-03 10:10:00'),
       ('97cdf010-0428-4c56-8d24-6abc7db87967', 'Тестовый салон', 'Москва', 'Улица имени Теста Тестовича, 123', POINT(55.732687, 37.631687), 8, 22, '2021-05-19 10:15:00');

INSERT INTO beautyshops_contacts (beautyshop_uuid, phone)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', '74957243851'),
       ('9fbec264-1655-4ccf-a368-da30b9019c0b', '79268887070'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', '74956284702'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', '74956286975'),
       ('97cdf010-0428-4c56-8d24-6abc7db87967', '71111111111');

INSERT INTO workers (uuid, full_name, description)
VALUES ('42c9f442-203b-4deb-b8e7-ef2bee010494', 'Марина Тестовая',
        'Очень быстрый мастер, если вы хотите быстро - то это к ней'),
       ('c380b673-ffc6-4a48-9618-9ce997a42476', 'Тожетестовая Василиса',
        'Очень нежный мастер, если вы хотите нежно - то это к ней'),
       ('900376a1-17a6-4364-bbe4-2d03b9dfe976', 'Мария Тестовторая',
        'Очень дешевый мастер, если вы хотите дешево - то это к ней'),
       ('e5f22585-b722-4b15-b552-2d0243625a9d', 'Анжелика Какаятотестовая', ''),
       ('13ab06a3-3cfa-4b56-8fde-97905fc4c78f', 'Тестовоеимя Ещёоднатестоваяфамилия', 'Просто тестовое описание');

INSERT INTO workers_service_types (worker_uuid, service_type_uuid, price, duration)
VALUES ('42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', 1000, 60),
       ('42c9f442-203b-4deb-b8e7-ef2bee010494', '4cf6b154-e3e8-4135-915e-407943fff873', 1000, 30),
       ('42c9f442-203b-4deb-b8e7-ef2bee010494', 'f98342b8-18bd-4634-ba7d-ab0075827fd8', 1000, 30),
       ('c380b673-ffc6-4a48-9618-9ce997a42476', '2e0668af-ef32-4702-9bfb-16876957431d', 1500, 45);

INSERT INTO beautyshops_workers (beautyshop_uuid, worker_uuid)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', '42c9f442-203b-4deb-b8e7-ef2bee010494'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', 'c380b673-ffc6-4a48-9618-9ce997a42476'),
       ('73b00c6d-a503-46b2-ae50-2bf609a82973', '900376a1-17a6-4364-bbe4-2d03b9dfe976'),
       ('9fbec264-1655-4ccf-a368-da30b9019c0b', 'e5f22585-b722-4b15-b552-2d0243625a9d'),
       ('69bf0453-d683-4457-a93c-b150b5c36e70', '13ab06a3-3cfa-4b56-8fde-97905fc4c78f');

INSERT INTO beautyshops_admins (beautyshop_uuid, client_uuid)
VALUES ('73b00c6d-a503-46b2-ae50-2bf609a82973', '66c937fe-f857-45a6-8ed2-d6fcb88216ff');

INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('0c43d9de-998f-4373-8529-f3622f8e371b', '73b00c6d-a503-46b2-ae50-2bf609a82973', '66c937fe-f857-45a6-8ed2-d6fcb88216ff', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-02-03 10:00:00.000000', '2021-02-03 10:30:00.000000', 1000.00, false, '2021-03-30 22:49:29.592961');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('12b15919-b845-4019-9bc6-78514d5c6d49', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-04-01 01:00:00.000000', '2021-04-01 01:00:00.000000', 1000.00, false, '2021-03-31 23:54:43.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('047934c2-01f0-458d-b2e2-6b5fb0373c7d', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-04-09 04:00:00.000000', '2021-04-09 04:00:00.000000', 1000.00, false, '2021-04-08 22:10:59.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('f84679a2-524f-42f1-9197-9a0d02f1ed70', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '4cf6b154-e3e8-4135-915e-407943fff873', '2021-05-06 06:00:00.000000', '2021-05-06 06:00:00.000000', 1000.00, true, '2021-05-06 01:19:27.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('c01fa211-fe35-4ab2-b2b5-ce56eff20461', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-06 03:00:00.000000', '2021-05-06 03:00:00.000000', 1000.00, true, '2021-05-06 00:37:39.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('2f3a14f3-b290-47e7-9553-629fde8fab55', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-06 03:00:00.000000', '2021-05-06 03:00:00.000000', 1000.00, true, '2021-05-06 01:21:30.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('e805547a-9d9a-4193-a97c-20322584ebe2', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', 'c380b673-ffc6-4a48-9618-9ce997a42476', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-06 03:00:00.000000', '2021-05-06 03:00:00.000000', 1500.00, false, '2021-05-06 01:29:51.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('d175ded4-8455-4d9a-b499-c0f513449adf', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', 'f98342b8-18bd-4634-ba7d-ab0075827fd8', '2021-05-06 05:30:00.000000', '2021-05-06 05:30:00.000000', 1000.00, false, '2021-05-06 01:29:59.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('fac447f9-a2fe-4361-87ac-125bb32c8638', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', 'c380b673-ffc6-4a48-9618-9ce997a42476', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-07 05:30:00.000000', '2021-05-07 05:30:00.000000', 1500.00, false, '2021-05-07 00:54:51.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('33a52351-1438-40a3-9a80-e682f63139de', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-08 04:00:00.000000', '2021-05-08 04:00:00.000000', 1000.00, false, '2021-05-07 23:44:04.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('ad49f3bf-9bdb-473b-8b06-8a05d4e6a7fb', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-11 04:30:00.000000', '2021-05-11 04:30:00.000000', 1000.00, true, '2021-05-10 23:48:33.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('47c7eda1-ec4a-49c6-8cf1-ffe164ca5edd', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-11 05:00:00.000000', '2021-05-11 05:00:00.000000', 1000.00, true, '2021-05-10 23:51:33.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('61919d52-20c0-434f-a03f-41cabb452e34', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-11 04:00:00.000000', '2021-05-11 04:00:00.000000', 1000.00, true, '2021-05-10 23:56:41.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('71c20c47-bd51-4178-aff0-badd9f0709ed', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-11 04:00:00.000000', '2021-05-11 04:00:00.000000', 1000.00, true, '2021-05-10 23:57:54.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('fe247a1a-d1e6-4ba2-94a7-396c575c2b07', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-11 05:00:00.000000', '2021-05-11 05:00:00.000000', 1000.00, false, '2021-05-11 00:00:46.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('03afc831-a804-438a-b797-4d6da6c56174', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-12 05:30:00.000000', '2021-05-12 05:30:00.000000', 1000.00, true, '2021-05-12 00:02:04.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('8f4261fb-958b-457a-98c0-c0f50c5ece50', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-12 05:30:00.000000', '2021-05-12 05:30:00.000000', 1000.00, false, '2021-05-12 04:09:09.000000');
INSERT INTO checkin_list (uuid, beautyshop_uuid, client_uuid, worker_uuid, service_type_uuid, start_date, end_date, price, deleted, created_date) VALUES ('aa32b73a-a79f-4ff1-b1c0-2b5d5f3d1ff0', '73b00c6d-a503-46b2-ae50-2bf609a82973', 'a3c2e100-cd5f-4b41-91aa-34b1dd810020', '42c9f442-203b-4deb-b8e7-ef2bee010494', '2e0668af-ef32-4702-9bfb-16876957431d', '2021-05-13 05:00:00.000000', '2021-05-13 05:00:00.000000', 1000.00, false, '2021-05-13 00:29:53.000000');

INSERT INTO clients (uuid, full_name, phone, session_id, session_private_id, salt)
VALUES ('66c937fe-f857-45a6-8ed2-d6fcb88216ff',
        'Тестовый чувак',
        '79991112233',
        '084d412c63f873244dd1b73edf11d1b953996d8493fd428de1714928b74914ea',
        'df3ab2de8ed5d68e4866a9c346e9a7efbd1b7b812ef3124a64333f72bf9f34f5',
        '4dde0a47095a425e755d55ac2d8ec41fb8b410a1da796c13d084b4ef66d3e875');
