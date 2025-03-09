-- +goose Up
-- +goose StatementBegin
CREATE TABLE device
(
    id   integer      PRIMARY KEY,
    name varchar(255) NOT NULL
);

COMMENT ON TABLE device IS 'Хранилище узи аппаратов на которых делались снимки';
COMMENT ON COLUMN device.name IS 'Название аппарата';

CREATE TABLE mri
(
    id         uuid         PRIMARY KEY,
    projection varchar(255) NOT NULL,
    checked    boolean      NOT NULL,
    create_at  date         NOT NULL,
    patient_id uuid         NOT NULL,
    device_id  integer      NOT NULL REFERENCES device (id)
);

COMMENT ON TABLE mri IS 'Хранилище описаний и характеристик узи';
COMMENT ON COLUMN mri.projection IS 'Проекция в которой было сделано узи';
COMMENT ON COLUMN mri.patient_id IS 'Идентификатор пациента к которому относится узи';
COMMENT ON COLUMN mri.device_id IS 'Идентификатор узи аппарата на котором снято узи';

CREATE TABLE image
(
    id     uuid PRIMARY KEY,
    mri_id uuid    NOT NULL REFERENCES mri (id),
    page   integer NOT NULL
);

COMMENT ON TABLE image IS 'Хранилище кадров в узи';

CREATE TABLE node
(
    id        uuid    PRIMARY KEY,
    ai        boolean NOT NULL,
    knosp_012 real    NOT NULL,
    knosp_3  real    NOT NULL,
    knosp_4  real    NOT NULL
);

COMMENT ON TABLE node IS 'Хранилище узлов в узи';
COMMENT ON COLUMN node.ai IS 'Автор узла(нейронка ли)';
COMMENT ON COLUMN node.knosp_012 IS 'процент отношения к классу knosp_012';
COMMENT ON COLUMN node.knosp_3 IS 'процент отношения к классу knosp_3';
COMMENT ON COLUMN node.knosp_4 IS 'процент отношения к классу knosp_4';

CREATE TABLE segment
(
    id        uuid PRIMARY KEY,
    node_id   uuid NOT NULL REFERENCES node (id),
    image_id  uuid NOT NULL REFERENCES image (id),
    contor    text NOT NULL,
    knosp_012 real NOT NULL,
    knosp_3  real NOT NULL,
    knosp_4  real NOT NULL
);

COMMENT ON TABLE segment IS 'Хранилище сегментов в узи';
COMMENT ON COLUMN segment.contor IS 'контур узла (JSON)';
COMMENT ON COLUMN segment.knosp_012 IS 'процент отношения к классу knosp_012';
COMMENT ON COLUMN segment.knosp_3 IS 'процент отношения к классу knosp_3';
COMMENT ON COLUMN segment.knosp_4 IS 'процент отношения к классу knosp_4';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS node CASCADE;
DROP TABLE IF EXISTS segment CASCADE;
DROP TABLE IF EXISTS image CASCADE;
DROP TABLE IF EXISTS device CASCADE;
DROP TABLE IF EXISTS mri CASCADE;
-- +goose StatementEnd
