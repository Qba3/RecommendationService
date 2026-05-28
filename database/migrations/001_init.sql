CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL UNIQUE,
                       created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT,
                          price NUMERIC(10,2) NOT NULL,
                          image_url TEXT,
                          created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER REFERENCES users(id),
                        created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE order_items (
                             id SERIAL PRIMARY KEY,
                             order_id INTEGER REFERENCES orders(id),
                             product_id INTEGER REFERENCES products(id),
                             quantity INTEGER NOT NULL DEFAULT 1
);