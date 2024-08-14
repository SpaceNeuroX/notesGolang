CREATE TABLE notes (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       content TEXT NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT NOW()
);
