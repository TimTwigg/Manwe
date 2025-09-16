create table
    "public"."encounter" (
        "encounterid" serial primary KEY,
        "name" varchar(20),
        "description" text,
        "creationdate" date not null,
        "accesseddate" date not null,
        "campaign" integer not null,
        "started" boolean not null default 'false',
        "round" integer default '0',
        "turn" integer default '0',
        "haslair" boolean not null default 'false',
        "lairownerid" integer default '-1',
        "activeid" text,
        "username" text not null,
        "published" boolean not null default 'false',
        constraint "constraint_1" foreign KEY ("campaign", "username") references "public"."campaign" ("campaign", "username") on update CASCADE on delete CASCADE,
        constraint "constraint_2" foreign KEY ("username") references "public"."users" ("username") on update CASCADE on delete CASCADE
    )