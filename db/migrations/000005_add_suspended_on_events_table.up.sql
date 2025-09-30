ALTER TABLE tb_events ADD COLUMN suspended_event BOOLEAN DEFAULT FALSE;

CREATE INDEX idx_suspended_event ON tb_events (suspended_event);