CREATE TABLE IF NOT EXISTS public.rooms
(
    id bigserial NOT NULL,
    name character varying(32) NOT NULL,
    status smallint NOT NULL DEFAULT 1,
    created_at timestamp without time zone NOT NULL,
    created_by bigint NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT user_key FOREIGN KEY (created_by)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);