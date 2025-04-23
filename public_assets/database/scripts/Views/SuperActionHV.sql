CREATE VIEW
    SuperActionHV AS
SELECT
    EntityID,
    Type,
    Description,
    Points
FROM
    SuperAction
WHERE
    Type in ('Legendary', 'Mythic')
    AND Name = 'X'