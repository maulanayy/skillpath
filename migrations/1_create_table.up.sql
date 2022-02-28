
CREATE TABLE IF NOT EXISTS public.users(
    id bigserial NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL,
    token varchar ,
    info jsonb,
    created_by varchar ,
    created_at timestamptz NOT NULL DEFAULT now(), 
    deleted_at timestamptz, 
    updated_at timestamptz
);