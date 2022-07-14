alter table transactions
    drop column source_transaction_mined_at;
alter table transactions
    add column source_transaction_mined_at timestamp null;
