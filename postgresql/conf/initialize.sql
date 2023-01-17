CREATE DATABASE medmanage;
\c medmanage


CREATE ROLE medmanage WITH LOGIN PASSWORD 'medmanage';
GRANT ALL PRIVILEGES ON DATABASE medmanage TO medmanage;