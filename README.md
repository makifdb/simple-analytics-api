# Simple Analytics API – Challenges and Key Learnings

## Project Overview

**Objective:**
Develop a simple analytics API in Go to track and query user events using GraphQL. The project utilizes:
- **GraphQL** for the API.
- **ClickHouse** as the database for fast analytical queries.
- **Entgo** as the ORM to define the data structure in Go.

**Key Components:**
- **GraphQL API:**
  - Mutation to record new events.
  - Queries to retrieve events for a specific user.
  - Query for aggregating event counts by event type within a specified time range.
- **ClickHouse:**
  - Chosen for high performance in analytical workloads.
- **Entgo:**
  - Used to define the schema (event structure) as Go structs.
  - Due to limitations with ClickHouse, automatic migrations were not possible, so manual SQL is used instead.

---

## Main Challenges & Key Learnings

### 1. Entgo and ClickHouse Compatibility

- **Challenge:**
  Entgo is optimized for traditional relational databases (PostgreSQL, MySQL, SQLite) and doesn’t fully support ClickHouse.
  - **Issue:**
    Using standard Entgo methods (e.g. `Save()`) led to unexpected behaviors such as duplicate events or errors like a missing default `id` column.

- **Solution:**
  - **Raw SQL for Inserts and Queries:**
    To gain full control, raw SQL commands were used to perform inserts and queries. This bypassed the problematic ORM functions.
  - **Custom Dialect Handling:**
    The project created a `*sql.DB` connection with the ClickHouse driver and wrapped it using `entsql.OpenDB`, passing a custom dialect via `dialect.Other("clickhouse")`. This prevented the “unsupported driver” error.

---

### 2. Manual Schema Management

- **Challenge:**
  Automatic schema creation with Entgo failed when used with ClickHouse. Entgo’s migration tool expects columns (like a default `id`) that were not present in our manually defined table.

- **Solution:**
  - **Manual SQL Migration:**
    A raw SQL migration script was written to create the `events` table with precise ClickHouse syntax (using the MergeTree engine, proper `ORDER BY`, and defining primary keys).
  - **Schema Alignment:**
    The schema was adjusted to include both an `id` column (expected by Entgo) and an `event_id` column (to reflect the domain model). Both columns store the same UUID value.

---

### 3. Data Handling and Conversion

- **Challenge:**
  Ensuring consistency between how data is stored in ClickHouse and how it's used in the API:
  - **JSON Metadata as Text:**
    Metadata needed to be stored as a JSON string, yet processed as structured data in Go.
  - **Duplicate Event Issues:**
    The ORM methods sometimes introduced duplicate events due to schema mismatches.

- **Solution:**
  - **Manual Conversion:**
    Resolvers were refactored to manually validate and parse JSON metadata (using `json.Unmarshal`). The raw JSON string is then used for insertion while the parsed data is kept for API responses.
  - **Constructing Return Objects Manually:**
    After a raw SQL insert, the `ent.Event` object is manually constructed to match the expected GraphQL response and the Ent schema.

---

### 4. GraphQL API and Data Aggregation

- **Challenge:**
  Mapping between the GraphQL API and the underlying data:
  - **Field Expectations:**
    Entgo-generated queries expected an `id` field, which was missing if only `event_id` was defined.

- **Solution:**
  - **Schema Mapping Adjustments:**
    The migration script was updated to include both `id` and `event_id` (with identical UUIDs) to satisfy both raw SQL insertions and ORM queries.
  - **Refactoring Resolvers:**
    Resolvers were adapted to use raw SQL commands where necessary, ensuring that the data passed between the API and the database was consistent.

---

## Final Takeaways

- **Flexibility vs. Abstraction:**
  ORMs like Entgo simplify many database operations, but for ClickHouse—where the database behaves very differently—raw SQL commands provided the necessary control.

- **Manual Intervention:**
  Automatic migrations do not always work with non-standard databases, necessitating manual migration scripts to ensure data integrity.

- **Understanding Tool Limitations:**
  Integrating GraphQL, ClickHouse, and Entgo demonstrated that every layer has specific requirements and limitations. It’s vital to understand these to design robust systems.

- **Error Handling and Consistency:**
  Careful error handling, consistent data conversion, and schema alignment were crucial to bridge the gap between expected ORM behavior and ClickHouse’s requirements.

**Conclusion:**
While Entgo proved useful for defining the Go data structures, practical use with ClickHouse revealed several challenges. A hybrid approach—using Entgo for type definitions and raw SQL for data operations—offered the stability and performance required for the analytics API. This project highlights the need for flexibility and in-depth understanding of both the ORM and the database when integrating diverse technologies.
