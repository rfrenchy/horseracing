BEGIN;

CREATE TABLE IF NOT EXISTS racingpost (
        -- race_info
        "date" TEXT,
        region TEXT,
        course_id TEXT,
        course TEXT,
        race_id TEXT,
        off TEXT,
        race_name TEXT,
        "type" TEXT,
        class TEXT,
        pattern TEXT,
        rating_band TEXT,
        age_band TEXT,
        sex_rest TEXT,
        dist_m TEXT,
        going TEXT,
        surface TEXT,
        ran TEXT,

        -- runner_info
        num TEXT,
        pos TEXT,
        draw TEXT,
        ovr_btn TEXT,
        btn TEXT,
        horse_id TEXT,
        horse TEXT,
        age TEXT,
        sex TEXT,
        lbs TEXT,
        hg TEXT,
        time TEXT,
        "dec" TEXT,
        jockey_id TEXT,
        jockey TEXT,
        trainer_id TEXT,
        trainer TEXT,
        prize TEXT,
        "or" TEXT,
        rpr TEXT,
        ts TEXT,
        sire_id TEXT,
        sire TEXT,
        dam_id TEXT,
        dam TEXT,
        damsire_id TEXT,
        damsire TEXT,
        owner_id TEXT,
        owner TEXT,
        silk_url TEXT,
        "comment" TEXT
);

CREATE TABLE IF NOT EXISTS horse (
        horse_id INT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        sex VARCHAR(3) NOT NULL
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
        over_beaten INT,
        beaten INT,
        age VARCHAR(255),
        weight VARCHAR(255),
        headgear VARCHAR(255),
        time TIMESTAMP,
        odds REAL,
        jockey_id INT REFERENCES jockey (jockey_id),
        trainer_id INT REFERENCES trainer (trainer_id),
        prizemoney NUMERIC(10,2),
        official_rating REAL,
        rp_rating REAL,
        ts_rating REAL,
        owner_id INT REFERENCES owner (owner_id),
        "comment" TEXT
);

COMMIT;
