CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       first_name VARCHAR(100),
                       last_name VARCHAR(100),
                       date_of_birth DATE,
                       gender VARCHAR(10),
                       phone_number VARCHAR(20),
                       address TEXT,
                       is_active BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE roles (
                       role_id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL UNIQUE,
                       description TEXT
);

CREATE TABLE user_roles (
                            user_role_id SERIAL PRIMARY KEY,
                            user_id INT NOT NULL,
                            role_id INT NOT NULL,
                            FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
                            FOREIGN KEY (role_id) REFERENCES roles(role_id) ON DELETE CASCADE
);

CREATE TABLE tokens (
                        token_id SERIAL PRIMARY KEY,
                        user_id INT,
                        token VARCHAR(255) NOT NULL,
                        expiration TIMESTAMP NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE activity_logs (
                               log_id SERIAL PRIMARY KEY,
                               user_id INT, -- Tidak ada referensi langsung ke basis data User Service
                               activity_type VARCHAR(50) NOT NULL,
                               timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               ip_address VARCHAR(50)
);
