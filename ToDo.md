# To Do

- Allow multiple entities with same name
    - DB allows, but parsing needs to handle this
- Pagination in api calls
- Rework so that stats are not hardcoded - allow custom stats and saving throws
    - Think this is done - review
- How to handle entities in encounter being non-statblocks (ie players or temp npcs)
- Create campaigns as actual objects so that players can be grouped in a campaign and easily all pulled into a new encounter

# Database Design

## General

### Notes

- Entities created in encounters as custom entries or as players are stored in Entity table.
    - RecordSource column defines what kind of entity it is:
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
