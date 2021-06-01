CREATE TABLE IF NOT EXISTS public.room_users
(
    user_id bigint NOT NULL,
    room_id bigint NOT NULL,
    CONSTRAINT user_key FOREIGN KEY (user_id)
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