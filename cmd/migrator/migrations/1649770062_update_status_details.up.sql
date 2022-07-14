begin;
alter table transactions drop constraint trx_status_details;
alter table transactions
    add constraint trx_status_details check (
                    status = 'pending' and
                    hash is null
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
        );
commit;

