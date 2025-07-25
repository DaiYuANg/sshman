-- Add migration script here
CREATE TABLE IF NOT EXISTS log_entries
(
    id
    INTEGER
    PRIMARY
    KEY
    AUTOINCREMENT,
    session_id
    TEXT
    NOT
    NULL,
    command
    TEXT
    NOT
    NULL,
    output
    TEXT,
    exit_code
    INTEGER,
    created_at
    DATETIME
    DEFAULT
    CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS remote_connections
(
    id
    INTEGER
    PRIMARY
    KEY
    AUTOINCREMENT,
    name
    TEXT
    NOT
    NULL,
    host
    TEXT
    NOT
    NULL,
    port
    INTEGER
    NOT
    NULL
    DEFAULT
    22,
    username
    TEXT
    NOT
    NULL,
    auth_method
    TEXT
    NOT
    NULL
    CHECK (
    auth_method
    IN
(
    'password',
    'key'
)),
    password TEXT, -- 可选（如果是 password 登录）
    private_key TEXT, -- 可选（如果是 key 登录）
    passphrase TEXT, -- 可选（私钥密码）
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
