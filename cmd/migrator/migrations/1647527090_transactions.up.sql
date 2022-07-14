create type transaction_status as enum (
    'pending',
    'sending',
    'sent',
    'failed',
    'confirmed'
);

create table transactions (
    id char(24) primary key not null constraint trx_id_prefix             check (left(id, 4) = 'trx_'),
        created_at timestamp            not null default current_timestamp,
          chain_id numeric              not null,
    bridge_address bytea                not null constraint trx_bridge_address_length check (length(bridge_address) = 20),
      receive_side bytea                not null constraint trx_receive_side_length   check (length(receive_side) = 20),
         call_data bytea                not null,
         signature bytea                not null,
            status transaction_status   not null default 'pending',
    sender_address bytea                not null constraint trx_sender_address_length check (length(sender_address) = 20),
        updated_at timestamp                null constraint trx_updated_at_follows    check (updated_at >= created_at),
             nonce bigint                   null constraint trx_nonce_positive        check (nonce >= 0),
         gas_limit numeric                  null constraint trx_gas_limit_positive    check (gas_limit >= 0),
         gas_price numeric                  null constraint trx_gas_price_positive    check (gas_price >= 0),
              hash bytea                    null constraint trx_hash_length           check (length(hash) = 32),
             error text                     null,

    constraint trx_nonce unique (
        sender_address,
        nonce
    ),

    constraint trx_message unique (
        chain_id,
        bridge_address,
        receive_side,
        call_data
    ),

    constraint trx_status_details check (
        status = 'pending' and
            updated_at is null and
                 nonce is null and
                  hash is null and
             gas_limit is null and
             gas_price is null and
                 error is null

        or

        status = 'sending' and
            updated_at is not null and
                 nonce is not null and
             gas_limit is not null and
             gas_price is not null and
                  hash is null and
                 error is null

        or

        status in ('sent', 'confirmed') and
            updated_at is not null and
                 nonce is not null and
             gas_limit is not null and
             gas_price is not null and
                  hash is not null and
                 error is null

        or

        status = 'failed' and
            updated_at is not null and
             gas_limit is not null and
             gas_price is not null and
                 error is not null
    )
);
