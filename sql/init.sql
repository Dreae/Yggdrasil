CREATE TABLE IF NOT EXISTS daemon_servers (
  id VARCHAR(255) PRIMARY KEY,
  ip VARCHAR(16),
  port INT
);

CREATE TABLE IF NOT EXISTS users (
  username VARCHAR(128) PRIMARY KEY,
  password VARCHAR(255),
  level VARCHAR(16)
);

CREATE TABLE IF NOT EXISTS api_keys (
  key VARCHAR(128) PRIMARY KEY,
  username VARCHAR(128) REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS images (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS dameon_images (
  daemon VARCHAR(255) REFERENCES daemon_servers(id),
  image INT REFERENCES images(id)
);

CREATE TABLE IF NOT EXISTS game_servers (
  id VARCHAR(255) PRIMARY KEY,
  daemon VARCHAR(255) REFERENCES daemon_servers(id),
  username VARCHAR(255) REFERENCES users(username),
  image VARCHAR(255),
  args TEXT
);
