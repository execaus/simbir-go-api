INSERT INTO "Role" (name) VALUES ('USER');
INSERT INTO "Role" (name) VALUES ('ADMIN');

INSERT INTO "RentType" (name) VALUES ('DAYS');
INSERT INTO "RentType" (name) VALUES ('MINUTES');

INSERT INTO "TransportType" (name) VALUES ('CAR');
INSERT INTO "TransportType" (name) VALUES ('BIKE');
INSERT INTO "TransportType" (name) VALUES ('SCOOTER');

INSERT INTO "Account" (id, username, "password", balance, deleted)
VALUES (0, 'admin', '$2a$10$y1ydJezJeC7ZfQejQzBY8.lMhe17plLGI2D6eOuuPV6b8eqOI.2bm', 0, false);

INSERT INTO "AccountRole" (account, "role") VALUES (0, 'USER');
INSERT INTO "AccountRole" (account, "role") VALUES (0, 'ADMIN');

CREATE EXTENSION earthdistance CASCADE;
