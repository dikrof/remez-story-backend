CREATE TABLE progress
(
    player_id       VARCHAR(128) PRIMARY KEY,
    current_node_id BIGINT REFERENCES nodes (id) ON DELETE SET NULL,
    state           JSONB        NOT NULL DEFAULT '{"events": []}',
    decisions       JSONB        NOT NULL DEFAULT '[]',
    started_at      TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_progress_current_node ON progress (current_node_id);
CREATE INDEX idx_progress_state        ON progress USING GIN (state);
CREATE INDEX idx_progress_updated_at   ON progress (updated_at);
