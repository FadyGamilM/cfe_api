CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS coffees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    region VARCHAR(200) NOT NULL,
    -- light roast, medium roast, medium-dark roast, and dark roast.    
    roast_level VARCHAR(50) NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    grind_unit INT NOT NULL,
    image VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

