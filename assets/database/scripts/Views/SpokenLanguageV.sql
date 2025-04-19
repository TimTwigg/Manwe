CREATE VIEW
    SpokenLanguageV AS
SELECT
    EntityID,
    Language,
    coalesce(Description, '') as Description
FROM
    SpokenLanguage