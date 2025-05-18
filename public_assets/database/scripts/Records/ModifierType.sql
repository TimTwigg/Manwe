INSERT
OR IGNORE INTO "ModifierType" (
    "ModifierType",
    "Description",
    "IsProficiencyRelevant"
)
VALUES
    ('DR', 'Damage Resistance', ''),
    ('DV', 'Damage Vulnerability', ''),
    ('DI', 'Damage Immunity', ''),
    ('CI', 'Condition Immunity', ''),
    ('SK', 'Skill', 'X'),
    ('ST', 'Saving Throw', 'X'),
    ('SE', 'Sense', ''),
    ('TR', 'Trait', '');