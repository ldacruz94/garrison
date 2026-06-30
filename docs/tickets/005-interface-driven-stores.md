# 005 - Interface-Driven Store Layer

**Status:** Planned

## Scope

Define interfaces for each store so handlers depend on the interface, not the concrete type. This is the prerequisite for writing tests with mock stores.

## Acceptance Criteria

- Each store (`MissionStore`, `PersonnelStore`, `AssetStore`, `AuditLogStore`) has a corresponding interface
- Handlers accept the interface, not the concrete struct
- Concrete store structs satisfy their interfaces (no behavior change)
