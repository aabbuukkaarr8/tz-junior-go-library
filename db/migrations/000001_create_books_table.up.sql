CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    pages INT NOT NULL CHECK (pages > 0),
    published_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    );