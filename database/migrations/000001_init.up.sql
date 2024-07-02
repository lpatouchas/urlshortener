CREATE SCHEMA urlshortener;

CREATE TABLE short_urls
(
    id         serial4 NOT NULL,
    externalid varchar NOT NULL,
    long_url   varchar NOT NULL,
    created_at timestamp NULL,
    CONSTRAINT short_urls_pk PRIMARY KEY (id),
    CONSTRAINT short_urls_unique UNIQUE (externalid)
);

CREATE TABLE short_url_visits
(
    id           serial4   NOT NULL,
    short_url_id int4      NOT NULL,
    visited_at   timestamp NOT NULL,
    CONSTRAINT short_url_visits_pk PRIMARY KEY (id)
);

ALTER TABLE short_url_visits ADD CONSTRAINT short_url_visits_short_urls_fk FOREIGN KEY (short_url_id) REFERENCES short_urls (id);
CREATE INDEX short_url_visits_short_url_id_idx ON short_url_visits USING btree (short_url_id);

