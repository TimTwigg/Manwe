CREATE VIEW
    SuperActionV AS
SELECT
    EntityID,
    Type,
    Name,
    Description,
    Points
FROM
    SuperAction
WHERE
    Type in ('Legendary', 'Mythic')
    and Name <> 'X'