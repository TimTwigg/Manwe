CREATE VIEW
    LairActionV AS
SELECT
    StatBlockID,
    Name,
    Description,
    coalesce(IsRegional, '') as IsRegional,
    Domain,
    Published
FROM
    SuperAction
WHERE
    Type = 'Lair'