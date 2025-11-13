CREATE TABLE chapters
(
    id          BIGSERIAL PRIMARY KEY,
    title       VARCHAR(240) NOT NULL,
    description TEXT,
    order_index INT          NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP    NOT NULL DEFAULT NOW(),

    CONSTRAINT unique_chapter_order UNIQUE (order_index)
);

CREATE INDEX idx_chapters_order ON chapters (order_index);
