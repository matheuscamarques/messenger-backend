CREATE TABLE "user" (
  "id" int,
  "username" int,
  "password" int,
  "status" int
);

CREATE TABLE "info_user" (
  "id" int,
  "infos" string,
  "user_id" int
);

CREATE TABLE "friend" (
  "id" int,
  "user_id" int,
  "friend_id" int
);

CREATE TABLE "embed" (
  "id" int,
  "type" string,
  "src" string
);

CREATE TABLE "messages" (
  "id" int,
  "from" int,
  "to" int,
  "embed" int
);

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "info_user" ("user_id");

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "friend" ("user_id");

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "friend" ("friend_id");

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "messages" ("from");

ALTER TABLE "user" ADD FOREIGN KEY ("id") REFERENCES "messages" ("to");

ALTER TABLE "messages" ADD FOREIGN KEY ("embed") REFERENCES "embed" ("id");
