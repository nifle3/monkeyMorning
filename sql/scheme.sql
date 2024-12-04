CREATE TABLE MorningPolls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_name TEXT NOT NULL
);

CREATE TABLE MorningPollsOption (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    morning_poll_id INTEGER NOT NULL,
    option_tetxt TEXT NOT NULL,

    FOREIGN KEY morning_poll_id REFERENCES MorningPolls(id)
);


CREATE TABLE SendedPolls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    poll_tg_id TEXT NOT NULL
);
