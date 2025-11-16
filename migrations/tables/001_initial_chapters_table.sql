CREATE TABLE chapters (
                          id BIGSERIAL PRIMARY KEY,
                          title VARCHAR(240) NOT NULL,
                          description TEXT,
                          order_index INT NOT NULL,

                          CONSTRAINT unique_chapter_order UNIQUE(order_index)
);

CREATE INDEX idx_chapters_order ON chapters(order_index);