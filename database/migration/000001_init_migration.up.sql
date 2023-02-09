CREATE TABLE
    Organizers (
                   uuid UUID PRIMARY KEY,
                   email VARCHAR(256) UNIQUE,
                   first_name VARCHAR(50) NOT NULL,
                   last_name VARCHAR(50),
                   bio VARCHAR(200),
                   gender VARCHAR(1),
                   password VARCHAR(100) NOT NULL,
                   created_at DATE NOT NULL,
                   last_login DATE NOT NULL
);

CREATE TABLE
    Hackathons (
                   id BIGSERIAL PRIMARY KEY,
                   name VARCHAR(30) NOT NULL,
                   tagline VARCHAR(200) NOT NULL,
                   description TEXT NOT NULL,
                   contact_email VARCHAR(256) NOT NULL,
                   host VARCHAR(100) NOT NULL,
                   hacking_start TIMESTAMP
                          WITH
                              TIME ZONE,
                   deadline TIMESTAMP
                          WITH
                              TIME ZONE,
                   created_by UUID REFERENCES Organizers(uuid)
);

CREATE TABLE
    Participants (
                     uuid UUID PRIMARY KEY,
                     email VARCHAR(256) UNIQUE,
                     first_name VARCHAR(50) NOT NULL,
                     last_name VARCHAR(50),
                     bio VARCHAR(200),
                     gender VARCHAR(1),
                     password VARCHAR(100) NOT NULL,
                     created_at DATE NOT NULL,
                     last_login DATE NOT NULL
);

CREATE TABLE
    Projects (
                 id BIGSERIAL PRIMARY KEY,
                 name VARCHAR(100) NOT NULL,
                 description TEXT,
                 source_code TEXT,
                 video_link TEXT,
                 screenshot_link TEXT,
                 hackathon_id BIGSERIAL REFERENCES Hackathons(id) NOT NULL
);

CREATE TABLE
    Teams (
              team_id UUID PRIMARY KEY,
              hackathon_id BIGSERIAL REFERENCES Hackathons(id) NOT NULL,
              name VARCHAR(50) NOT NULL,
              members TEXT [] NOT NULL,
              project_id BIGSERIAL REFERENCES Projects(id)
);


SET
timezone = 'Asia/Calcutta';