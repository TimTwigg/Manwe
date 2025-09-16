create table
    "public"."campaignentities" (
        "id" integer,
        "username" text,
        "rowid" serial,
        "statblockid" integer not null,
        "notes" text,
        constraint "campaignentities_p_key" primary KEY ("id", "username", "rowid"),
        constraint "campaign_foreign_key" foreign KEY ("id", "username") references "public"."campaign" ("id", "username") on update CASCADE on delete CASCADE,
        constraint "username_foreign_key" foreign KEY ("username") references "public"."users" ("username") on update CASCADE on delete CASCADE,
        constraint "statblockid_foreign_key" foreign KEY ("statblockid") references "public"."statblock" ("statblockid") on update CASCADE on delete CASCADE
    )