CREATE EXTENSION IF NOT EXISTS vector;

ALTER TABLE tb_user_events
ADD COLUMN confirmation_code_event VARCHAR(7);
