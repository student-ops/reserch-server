CREATE DATABASE medManage;
\c medManage

CREATE ROLE medManage WITH LOGIN PASSWORD 'medManage';
GRANT ALL PRIVILEGES ON DATABASE medManage TO medManage;