CREATE TABLE IF NOT EXISTS content.orders (
        id bigint PRIMARY KEY,
		status VARCHAR (15) DEFAULT 'NEW',
		accrual real,
		uploaded_at bigint NOT NULL,
		user_id integer REFERENCES content.users (id)
);