-- psql -U postgres -f deploy-pd.sql // run

-- DROP TABLE IF EXISTS orders;
-- DROP TABLE IF EXISTS delivery;
-- DROP TABLE IF EXISTS payments;
-- DROP TABLE IF EXISTS items_lists;
-- DROP TABLE IF EXISTS items;
-- DROP DATABASE IF EXISTS demo_service;
-- DROP ROLE IF EXISTS gaby;

CREATE DATABASE demo_service;
CREATE USER gaby WITH PASSWORD 'forza' SUPERUSER;
GRANT ALL PRIVILEGES ON DATABASE demo_service TO gaby;
ALTER DATABASE demo_service OWNER TO gaby;

\c demo_service;

CREATE TABLE orders (
    order_uid VARCHAR(50) PRIMARY KEY,
    track_number VARCHAR(50) NOT NULL,
    entry VARCHAR(50) NOT NULL,
    delivery_id INT NOT NULL, 
    payment_id INT NOT NULL,
    items_list_id INT NOT NULL,
    locale VARCHAR(50) NOT NULL,
    internal_signature VARCHAR(50),
    customer_id INT NOT NULL,
    delivery_service VARCHAR(50) NOT NULL,
    shardkey VARCHAR(50) NOT NULL,
    sm_id INT NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard VARCHAR(50) NOT NULL
); 

CREATE TABLE delivery (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    phone VARCHAR(50),
    zip VARCHAR(50),
    city VARCHAR(50),
    address VARCHAR(50),
    region VARCHAR(50),
    email VARCHAR(50)
);

CREATE TABLE payments (
    transaction VARCHAR(50) PRIMARY KEY,
    request_id VARCHAR(50),
    currency VARCHAR(50),
    provider VARCHAR(50),
    amount INT,
    payment_dt INT,
    bank    VARCHAR(50),
    delivery_cost INT,
    goods_total INT,
    custom_fee INT
);

CREATE TABLE items_lists (
    id SERIAL PRIMARY KEY,
    item_id INT
);


CREATE TABLE items(
    chrt_id SERIAL PRIMARY KEY,
    track_number INT,
    price INT,
    rid VARCHAR(50),
    name VARCHAR(50),
    sale INT,
    size VARCHAR(50),
    total_price INT,
    nm_id INT,
    brand VARCHAR(50),
    status INT
);
