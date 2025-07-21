# To Do

- Allow multiple entities with same name
    - DB allows, but parsing needs to handle this
        - New pgx QueryRow where only one is expected will select only top row if multiple are present.
            - Should this be explicitly handled?
- Pagination in api calls
- How to handle entities in encounter being non-statblocks (ie players or temp npcs)
    - Difference type
    - Should be handled now, test
- Finish reworking db CRUD operations to use pgx instead of pq
    - EncounterReader and ConditionReader done

# Database Design

## General

### Notes

## Security / Ownership Management

### Notes

- All queries should be filtered to objects within the user's area
    - read access to public domain
    - read/write access to their own domain
