CREATE TABLE IF NOT EXISTS race_details (
    id BIGINT NOT NULL AUTO_INCREMENT,
    race_id BIGINT NOT NULL,
    horse_number TINYINT UNSIGNED NOT NULL,
    frame_number TINYINT UNSIGNED NOT NULL,
    horse_name VARCHAR(128) NOT NULL,
    sex VARCHAR(16) NOT NULL,
    age TINYINT UNSIGNED NOT NULL,
    weight INT UNSIGNED NOT NULL,
    jockey VARCHAR(64) NOT NULL,
    stable VARCHAR(64) NOT NULL,
    body_weight SMALLINT UNSIGNED NOT NULL,
    body_weight_change SMALLINT NOT NULL,
    odds DECIMAL(8, 2) NOT NULL,
    popularity TINYINT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    KEY idx_race_details_race_id (race_id),
    UNIQUE KEY uq_race_details_race_horse_number (race_id, horse_number),
    CONSTRAINT fk_race_details_race
        FOREIGN KEY (race_id) REFERENCES races (id)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
