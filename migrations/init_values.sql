INSERT INTO "Role" (name) VALUES ('USER');
INSERT INTO "Role" (name) VALUES ('ADMIN');

INSERT INTO "RentType" (name) VALUES ('DAYS');
INSERT INTO "RentType" (name) VALUES ('MINUTES');

INSERT INTO "TransportType" (name) VALUES ('CAR');
INSERT INTO "TransportType" (name) VALUES ('BIKE');
INSERT INTO "TransportType" (name) VALUES ('SCOOTER');

CREATE EXTENSION earthdistance CASCADE;
