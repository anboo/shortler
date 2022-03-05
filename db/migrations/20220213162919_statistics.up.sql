CREATE TABLE click_statistics (
    id uuid NOT NULL,
    ip varchar DEFAULT NULL,
    click_at bigint,
    token varchar
);
CREATE UNIQUE INDEX click_statistics_uuid ON click_statistics (id);
CREATE INDEX click_statistics_token ON click_statistics (token);