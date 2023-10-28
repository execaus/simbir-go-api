CREATE TABLE "Account" (
	"id" serial NOT NULL,
	"username" CHARACTER VARYING NOT NULL UNIQUE,
	"password" CHARACTER VARYING NOT NULL,
	"balance" DOUBLE PRECISION NOT NULL,
	"deleted" bool NOT NULL,
	CONSTRAINT "Account_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Transport" (
	"id" serial NOT NULL,
	"owner" int NOT NULL,
	"type" CHARACTER VARYING NOT NULL,
	"can_rented" BOOLEAN NOT NULL,
	"model" CHARACTER VARYING NOT NULL,
	"color" CHARACTER VARYING NOT NULL,
	"identifier" CHARACTER VARYING NOT NULL UNIQUE,
	"description" TEXT,
	"latitude" DOUBLE PRECISION NOT NULL,
	"longitude" DOUBLE PRECISION NOT NULL,
	"minute_price" DOUBLE PRECISION,
	"day_price" DOUBLE PRECISION,
	"deleted" bool NOT NULL,
	CONSTRAINT "Transport_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "TransportType" (
	"name" CHARACTER VARYING NOT NULL,
	CONSTRAINT "TransportType_pk" PRIMARY KEY ("name")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Rent" (
	"id" serial NOT NULL,
	"account" int NOT NULL,
	"transport" int NOT NULL,
	"time_start" TIMESTAMP NOT NULL,
	"time_end" TIMESTAMP,
	"price_unit" DOUBLE PRECISION NOT NULL,
	"price_type" CHARACTER VARYING NOT NULL,
	"deleted" bool NOT NULL,
	CONSTRAINT "Rent_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "AccountRole" (
	"account" int NOT NULL,
	"role" CHARACTER VARYING NOT NULL,
	CONSTRAINT "AccountRole_pk" PRIMARY KEY ("account","role")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "TokenBlackList" (
	"token" CHARACTER VARYING NOT NULL,
	CONSTRAINT "TokenBlackList_pk" PRIMARY KEY ("token")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "RentType" (
	"name" CHARACTER VARYING NOT NULL,
	CONSTRAINT "RentType_pk" PRIMARY KEY ("name")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Role" (
	"name" CHARACTER VARYING NOT NULL,
	CONSTRAINT "Role_pk" PRIMARY KEY ("name")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "Transport" ADD CONSTRAINT "Transport_fk0" FOREIGN KEY ("owner") REFERENCES "Account"("id") ON UPDATE CASCADE;
ALTER TABLE "Transport" ADD CONSTRAINT "Transport_fk1" FOREIGN KEY ("type") REFERENCES "TransportType"("name") ON UPDATE CASCADE;


ALTER TABLE "Rent" ADD CONSTRAINT "Rent_fk0" FOREIGN KEY ("account") REFERENCES "Account"("id") ON UPDATE CASCADE;
ALTER TABLE "Rent" ADD CONSTRAINT "Rent_fk1" FOREIGN KEY ("transport") REFERENCES "Transport"("id") ON UPDATE CASCADE;
ALTER TABLE "Rent" ADD CONSTRAINT "Rent_fk2" FOREIGN KEY ("price_type") REFERENCES "RentType"("name") ON UPDATE CASCADE;

ALTER TABLE "AccountRole" ADD CONSTRAINT "AccountRole_fk0" FOREIGN KEY ("account") REFERENCES "Account"("id") ON UPDATE CASCADE;
ALTER TABLE "AccountRole" ADD CONSTRAINT "AccountRole_fk1" FOREIGN KEY ("role") REFERENCES "Role"("name") ON UPDATE CASCADE;

