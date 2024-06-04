-- Insert John Doe as admin
INSERT INTO users (name, phone, password, role, created_at, updated_at, status)
VALUES ('John Doe', '01700000001', '$2a$10$8tTTIoxEgIXRINozlLZYZuFNoiOLnx.6I2S5KoXlBCPfh1wYItnva', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'pending');

-- Insert Jane Doe as manager
INSERT INTO users (name, phone, password, role, created_at, updated_at, status)
VALUES ('Jane Doe', '01700000002', '$2a$10$8tTTIoxEgIXRINozlLZYZuFNoiOLnx.6I2S5KoXlBCPfh1wYItnva', 'manager', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'pending');

-- ... more user data inserts
