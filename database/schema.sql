-- VCM Medical Platform Database Schema
-- Run this in your Railway PostgreSQL database

-- Enable extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- User types
CREATE TABLE usertype (
    usertype           SMALLINT PRIMARY KEY,
    usertype_name      VARCHAR(64) NOT NULL
);

-- Users table
CREATE TABLE users (
    cd_user            SERIAL PRIMARY KEY,
    user_status        VARCHAR(24) NOT NULL DEFAULT 'Registered',
    ty_user            SMALLINT NOT NULL,
    subtype_user       SMALLINT NOT NULL DEFAULT 0,
    email              VARCHAR(320) UNIQUE NOT NULL,
    password           VARCHAR(60) NOT NULL,
    otp_code           VARCHAR(6) NOT NULL DEFAULT '',
    otp_created_at     TIMESTAMP WITH TIME ZONE DEFAULT '1970-01-01 00:00:01'::timestamp,
    first_name         VARCHAR(64) NOT NULL DEFAULT '',
    last_name          VARCHAR(64) NOT NULL DEFAULT '',
    gender             VARCHAR(16) NOT NULL DEFAULT 'Other',
    phone_number       VARCHAR(30) NOT NULL DEFAULT '',
    date_of_birth      DATE NOT NULL DEFAULT '1900-01-01',
    wechat_id          VARCHAR(64) NOT NULL DEFAULT '',
    languages          VARCHAR(128) NOT NULL DEFAULT '',
    occupation         VARCHAR(128) NOT NULL DEFAULT '',
    religion           VARCHAR(64) NOT NULL DEFAULT '',
    height_cm          SMALLINT NOT NULL DEFAULT 0,
    weight_kg          SMALLINT NOT NULL DEFAULT 0,
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ty_user) REFERENCES usertype(usertype)
);

-- Basic tables for the demo
CREATE TABLE af_psoriasis (
    cd_assessment      SERIAL PRIMARY KEY,
    cd_user            INTEGER NOT NULL,
    status             SMALLINT NOT NULL DEFAULT 0,
    cd_disease         SMALLINT NOT NULL DEFAULT 1,
    cd_product         SMALLINT NOT NULL DEFAULT 1,
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cd_user) REFERENCES users(cd_user)
);

CREATE TABLE appointments (
    cd_appointment     SERIAL PRIMARY KEY,
    cd_doctor          INTEGER NOT NULL,
    cd_user            INTEGER NOT NULL,
    appointment_date   DATE NOT NULL DEFAULT CURRENT_DATE,
    appointment_time   TIME NOT NULL DEFAULT '09:00:00',
    duration_minutes   SMALLINT NOT NULL DEFAULT 30,
    status             VARCHAR(32) NOT NULL DEFAULT 'scheduled',
    notes              TEXT NOT NULL DEFAULT '',
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cd_doctor) REFERENCES users(cd_user),
    FOREIGN KEY (cd_user) REFERENCES users(cd_user)
);

CREATE TABLE "order" (
    cd_order           SERIAL PRIMARY KEY,
    cd_user            INTEGER NOT NULL,
    total_amount       DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    status             VARCHAR(32) NOT NULL DEFAULT 'pending',
    order_reference    VARCHAR(64) NOT NULL DEFAULT '',
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cd_user) REFERENCES users(cd_user)
);

CREATE TABLE chat_room (
    cd_chat_room       SERIAL PRIMARY KEY,
    cd_patient         INTEGER NOT NULL,
    cd_staff           INTEGER DEFAULT NULL,
    cd_room_type       SMALLINT NOT NULL DEFAULT 1,
    status             VARCHAR(32) NOT NULL DEFAULT 'waiting',
    subject            VARCHAR(256) NOT NULL DEFAULT '',
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cd_patient) REFERENCES users(cd_user) ON DELETE CASCADE,
    FOREIGN KEY (cd_staff) REFERENCES users(cd_user) ON DELETE SET NULL
);

CREATE TABLE chat_message (
    cd_message         SERIAL PRIMARY KEY,
    cd_chat_room       INTEGER NOT NULL,
    cd_user            INTEGER NOT NULL,
    content            TEXT NOT NULL,
    is_read            BOOLEAN NOT NULL DEFAULT FALSE,
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cd_chat_room) REFERENCES chat_room(cd_chat_room) ON DELETE CASCADE,
    FOREIGN KEY (cd_user) REFERENCES users(cd_user) ON DELETE CASCADE
);

-- Insert initial data
INSERT INTO usertype (usertype, usertype_name) VALUES
(0, 'Client/Patient'),
(1, 'Agent'),
(2, 'Sales Channel'),
(3, 'Influencer'),
(4, 'Distributor'),
(5, 'Doctor'),
(10, 'Operator'),
(11, 'Admin'),
(12, 'Super Admin');

-- Create indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_type ON users(ty_user);
CREATE INDEX idx_users_status ON users(user_status);
