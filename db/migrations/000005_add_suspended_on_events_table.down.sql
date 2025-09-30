DROP INDEX idx_suspended_event ON tb_events;

ALTER TABLE tb_events DROP COLUMN suspended_event;