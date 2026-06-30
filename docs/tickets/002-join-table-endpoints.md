# 002 - Join Table Endpoints

**Status:** Planned

## Scope

Add endpoints to assign personnel and assets to missions using the existing `MISSION_PERSONNEL` and `MISSION_ASSET` join tables.

## Endpoints

- `POST /missions/{id}/personnel` - assign a personnel record to a mission with a role
- `DELETE /missions/{id}/personnel/{personnel_id}` - remove personnel from a mission
- `POST /missions/{id}/assets` - assign an asset to a mission with a role
- `DELETE /missions/{id}/assets/{asset_id}` - remove an asset from a mission

## Acceptance Criteria

- Request body for POST includes `personnel_id`/`asset_id` and `role`
- Returns 404 if the mission or the personnel/asset does not exist
- Returns 409 if the assignment already exists
- Successful POST returns the created assignment
