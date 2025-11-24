CREATE TABLE IF NOT EXISTS devices (
    name VARCHAR NOT NULL,
    room VARCHAR(24) NOT NULL,
    location VARCHAR(24) NOT NULL,
    device_id VARCHAR(24),
    address VARCHAR(24),
    make VARCHAR(24) NOT NULL,
    model VARCHAR(24) NOT NULL,

    FOREIGN KEY (make, model) REFERENCES device_models(make, model),
    PRIMARY KEY (name)
);