CREATE TABLE "Account" (
	"username" character varying NOT NULL,
	"password" character varying NOT NULL,
	"balance" double precision NOT NULL,
	CONSTRAINT "Account_pk" PRIMARY KEY ("username")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Transport" (
	"id" character varying NOT NULL,
	"owner" character varying NOT NULL,
	"transport" character varying NOT NULL,
	"model" character varying NOT NULL,
	"color" character varying NOT NULL,
	"description" TEXT,
	"latitude" double precision NOT NULL,
	"longitude" double precision NOT NULL,
	"minute_price" double precision,
	"day_price" double precision,
	CONSTRAINT "Transport_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "TransportType" (
	"name" character varying NOT NULL,
	CONSTRAINT "TransportType_pk" PRIMARY KEY ("name")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Rent" (
	"id" serial NOT NULL,
	"account" character varying NOT NULL,
	"transport" character varying NOT NULL,
	"time_start" TIMESTAMP NOT NULL,
	"time_end" TIMESTAMP,
	"price_unit" double precision NOT NULL,
	"price_type" character varying NOT NULL,
	CONSTRAINT "Rent_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Role" (
	"name" character varying NOT NULL,
	CONSTRAINT "Role_pk" PRIMARY KEY ("name")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "AccountRole" (
	"account" character varying NOT NULL,
	"role" character varying NOT NULL,
	CONSTRAINT "AccountRole_pk" PRIMARY KEY ("account","role")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "Transport" ADD CONSTRAINT "Transport_fk0" FOREIGN KEY ("owner") REFERENCES "Account"("username");
ALTER TABLE "Transport" ADD CONSTRAINT "Transport_fk1" FOREIGN KEY ("transport") REFERENCES "TransportType"("name");


ALTER TABLE "Rent" ADD CONSTRAINT "Rent_fk0" FOREIGN KEY ("account") REFERENCES "Account"("username");
ALTER TABLE "Rent" ADD CONSTRAINT "Rent_fk1" FOREIGN KEY ("transport") REFERENCES "Transport"("id");


ALTER TABLE "AccountRole" ADD CONSTRAINT "AccountRole_fk0" FOREIGN KEY ("account") REFERENCES "Account"("username");
ALTER TABLE "AccountRole" ADD CONSTRAINT "AccountRole_fk1" FOREIGN KEY ("role") REFERENCES "Role"("name");







