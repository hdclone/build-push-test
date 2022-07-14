begin;
drop type network_type;
alter table transactions drop column network_type;
commit;