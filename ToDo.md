# To Do

- Allow multiple entities with same name
    - DB allows, but parsing needs to handle this
- Pagination in api calls
- Rework so that stats are not hardcoded - allow custom stats and saving throws
- How to handle entities in encounter being non-statblocks (ie players or temp npcs)

# Database Design

## General

### Notes

- Entities created in encounters as custom entries or as players are stored in Entity table.

### Tables / Columns

- Add another type column to Entity table - RecordSource
    - "Statblock"
        - Standard or custom entities which have a stat block. Equivalent to standard / custom stat blocks in DDB Encounter Manager.
    - "Player"
        - Player entities. These are actual players correlating usually to real people, but at minimum to a character sheet (in whichever system).
    - "Custom"
        - Temporary entities created to track NPCs who are not worth creating an actual stat block for, or just until an actual stat block is created.
        - These should be cleaned up at semi-regular intervals.

## Security / Ownership Management

### Notes

- All queries should be filtered to objects within the user's area
    - read access to public domain
    - read/write access to their own domain

### Tables / Columns

- Add "Domain" column to all tables - foreign key to User table
    - "public" for standard objects
    - user's userid for custom objects
- Add "published" column to all tables
    - boolean
        - true means it is effectively public domain.
            - if the domain is not "public" then it is the creator of the object.
