CREATE VIEW
    SuperActionV AS
SELECT
    StatBlockID,
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