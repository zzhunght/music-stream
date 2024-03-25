CREATE TABLE "session" (
  "id" uuid PRIMARY KEY,
  "email" varchar NOT NULL,
  "client_agent" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_block" bool DEFAULT false,
  "expired_at" timestamp NOT NULL
);

CREATE TABLE "roles" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "accounts" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role_id" int NOT NULL,
  "is_verify" boolean NOT NULL DEFAULT (false),
  "secret_key" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "songs" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "thumbnail" varchar,
  "path" varchar,
  "lyrics" text,
  "duration" int,
  "release_date" date,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "favorite_songs" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "account_id" int NOT NULL,
  "song_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "song_categories" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "song_id" int NOT NULL,
  "category_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "artist" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "avatar_url" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "songs_artist" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "song_id" int NOT NULL,
  "artist_id" int NOT NULL,
  "owner" boolean NOT NULL
);

CREATE TABLE "albums" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "artist_id" int NOT NULL,
  "name" varchar NOT NULL,
  "release_date" date NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "albums_songs" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "song_id" int NOT NULL,
  "album_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "playlist" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" int NOT NULL,
  "account_id" int NOT NULL,
  "description" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "playlist_song" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "playlist_id" int NOT NULL,
  "song_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "favorite_albums" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "account_id" int NOT NULL,
  "album_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "artist_follow" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "account_id" int NOT NULL,
  "artist_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "songs" ("name");

ALTER TABLE "accounts" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "favorite_songs" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "favorite_songs" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "song_categories" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "song_categories" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "songs_artist" ADD FOREIGN KEY ("artist_id") REFERENCES "artist" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "songs_artist" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "albums" ADD FOREIGN KEY ("artist_id") REFERENCES "artist" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "albums_songs" ADD FOREIGN KEY ("album_id") REFERENCES "albums" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "albums_songs" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "playlist" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "playlist_song" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlist" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "playlist_song" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "favorite_albums" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "favorite_albums" ADD FOREIGN KEY ("album_id") REFERENCES "albums" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "artist_follow" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "artist_follow" ADD FOREIGN KEY ("artist_id") REFERENCES "artist" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
