CREATE TABLE　IF NOT EXISTS races (
    id BIGINT NOT NULL AUTO_INCREMENT,
    race_date DATE NOT NULL,
    race_course_id BIGINT NOT NULL,
    race_number TINYINT UNSIGNED NOT NULL,
    race_name VARCHAR(128) NOT NULL,
    start_time TIME NOT NULL,
    surface VARCHAR(16) NOT NULL,
    distance INT UNSIGNED NOT NULL,
    direction VARCHAR(16) NOT NULL,
    weather VARCHAR(16) NOT NULL,
    track_condition VARCHAR(16) NOT NULL,
    race_conditions VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uq_races_date_course_number (race_date, race_course_id, race_number),
    CONSTRAINT fk_races_race_course
        FOREIGN KEY (race_course_id) REFERENCES race_courses (id)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
