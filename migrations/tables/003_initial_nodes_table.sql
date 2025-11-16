CREATE TABLE nodes
(
    id          BIGSERIAL PRIMARY KEY,
    chapter_id  BIGINT      NOT NULL REFERENCES chapters (id) ON DELETE CASCADE,
    scene_label VARCHAR(128),
    kind        VARCHAR(32) NOT NULL CHECK (
        kind IN (
                 'NARRATION',
                 'DIALOGUE',
                 'CHOICE',
                 'SYSTEM_NOTIFICATION',
                 'CHOICE_OPTION'
            )
        ),
    speaker     VARCHAR(128),
    text        TEXT,
    next_id     BIGINT      REFERENCES nodes (id) ON DELETE SET NULL,
    choices     JSONB,
    conditional JSONB,

    CONSTRAINT check_choice_has_choices CHECK (
        (kind != 'CHOICE')
            OR (choices IS NOT NULL AND jsonb_array_length(choices) > 0)
        ),

    CONSTRAINT check_linear_has_next CHECK (
        (kind IN ('CHOICE', 'SYSTEM_NOTIFICATION'))
            OR (next_id IS NOT NULL)
        )
);

CREATE INDEX idx_nodes_chapter_id ON nodes (chapter_id);
CREATE INDEX idx_nodes_kind ON nodes (kind);
CREATE INDEX idx_nodes_scene_label ON nodes (scene_label);
CREATE INDEX idx_nodes_next_id ON nodes (next_id);

CREATE INDEX idx_nodes_choices ON nodes USING GIN (choices);
CREATE INDEX idx_nodes_conditional ON nodes USING GIN (conditional);
