-- CREATE TABLE race_raw (name VARCHAR(255), json JSONB);

CREATE TABLE lingfield (
  "date" DATE,
  region VARCHAR(20), 
  course VARCHAR(255),
  off VARCHAR(100),
  race_name TEXT,
  type VARCHAR(10),
  class VARCHAR(10),
  pattern VARCHAR(255),
  rating_band VARCHAR(255),
  age_band VARCHAR(255), -- todo, potentially convert this fk or something else to make queryable. Might have to do for multiple columns.
  sex_rest VARCHAR(255),
  dist VARCHAR(255),
  dist_f VARCHAR(255),
  dist_m smallint,
  going VARCHAR(255),
  ran smallint,
  num smallint,
  pos smallint,
  draw smallint,
  ovr_btn smallint,
  btn smallint,
  horse VARCHAR(255),
  age smallint,
  sex VARCHAR(3),
  lbs smallint,
  hg VARCHAR(255),
  time INTERVAL,
  secs INTERVAL,
  "dec" VARCHAR(255),
  jockey VARCHAR(255),
  trainer VARCHAR(255),
  prize integer,
  "or" smallint,
  rpr smallint,
  sire VARCHAR(255),
  dam VARCHAR(255),
  damsire VARCHAR(255),
  owner VARCHAR(255),
  "comment" TEXT
);

--3,C,136,,1:53.98,113.98,1.44,Hector Crouch,Archie Watson,4104,77,82,Zoffany (IRE),Countess Chrissy GB,Declaration Of War,Hambleton Racing Ltd Xviii Partner,Made all - took keen hold - pushed along over 1f out - ridden and kept on inside final furlong(op 4/11)


