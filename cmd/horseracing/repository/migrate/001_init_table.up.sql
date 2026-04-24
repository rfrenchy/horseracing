BEGIN;

CREATE TABLE IF NOT EXISTS horse (
        horse_id INT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        sex VARCHAR(3) NOT NULL,
        dam_id INT REFERENCES horse (horse_id), -- Mum
        damsire_id INT REFERENCES horse (horse_id), -- Grandad Mum's side
        sire_id INT REFERENCES horse (horse_id) -- Dad
);

CREATE TABLE IF NOT EXISTS trainer (
        trainer_id INT PRIMARY KEY,
        name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS jockey (
        jockey_id INT PRIMARY KEY,
        name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS owner (
        owner_id INT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        silk_url VARCHAR(500)
);

CREATE TABLE IF NOT EXISTS region (
        region_id INT PRIMARY KEY,
        region_code VARCHAR(5) NOT NULL,
        name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS course (
        course_id INT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        region_id INT REFERENCES region (region_id)
);

CREATE TABLE IF NOT EXISTS race (
        race_id INT PRIMARY KEY,
        name VARCHAR(255),
        "date" DATE,
        course_id INT REFERENCES course (course_id),
        off_time TIME,
        race_type VARCHAR(255),
        race_class VARCHAR(255),
        pattern VARCHAR(255),
        rating_band VARCHAR(255),
        age_band_restriction VARCHAR(255),
        sex_restriction VARCHAR(255),
       distance REAL,
        going VARCHAR(255),
        surface VARCHAR(255),
        ran INT
);

CREATE TABLE IF NOT EXISTS runner (
        horse_id INT REFERENCES horse (horse_id),
        race_id INT REFERENCES race (race_id),
        race_card_number INT,
        position INT,
        draw INT,
        over_beaten REAL,
        beaten REAL,
        age VARCHAR(255),
        weight VARCHAR(255),
        headgear VARCHAR(255),
        time TIME,
        odds REAL,
        jockey_id INT REFERENCES jockey (jockey_id),
        trainer_id INT REFERENCES trainer (trainer_id),
        prizemoney NUMERIC(10,2),
        official_rating REAL,
        rp_rating REAL,
        ts_rating REAL,
        owner_id INT REFERENCES owner (owner_id),
        "comment" TEXT
        PRIMARY KEY(horse_id, race_id)
);

CREATE TABLE IF NOT EXISTS racingpost (
        course_id INT NOT NULL REFERENCES course (course_id),
        year INT NOT NULL,
        csv TEXT NOT NULL,
        processed BOOLEAN NOT NULL
);

COMMIT;
