CREATE TABLE IF NOT EXISTS content.withdrawals (
        id serial PRIMARY KEY,
        amount real NOT NULL,
        processed_at bigint NOT NULL,
        order_id integer REFERENCES content.orders (id),
        user_id integer REFERENCES content.users (id)
);