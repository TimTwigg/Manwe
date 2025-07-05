CREATE VIEW
    SuperActionHV AS
SELECT
    StatBlockID,
    Type,
    Description,
    Points
FROM
    SuperAction
WHERE
    Type in ('Legendary', 'Mythic')
    AND Name = 'X'