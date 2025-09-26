CREATE EXTENSION IF NOT EXISTS vector;

CREATE INDEX idx_events_date ON tb_events (date_event);

CREATE INDEX idx_events_valid_active ON tb_events (date_event)
    WHERE active_event IS TRUE AND canceled_event IS FALSE AND deleted_event IS FALSE;

CREATE INDEX idx_events_embeddings_hnsw ON tb_events
    USING HNSW (embeddings_event vector_l2_ops);

CREATE INDEX idx_user_events_user ON tb_user_events (fk_id_user);

CREATE INDEX idx_user_events_event ON tb_user_events (fk_id_event);

CREATE UNIQUE INDEX idx_user_events_confirmation_code ON tb_user_events (confirmation_code_event);