CREATE TABLE IF NOT EXISTS withdrawals
(
    id             uuid                     default gen_random_uuid() not null
        constraint withdrawals_pk
            primary key,
    account_number varchar(64)                                        not null,
    amount         integer                  default 0                 not null,
    bank_code      varchar(8)                                         not null,
    remark         varchar(64)                                        not null,
    created_at     timestamp with time zone default now()             not null,
    is_success     boolean                  default false             not null
);

alter table withdrawals
    owner to admin;

create unique index withdrawals_id_uindex
    on withdrawals (id);

CREATE TABLE IF NOT EXISTS bigflip_logs
(
    id               uuid    default gen_random_uuid() not null
        constraint bigflip_logs_pk
            primary key,
    transaction_id   bigint                            not null,
    amount           integer default 0                 not null,
    status           varchar(16)                       not null,
    trx_timestamp    timestamp with time zone,
    bank_code        varchar(8)                        not null,
    account_number   varchar(32)                       not null,
    beneficiary_name varchar(16)                       not null,
    remark           varchar(64)                       not null,
    receipt          text                              not null,
    time_served      timestamp with time zone,
    fee              integer default 0                 not null,
    withdrawal_id    uuid                              not null
        constraint bigflip_logs_withdrawals_id_fk
            references withdrawals
);

alter table bigflip_logs
    owner to admin;

create unique index bigflip_logs_id_uindex
    on bigflip_logs (id);

create unique index bigflip_logs_withdrawal_id_uindex
    on bigflip_logs (withdrawal_id);

CREATE TABLE IF NOT EXISTS bigflip_response
(
    id            uuid                     default gen_random_uuid() not null
        constraint bigflip_response_pk
            primary key,
    url           varchar(32)                                        not null,
    payload       jsonb,
    withdrawal_id uuid                     default gen_random_uuid() not null,
    created_at    timestamp with time zone default now()             not null,
    http_code     integer                                            not null
);

alter table bigflip_response
    owner to admin;

create unique index bigflip_response_id_uindex
    on bigflip_response (id);

