alter table transactions
    add constraint trx_nonce unique (
        sender_address,
        nonce
    );