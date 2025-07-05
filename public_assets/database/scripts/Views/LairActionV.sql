CREATE VIEW
    LairActionV AS
SELECT
    StatBlockID,
    Name,
    Description,
    coalesce(IsRegional, '') as IsRegional
FROM
    SuperAction
WHERE
    Type = 'Lair'