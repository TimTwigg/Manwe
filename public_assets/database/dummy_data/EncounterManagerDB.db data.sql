BEGIN TRANSACTION;

INSERT
OR IGNORE INTO "User" ("UserName")
VALUES
    ('Public'),
    ('Private');

INSERT
OR IGNORE INTO "RecordSource" ("RecordSource")
VALUES
    ('Statblock'),
    ('Player'),
    ('Custom');

INSERT
OR IGNORE INTO "Ability" ("Ability")
VALUES
    ('Strength'),
    ('Dexterity'),
    ('Constitution'),
    ('Intelligence'),
    ('Wisdom'),
    ('Charisma'),
    ('');

INSERT
OR IGNORE INTO "Condition" ("Condition")
VALUES
    ('Stunned');

INSERT
OR IGNORE INTO "DamageType" ("DamageType", "Description")
VALUES
    ('', ''),
    ('Piercing', ''),
    ('Slashing', '');

INSERT
OR IGNORE INTO "Language" ("Language", "Description")
VALUES
    ('Common', 'Everyone speaks common.');

INSERT
OR IGNORE INTO "Size" ("Size")
VALUES
    ('Tiny'),
    ('Small'),
    ('Medium'),
    ('Large'),
    ('Huge'),
    ('Gargantuan'),
    ('Medium or Small');

INSERT
OR IGNORE INTO "EntityType" ("EntityType")
VALUES
    ('Aberration'),
    ('Beast'),
    ('Celestial'),
    ('Construct'),
    ('Dragon'),
    ('Elemental'),
    ('Fey'),
    ('Fiend'),
    ('Giant'),
    ('Humanoid'),
    ('Monstrosity'),
    ('Ooze'),
    ('Plant'),
    ('Undead');

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

INSERT
OR IGNORE INTO "Entity" (
    "EntityID",
    "Name",
    "ChallengeRating",
    "ProficiencyBonus",
    "Source",
    "Size",
    "Type",
    "Alignment",
    "ArmorClass",
    "HitPoints1",
    "HitPoints2",
    "SWalk",
    "SFly",
    "SClimb",
    "SSwim",
    "SBurrow",
    "ReactionCount",
    "ArmorType",
    "RecordSource"
)
VALUES
    (
        1,
        'Winter Ghoul',
        1,
        2,
        'Homebrew',
        'Medium',
        'Undead',
        'Chaotic Evil',
        12,
        22,
        '5d8',
        30,
        0,
        0,
        0,
        0,
        1,
        'Natural Armor',
        'Statblock'
    );

INSERT
OR IGNORE INTO "EntityStats" ("EntityID", "Ability", "Value")
VALUES
    (1, 'Strength', 13),
    (1, 'Dexterity', 15),
    (1, 'Constitution', 10),
    (1, 'Intelligence', 7),
    (1, 'Wisdom', 10),
    (1, 'Charisma', 6);

INSERT
OR IGNORE INTO "Action" (
    "EntityID",
    "ActionID",
    "Name",
    "AttackType",
    "HitModifier",
    "Reach",
    "Targets",
    "Description"
)
VALUES
    (1, 1, 'Bite', 'Melee Weapon Attack', 2, 5, 1, ''),
    (
        1,
        2,
        'Claws',
        'Melee Weapon Attack',
        4,
        5,
        1,
        'If the target is a creature other than an undead, it must succeed on a DC 10 Constitution saving throw. On a failed save, a target begins to freeze and is restrained. The restrained target must repeat the saving throw at the end of each of its turns. On a success, the effect ends on the target. On a failure, the target is stunned. If the target fails this saving throw again, they are frozen and petrified. The target remains petrified for 24 hours, after which they thaw, or until freed by the greater restoration spell or other magic.'
    );

INSERT
OR IGNORE INTO "ActionDamage" (
    "EntityID",
    "ActionID",
    "DamageID",
    "Amount",
    "Type",
    "AltDmgActive",
    "Amount2",
    "Type2",
    "AltDmgNote",
    "SaveDmgActive",
    "Ability",
    "DC",
    "HalfDamage",
    "SaveDmgNote"
)
VALUES
    (
        1,
        1,
        1,
        '2d6 + 2',
        'Piercing',
        '',
        '',
        '',
        '',
        '',
        '',
        0,
        '',
        ''
    ),
    (
        1,
        2,
        1,
        '2d4 + 2',
        'Slashing',
        '',
        '',
        '',
        '',
        '',
        '',
        0,
        '',
        ''
    );

