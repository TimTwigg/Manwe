CREATE VIEW
    SuperActionV AS
SELECT
    EntityID,
    Type,
    Name,
    Description,
    Points,
    Domain,
    Published
FROM
    SuperAction
WHERE
    Type in ('Legendary', 'Mythic')
    and Name <> 'X'