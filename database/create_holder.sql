CREATE TABLE holder (
    proxy_wallet VARCHAR(255) NOT NULL,
    bio TEXT,
    asset VARCHAR(255) NOT NULL,
    pseudonym VARCHAR(255),
    amount NUMERIC(10, 2) NOT NULL,
    display_username_public BOOLEAN NOT NULL,
    outcome_index INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    profile_image TEXT,
    profile_image_optimized TEXT
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    market_id VARCHAR(50)
);