create table units
(
    id                   bigserial
        constraint units_pkey
            primary key,
    name                 varchar,
    geom                 bytea,
);