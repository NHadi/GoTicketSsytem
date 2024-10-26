IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'tickets_db')
BEGIN
    CREATE DATABASE tickets_db;
END
