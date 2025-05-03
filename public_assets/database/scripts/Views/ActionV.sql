CREATE VIEW
    ActionV AS
SELECT
    EntityID,
    ActionID,
    Name,
    coalesce(AttackType, '') as AttackType,
    coalesce(HitModifier, 0) as HitModifier,
    coalesce(Reach, 0) as Reach,
    coalesce(Targets, 0) as Targets,
    coalesce(Description, '') as Description,
    Domain,
    Published
FROM
    Action