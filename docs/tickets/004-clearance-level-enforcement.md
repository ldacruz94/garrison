# 004 - Clearance Level Enforcement

**Status:** Planned

## Scope

Enforce clearance levels so that personnel can only be assigned to missions that match their clearance. Currently `clearance_level` on personnel is just a data field with no logic attached.

## Clearance Levels (low to high)

`UNCLASSIFIED` < `CONFIDENTIAL` < `SECRET` < `TOP_SECRET`

## Acceptance Criteria

- Assigning personnel to a mission checks their clearance level against the mission type or a mission clearance field
- Returns 403 if the personnel's clearance is insufficient
- Clearance check lives in the service/store layer, not the handler
