# Simple Analytics API

## Objective
Create a simple analytics API that tracks and queries user events using Go, with the following requirements:
- Use GraphQL to define and expose the API.
- Use ClickHouse as the database to store and query the events.
- Use Entgo for ORM and schema management.

## Task Details

### 1. Setup and Environment
- Set up a Go project with appropriate module dependencies.
- Use Docker to provide a local instance of ClickHouse for testing and development.

### 2. Schema Design
- Define an event schema using Entgo. An event should have the following fields:
  - `EventID` (UUID)
  - `UserID` (UUID)
  - `EventType` (string)
  - `Timestamp` (datetime)
  - `Metadata` (JSONB or similar for additional event data)

### 3. API Endpoints
- Implement GraphQL mutations to:
  - Record a new event.
- Implement GraphQL queries to:
  - Retrieve events for a specific user.
  - Aggregate and return event counts by `EventType` within a specified time range.

### 4. Database Integration
- Configure Entgo to generate and manage the necessary schema in ClickHouse.
- Implement repository functions to handle CRUD operations for the events.
