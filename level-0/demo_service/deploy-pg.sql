CREATE DATABASE demo_service;
CREATE USER gaby WITH PASSWORD 'forza' SUPERUSER;
GRANT ALL PRIVILEGES ON DATABASE demo_service TO gaby;
ALTER DATABASE demo_service OWNER TO gaby;

\c demo_service;

CREATE TABLE orders (
    order_uid VARCHAR(50) PRIMARY KEY,
    track_number VARCHAR(50) NOT NULL UNIQUE,
    entry VARCHAR(50) NOT NULL,
    delivery_id INT NOT NULL,
    locale VARCHAR(50) NOT NULL,
    internal_signature VARCHAR(50),
    customer_id VARCHAR(50) NOT NULL,
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
    transaction VARCHAR(50) references orders(order_uid),
    request_id VARCHAR(50),
    currency VARCHAR(50) NOT NULL,
    provider VARCHAR(50) NOT NULL,
    amount INT NOT NULL,
    payment_dt INT,
    bank    VARCHAR(50) NOT NULL,
    delivery_cost INT,
    goods_total INT NOT NULL,
    custom_fee INT NOT NULL
);

CREATE TABLE items(
    chrt_id INT PRIMARY KEY,
    track_number VARCHAR(50) references orders(track_number),
    price INT NOT NULL,
    rid VARCHAR(50),
    name VARCHAR(50) NOT NULL,
    sale INT DEFAULT 0,
    size VARCHAR(50),
    total_price INT NOT NULL,
    nm_id INT,
    brand VARCHAR(50),
    status INT
);
