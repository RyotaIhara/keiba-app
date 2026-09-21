INSERT INTO users (code, name, password)
VALUES
    ('test001', 'テスト001', 'password'),
    ('test002', 'テスト002', 'password'),
    ('test003', 'テスト003', 'password')
ON DUPLICATE KEY UPDATE
    name = VALUES(name),
    password = VALUES(password);
