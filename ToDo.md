# To Do

- Allow multiple entities with same name
    - DB allows, but parsing needs to handle this
- Pagination in api calls
- Rework so that stats are not hardcoded - allow custom stats and saving throws
- Save Encounter should save and then re-read/return saved encounter.
    - For new encounters the incoming one will have id=0, return the real encounter with a valid id
