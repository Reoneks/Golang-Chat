CREATE TABLE IF NOT EXISTS public.users
(
    id bigserial NOT NULL,
    first_name character varying(32) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(32) COLLATE pg_catalog."default" NOT NULL,
    email character varying(45) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    status smallint NOT NULL DEFAULT 1,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);