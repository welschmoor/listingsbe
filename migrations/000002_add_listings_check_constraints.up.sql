ALTER TABLE
  listings
ADD
  CONSTRAINT listings_price_check CHECK (price >= 0);

ALTER TABLE
  listings
ADD
  CONSTRAINT categories_length_check CHECK (
    array_length(categories, 1) BETWEEN 1
    AND 5
  );