CREATE TABLE events
(
    id          BIGSERIAL PRIMARY KEY,
    code        VARCHAR(64)  UNIQUE NOT NULL,
    title       VARCHAR(240) NOT NULL,
    description TEXT,
    deprecated  BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_events_code       ON events (code);
CREATE INDEX idx_events_deprecated ON events (deprecated);
