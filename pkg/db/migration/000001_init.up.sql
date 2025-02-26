CREATE TABLE IF NOT EXISTS tokens (
    id VARCHAR(42) PRIMARY KEY,  -- Ethereum address (0x + 40 hex chars)
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    decimals INTEGER NOT NULL,
    is_stable BOOLEAN NOT NULL DEFAULT FALSE,
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


--- Pool Histories ---
CREATE TABLE IF NOT EXISTS prices (
    id BIGSERIAL,
    pool_id VARCHAR(42) NOT NULL REFERENCES pools(id),
    block_number BIGINT NOT NULL,
    block_timestamp BIGINT NOT NULL,
    sender VARCHAR(42) NOT NULL,
    recipient VARCHAR(42) NOT NULL,
    amount0 BIGINT NOT NULL,
    amount1 BIGINT NOT NULL,
    sqrt_price_x96 BIGINT NOT NULL,
    liquidity BIGINT NOT NULL,
    tick INTEGER NOT NULL,
    price_usd NUMERIC(78,18) NOT NULL,
    timestamp_utc TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, timestamp_utc)
);

-- Create a hypertable for prices
SELECT create_hypertable('prices', by_range('timestamp_utc'));


--- Seed data ---

INSERT INTO tokens (id, name, symbol, decimals, is_stable) VALUES
('0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', 'USD Coin', 'USDC', 6, TRUE),
('0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2', 'Wrapped Ether', 'WETH', 18, FALSE);

INSERT INTO pools (id, token0_id, token1_id) VALUES
('0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640', '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2');




-- -- For tracking processed blocks
-- CREATE TABLE block_processing_state (
--     pool_address VARCHAR NOT NULL,
--     last_processed_block BIGINT NOT NULL,
--     is_backfill BOOLEAN NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     PRIMARY KEY (pool_address, is_backfill)
-- );

-- -- For storing swap events
-- CREATE TABLE swap_events (
--     id SERIAL PRIMARY KEY,
--     pool_address VARCHAR NOT NULL,
--     block_number BIGINT NOT NULL,
--     transaction_hash VARCHAR NOT NULL,
--     sender VARCHAR NOT NULL,
--     recipient VARCHAR NOT NULL,
--     amount0 NUMERIC NOT NULL,
--     amount1 NUMERIC NOT NULL,
--     sqrt_price_x96 NUMERIC NOT NULL,
--     liquidity NUMERIC NOT NULL,
--     tick INTEGER NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     UNIQUE(transaction_hash, pool_address)
-- );

-- -- For tracking backfill chunks
-- CREATE TABLE backfill_chunks (
--     id SERIAL PRIMARY KEY,
--     pool_address VARCHAR NOT NULL,
--     start_block BIGINT NOT NULL,
--     end_block BIGINT NOT NULL,
--     status VARCHAR NOT NULL, -- 'pending', 'processing', 'completed', 'failed'
--     error TEXT,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     UNIQUE(pool_address, start_block, end_block)
-- );