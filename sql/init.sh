#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER

  CREATE TABLE IF NOT EXISTS orders (
      uid TEXT PRIMARY KEY NOT NULL,
      track_number TEXT,
      entry TEXT,
      delivery_name TEXT,
      delivery_phone TEXT,
      delivery_zip TEXT,
      delivery_city TEXT,
      delivery_address TEXT,
      delivery_region TEXT,
      delivery_email TEXT,
      payment_transaction TEXT,
      payment_request_id TEXT,
      payment_currency TEXT,
      payment_provider TEXT,
      payment_amount INT,
      payment_dt TIMESTAMPTZ,
      payment_bank TEXT,
      payment_delivery_cost INT,
      payment_goods_total INT,
      payment_custom_fee INT,
      locale TEXT,
      internal_signature TEXT,
      customer_id TEXT,
      delivery_service TEXT,
      shardkey TEXT,
      sm_id INT,
      date_created TIMESTAMPTZ,
      oof_shard TEXT
  );


  CREATE TABLE IF NOT EXISTS items (
      order_id TEXT NOT NULL,
      chrt_id INT PRIMARY KEY NOT NULL ,
      track_number TEXT,
      price INT,
      rid TEXT,
      name TEXT,
      sale INT,
      size TEXT,
      total_price INT,
      nm_id INT,
      brand TEXT,
      status INT
  );

EOSQL
