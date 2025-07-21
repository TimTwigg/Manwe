CREATE VIEW
    actiondamagev AS
SELECT
    statblockid,
    actionid,
    coalesce(amount, '') AS amount,
    coalesce(type, '') AS type,
    altdamageactive,
    coalesce(amount2, '') AS amount2,
    coalesce(type2, '') AS type2,
    coalesce(altdamagenote, '') AS altdamagenote,
    savedamageactive,
    coalesce(ability, '') AS ability,
    dc,
    halfdamage,
    coalesce(savedamagenote, '') AS savedamagenote
FROM
    public.actiondamage