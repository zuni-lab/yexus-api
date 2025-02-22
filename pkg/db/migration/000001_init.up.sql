CREATE TABLE IF NOT EXISTS tokens (
    id VARCHAR(42) PRIMARY KEY,  -- Ethereum address (0x + 40 hex chars)
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    decimals INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pools (
    id VARCHAR(42) PRIMARY KEY,  -- Ethereum address (0x + 40 hex chars)
    token0_id VARCHAR(42) NOT NULL REFERENCES tokens(id),
    token1_id VARCHAR(42) NOT NULL REFERENCES tokens(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_pools_token0_id ON pools(token0_id);
CREATE INDEX idx_pools_token1_id ON pools(token1_id);


--- Price table ---

CREATE TABLE IF NOT EXISTS prices (
    id BIGSERIAL,
    pool_id VARCHAR(42) NOT NULL REFERENCES pools(id),
    price_usd NUMERIC(78,18) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, created_at)
);

-- Create a hypertable for prices
SELECT create_hypertable('prices', by_range('created_at'));


--- Seed data ---

INSERT INTO tokens (id, name, symbol, decimals) VALUES
('0x4200000000000000000000000000000000000006', 'Wrapped Ether', 'WETH', 18),
('0x833589fcd6edb6e08f4c7c32d4f71b54bda02913', 'USD Coin', 'USDC', 6);

INSERT INTO pools (id, token0_id, token1_id) VALUES
('0xd0b53d9277642d899df5c87a3966a349a798f224', '0x4200000000000000000000000000000000000006', '0x833589fcd6edb6e08f4c7c32d4f71b54bda02913');
