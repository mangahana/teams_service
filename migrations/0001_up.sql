CREATE TABLE teams (
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(30) NOT NULL,
  description   TEXT NOT NULL DEFAULT '',
  photo         TEXT NOT NULL DEFAULT '',
  owner_id      INTEGER NOT NULL,
  type_id       INTEGER NOT NULL,
  is_moderated  BOOLEAN NOT NULL DEFAULT FALSE,
  is_trusted    BOOLEAN NOT NULL DEFAULT FALSE,
  is_verified   BOOLEAN NOT NULL DEFAULT FALSE,
  created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE types (
  id   SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);
INSERT INTO types (name) VALUES ('Автор'), ('Баспа'), ('Аудармашы');


CREATE TABLE members (
  team_id     INTEGER NOT NULL,
  user_id     INTEGER NOT NULL,
  username    TEXT NOT NULL,
  user_photo  TEXT NOT NULL,
  permissions VARCHAR(50)[] NOT NULL DEFAULT '{}'
);
CREATE UNIQUE INDEX members_unique_key ON members (team_id, user_id);


CREATE TABLE permissions (
  slug VARCHAR(50) UNIQUE NOT NULL,
  name VARCHAR(50) NOT NULL
);
INSERT INTO permissions (slug, name) VALUES
('add_chapter', 'Тарау қосу'),
('remove_chapter', 'Тарау өшіру'),
('update_team', 'Топ ақпаратын өзгерту');




CREATE TABLE invites (
  team_id    INTEGER NOT NULL,
  user_id    INTEGER NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX invites_unique_key ON invites (team_id, user_id);