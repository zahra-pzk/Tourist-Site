-- +goose Up
CREATE TABLE tourist_places (
    attaction_id INTEGER PRIMARY KEY,
    attraction_name TEXT,
    category TEXT,
    categories TEXT,
    rating FLOAT,
    reviews INTEGER,
    address TEXT,
    city TEXT,
    country TEXT,
    province TEXT,
    zipcode TEXT,
    broader_category TEXT,
    Weighted_Score FLOAT,
    Weighted_Average FLOAT,
    All_Cities TEXT,
    description TEXT,
    latitude TEXT,
    longitude TEXT
);

-- +goose Down
DROP TABLE IF EXISTS tourist_places;
