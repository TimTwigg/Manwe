CREATE VIEW
    SpokenLanguageV AS
SELECT
    EntityID,
    Language,
    coalesce(Description, '') as Description,
    Domain,
    Published
FROM
    SpokenLanguage