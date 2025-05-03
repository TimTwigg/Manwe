CREATE VIEW
    ModifierV as
SELECT
    EntityID,
    Type,
    Name,
    coalesce(Value, 0) as Value,
    coalesce(Description, '') as Description,
    Domain,
    Published
FROM
    Modifiers