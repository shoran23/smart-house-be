CREATE TABLE IF NOT EXISTS device_models (
    make VARCHAR(24) NOT NULL,
    model VARCHAR(24) NOT NULL,
    purpose TEXT,
    device_type INT NOT NULL,

    PRIMARY KEY(make, model)
);