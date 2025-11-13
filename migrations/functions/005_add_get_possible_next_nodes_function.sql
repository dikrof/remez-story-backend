CREATE OR REPLACE FUNCTION get_possible_next_nodes(node_row nodes)
RETURNS BIGINT[] AS $$
DECLARE
result      BIGINT[] := '{}';
    choice_item JSONB;
    cond_item   JSONB;
BEGIN
    IF node_row.next_id IS NOT NULL THEN
        result := array_append(result, node_row.next_id);
END IF;

    IF node_row.choices IS NOT NULL THEN
        FOR choice_item IN
SELECT *
FROM jsonb_array_elements(node_row.choices)
         LOOP
    result := array_append(result, (choice_item->>'to_node_id')::BIGINT);
END LOOP;
END IF;

    IF node_row.conditional IS NOT NULL THEN
        FOR cond_item IN
SELECT *
FROM jsonb_array_elements(node_row.conditional)
         LOOP
    result := array_append(result, (cond_item->>'to_node_id')::BIGINT);
END LOOP;
END IF;

RETURN result;
END;
$$ LANGUAGE plpgsql IMMUTABLE;
