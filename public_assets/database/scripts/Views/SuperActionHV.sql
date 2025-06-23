CREATE VIEW
    SuperActionHV AS
SELECT
    StatBlockID,
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