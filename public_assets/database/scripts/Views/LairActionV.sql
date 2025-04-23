CREATE VIEW
    LairActionV AS
SELECT
    EntityID,
    Name,
    Description,
    coalesce(IsRegional, '') as IsRegional
FROM
    SuperAction
WHERE
    Type = 'Lair'