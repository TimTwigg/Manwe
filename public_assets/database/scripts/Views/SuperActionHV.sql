CREATE VIEW
    SuperActionHV AS
SELECT
    EntityID,
    Type,
    Description,
    Points,
    Domain,
    Published
FROM
    SuperAction
WHERE
    Type in ('Legendary', 'Mythic')
    AND Name = 'X'