FROM postgres:latest

RUN DEBIAN_FRONTEND=noninteractive apt-get update \
    &&  apt-get install -y postgresql-server-dev-14 \
    &&  rm -rf /var/lib/apt/lists/*
