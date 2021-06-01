CREATE TABLE IF NOT EXISTS public.messages
(
    id bigserial NOT NULL,
    text text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    created_by bigint NOT NULL,
    room_id bigint NOT NULL,
    status smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (id),
    CONSTRAINT user_key FOREIGN KEY (created_by)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT room_key FOREIGN KEY (room_id)
        REFERENCES public.rooms (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);