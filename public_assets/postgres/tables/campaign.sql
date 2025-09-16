create table
    "public"."campaign" (
        "id" serial,
        "username" text,
        "name" text,
        "description" text,
        "creationdate" date not null,
        "lastmodified" date not null,
        constraint "campaign_p_key" primary KEY ("id", "username"),
        constraint "username_foreign_key" foreign KEY ("username") references "public"."users" ("username") on update CASCADE on delete CASCADE
    )