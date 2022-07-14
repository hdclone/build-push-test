start transaction;
alter table transactions drop constraint  trx_message;
create unique index trx_message
    on transactions (chain_id, bridge_address, receive_side, sha256(call_data));
commit;