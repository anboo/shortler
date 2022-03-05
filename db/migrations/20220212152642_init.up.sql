CREATE TABLE links (
    token VARCHAR NOT NULL,
    created_by_id INT DEFAULT NULL,
    link TEXT NOT NULL,
    expires_at BIGINT DEFAULT NULL,
    PRIMARY KEY(token)
);
CREATE UNIQUE INDEX uniq_links_short ON links (token);