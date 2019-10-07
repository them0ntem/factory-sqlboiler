CREATE TABLE episodes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    season INT8 NULL,
    num INT8 NULL,
    title STRING NULL,
    stardate DECIMAL NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE quotes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    quote STRING NULL,
    characters STRING NULL,
    stardate DECIMAL NULL,
    episode_id UUID NOT NULL REFERENCES episodes (id),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