INSERT
OR IGNORE INTO "Encounter" (
    "EncounterID",
    "Name",
    "Description",
    "CreationDate",
    "AccessedDate",
    "Campaign",
    "Started",
    "Round",
    "Turn",
    "HasLair",
    "LairEntityName",
    "ActiveID"
)
VALUES
    (
        1,
        'Encounter 1',
        'Test Encounter',
        '20250419',
        '20250420',
        'Valez',
        '',
        0,
        0,
        '',
        '',
        ''
    );

INSERT
OR IGNORE INTO "EncounterEntities" (
    "EncounterID",
    "RowID",
    "EntityID",
    "Suffix",
    "Initiative",
    "MaxHitPoints",
    "TempHitPoints",
    "CurrentHitPoints",
    "ArmorClassBonus",
    "Notes",
    "IsHostile",
    "EncounterLocked",
    "ID"
)
VALUES
    (
        1,
        1,
        1,
        'A',
        10,
        22,
        0,
        10,
        0,
        'Wounded Ghoul',
        'X',
        '',
        ''
    ),
    (
        1,
        2,
        1,
        'B',
        4,
        13,
        3,
        13,
        -1,
        'Altered Ghoul',
        'X',
        'X',
        ''
    );

INSERT
OR IGNORE INTO "EncEntConditions" ("EncounterID", "RowID", "Condition", "Duration")
VALUES
    (1, 1, 'Stunned', 1);

INSERT
OR IGNORE INTO "Lair" ("EntityID", "Description", "Initiative")
VALUES
    (1, 'Ghoul lives in lair', 20);

INSERT
OR IGNORE INTO "Modifiers" (
    "EntityID",
    "Item",
    "Type",
    "Name",
    "Value",
    "Description"
)
VALUES
    (1, 1, 'DI', 'Cold', 0, ''),
    (1, 2, 'SE', 'Darkvision', 60, ''),
    (1, 3, 'SE', 'Passive Perception', 10, ''),
    (
        1,
        4,
        'TR',
        'Snow Camouflage',
        0,
        'The ghoul has advantage on Dexterity (Stealth) checks made to hide in snowy terrain.'
    );

INSERT
OR IGNORE INTO "Proficiencies" (
    "EntityID",
    "Item",
    "Type",
    "Name",
    "Level",
    "Override"
)
VALUES
    (1, 1, 'SK', 'Stealth', 1, 0),
    (1, 2, 'ST', 'Strength', 1, 0);

INSERT
OR IGNORE INTO "SimpleAction" (
    "EntityID",
    "ActionID",
    "Type",
    "Name",
    "Description"
)
VALUES
    (1, 1, 'Bonus', 'Dummy', 'Dummy Bonus Action'),
    (1, 2, 'Reaction', 'Dummy', 'Dummy Reaction');

INSERT
OR IGNORE INTO "SpokenLanguage" ("EntityID", "Language", "Description")
VALUES
    (1, 'Common', '');

INSERT
OR IGNORE INTO "SuperAction" (
    "EntityID",
    "ActionID",
    "Type",
    "Name",
    "Description",
    "Points",
    "IsRegional"
)
VALUES
    (
        1,
        1,
        'Legendary',
        'X',
        'Ghoul has 2 legendary actions.',
        2,
        ''
    ),
    (
        1,
        2,
        'Legendary',
        'Claws',
        'Ghoul attacks with claws',
        1,
        ''
    ),
    (1, 3, 'Legendary', 'Bite', 'Ghoul Bites', 2, ''),
    (
        1,
        1,
        'Mythic',
        'X',
        'If Ghoul is mythic, ghoul can use these as legendary actions',
        0,
        ''
    ),
    (
        1,
        2,
        'Mythic',
        'Claws',
        'Ghoul attacks twice with claws',
        1,
        ''
    ),
    (
        1,
        1,
        'Lair',
        'Despair',
        'Enemies despair',
        0,
        'X'
    ),
    (
        1,
        2,
        'Lair',
        'Summon',
        'Ghoul summons allies',
        0,
        ''
    ),
    (
        1,
        3,
        'Lair',
        'X',
        'Ghoul takes Lair Action',
        0,
        ''
    ),
    (1, 4, 'Lair', 'X', 'Ghoul''s home is bad', 0, 'X');

COMMIT;

PRAGMA INTEGRITY_CHECK;