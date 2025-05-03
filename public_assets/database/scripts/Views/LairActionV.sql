CREATE VIEW
    LairActionV AS
SELECT
    EntityID,
    Name,
    Description,
    coalesce(IsRegional, '') as IsRegional,
    Domain,
    Published
FROM
    SuperAction
WHERE
    Type = 'Lair'