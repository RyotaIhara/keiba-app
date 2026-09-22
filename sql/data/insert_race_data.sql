INSERT INTO races (
    race_date,
    race_course_id,
    race_number,
    race_name,
    start_time,
    surface,
    distance,
    direction,
    weather,
    track_condition,
    race_conditions
)
SELECT
    '2026-10-04',
    rc.id,
    1,
    '2歳未勝利',
    '10:10:00',
    'turf',
    1400,
    'left',
    'sunny',
    'firm',
    '2歳未勝利'
FROM race_courses rc
WHERE rc.code = '05'
UNION ALL
SELECT
    '2026-10-04',
    rc.id,
    2,
    '3歳以上1勝クラス',
    '10:40:00',
    'dirt',
    1600,
    'left',
    'sunny',
    'firm',
    '3歳以上1勝クラス'
FROM race_courses rc
WHERE rc.code = '05'
UNION ALL
SELECT
    '2026-10-04',
    rc.id,
    1,
    '秋風特別',
    '10:05:00',
    'turf',
    1600,
    'right',
    'cloudy',
    'good',
    '3歳以上2勝クラス'
FROM race_courses rc
WHERE rc.code = '06'
UNION ALL
SELECT
    '2026-10-11',
    rc.id,
    1,
    '京都新馬戦',
    '10:00:00',
    'turf',
    1800,
    'right',
    'sunny',
    'firm',
    '2歳新馬'
FROM race_courses rc
WHERE rc.code = '08'
UNION ALL
SELECT
    '2026-10-11',
    rc.id,
    1,
    '阪神ダート特別',
    '10:15:00',
    'dirt',
    1800,
    'right',
    'rainy',
    'soft',
    '3歳以上3勝クラス'
FROM race_courses rc
WHERE rc.code = '09'
ON DUPLICATE KEY UPDATE
    race_name = VALUES(race_name),
    start_time = VALUES(start_time),
    surface = VALUES(surface),
    distance = VALUES(distance),
    direction = VALUES(direction),
    weather = VALUES(weather),
    track_condition = VALUES(track_condition),
    race_conditions = VALUES(race_conditions);
