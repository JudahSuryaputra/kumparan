CREATE TABLE IF NOT EXISTS authors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS articles (
    id UUID PRIMARY KEY,
    author_id UUID NOT NULL REFERENCES authors(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Optional index for faster full-text search
CREATE INDEX IF NOT EXISTS idx_articles_title_body ON articles USING GIN (to_tsvector('english', title || ' ' || body));

-- Optional index for filtering by author name
CREATE INDEX IF NOT EXISTS idx_authors_name ON authors (name);
