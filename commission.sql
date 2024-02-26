CREATE TABLE tcomission_rules
(
    id          bigserial primary key,
    start_range numeric(15,4) not null,
    end_range   numeric(15,4),
    value       numeric(15,4) not null,
    type_id     bigint        not null,
    profile_id  bigint  not null,
    active      boolean       not null default true,
    created_at  timestamptz   not null default current_timestamp,
    updated_at  timestamptz   not null default current_timestamp,
    deleted_at  timestamptz
);

CREATE TABLE tcomission_types
(
    id         bigserial primary key,
    name       varchar(255) not null,
    code       varchar(12)  not null,
    active     boolean      not null default true,
    created_at timestamptz  not null default current_timestamp,
    updated_at timestamptz  not null default current_timestamp,
    deleted_at timestamptz
);

CREATE TABLE tcomission_profiles
(
    id          bigserial primary key,
    name text not null,
    description text,
    created_by  int,
    updated_by  int,
    active      boolean     not null default true,
    created_at  timestamptz not null default current_timestamp,
    updated_at  timestamptz not null default current_timestamp,
    deleted_at  timestamptz
);

INSERT INTO public.tcomission_types (id, name, code)
VALUES ( 1,'Фиксированная', 'FIXED'),
       ( 2, 'Поцентная', 'PERCENT');