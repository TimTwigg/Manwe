select
    'drop table ' || name || ';' as cmd
from
    sqlite_master
where
    type = 'table'
    and name <> 'sqlite_sequence'
union all
select
    'drop view ' || name || ';' as cmd
from
    sqlite_master
where
    type = 'view';