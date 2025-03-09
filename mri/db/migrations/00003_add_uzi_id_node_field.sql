-- +goose Up
-- +goose StatementBegin
ALTER TABLE node ADD mri_id uuid REFERENCES mri (id);

UPDATE node 
SET mri_id = subq.mri_id
FROM (
    SELECT DISTINCT i.mri_id, n.id
    FROM image i
    JOIN segment s ON s.image_id = i.id
    JOIN node n ON s.node_id = n.id
) as subq
WHERE node.id = subq.id;

ALTER TABLE node ALTER COLUMN mri_id SET NOT NULL;

COMMENT ON COLUMN node.mri_id IS 'Идентификатор узи';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE node
    DROP COLUMN mri_id;
-- +goose StatementEnd
