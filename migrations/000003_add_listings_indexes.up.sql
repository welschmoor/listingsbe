CREATE INDEX IF NOT EXISTS listings_title_idx ON listings USING GIN (to_tsvector('simple', title)); 
CREATE INDEX IF NOT EXISTS listings_categories_idx ON listings USING GIN (categories);