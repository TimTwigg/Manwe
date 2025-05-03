CREATE VIEW
    ActionDamageV as
SELECT
    EntityID,
    ActionID,
    coalesce(Amount, '') as Amount,
    coalesce(Type, '') as Type,
    coalesce(AltDmgActive, '') as AltDmgActive,
    coalesce(Amount2, '') as Amount2,
    coalesce(Type2, '') as Type2,
    coalesce(AltDmgNote, '') as AltDmgNote,
    coalesce(SaveDmgActive, '') as SaveDmgActive,
    coalesce(Ability, '') as Ability,
    coalesce(DC, 0) as DC,
    coalesce(HalfDamage, '') as HalfDamage,
    coalesce(SaveDmgNote, '') as SaveDmgNote,
    Domain,
    Published
FROM
    ActionDamage