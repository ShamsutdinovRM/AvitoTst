CREATE TABLE users (
    id TEXT PRIMARY KEY UNIQUE NOT NULL,
    balance DECIMAL CHECK ( balance >= 0 )
)