INSERT INTO race_courses (code, name)
VALUES
    ('tokyo', '東京競馬場'),
    ('nakayama', '中山競馬場'),
    ('kyoto', '京都競馬場'),
    ('hanshin', '阪神競馬場')
ON DUPLICATE KEY UPDATE
    name = VALUES(name);
