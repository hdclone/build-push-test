begin;
create type network_type as enum (
    'evm',
    'cosmos'
);

alter table transactions add column network_type network_type default 'evm';
commit;