CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(15) NOT NULL UNIQUE,
    password VARCHAR(20) NOT NULL,
    role VARCHAR(7) NOT NULL DEFAULT 'chef',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(8) NOT NULL DEFAULT 'pending',
    UNIQUE(name, phone)
);

CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_phone ON users(phone);
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_users_updated_at ON users(updated_at);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_role ON users(role);


DROP TABLE IF EXISTS users;