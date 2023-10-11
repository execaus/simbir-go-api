CREATE TABLE "Account" (
	"username" character varying NOT NULL,
	"password" character varying NOT NULL,
	"isAdmin" BOOLEAN NOT NULL,
	"balance" double precision NOT NULL,
	CONSTRAINT "Account_pk" PRIMARY KEY ("username")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "Transport" (
	"id" character varying NOT NULL,
	"ownerId" character varying NOT NULL,
	"transportType" character varying NOT NULL,
	"model" character varying NOT NULL,
	"color" character varying NOT NULL,
	"description" TEXT,
	"latitude" double precision NOT NULL,
	"longitude" double precision NOT NULL,
	"minutePrice" double precision,
	"dayPrice" double precision,
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
	"userId" character varying NOT NULL,
	"transportId" character varying NOT NULL,
	"timeStart" TIMESTAMP NOT NULL,
	"timeEnd" TIMESTAMP,
	"priceOfUnit" double precision NOT NULL,
	"priceType" character varying NOT NULL,
	"finalPrice" double precision,
	CONSTRAINT "Rent_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "Transport" ADD CONSTRAINT "Transport_fk0" FOREIGN KEY ("ownerId") REFERENCES "Account"("username");
ALTER TABLE "Transport" ADD CONSTRAINT "Transport_fk1" FOREIGN KEY ("transportType") REFERENCES "TransportType"("name");


ALTER TABLE "Rent" ADD CONSTRAINT "Rent_fk0" FOREIGN KEY ("userId") REFERENCES "Account"("username");
ALTER TABLE "Rent" ADD CONSTRAINT "Rent_fk1" FOREIGN KEY ("transportId") REFERENCES "Transport"("id");





