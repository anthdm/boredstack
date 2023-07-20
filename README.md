# The bored stack

Programming stack for the no bullshit builder.

### Database utils

You can run the following database util commands using make:

* **db-init** -- Initialise bun migrations table
* **db-migrate** -- Run migrations
* **db-rollback** -- Rollback migrations
* **db-lock** -- Locks the database to prevent future migrations
* **db-unlock** -- Unlocks the database to enable future migrations
* **db-create-sql name=migration_name** -- Creates new sql migration with given name
* **db-create-go name=migration_name** -- Creates new go migration with given name
* **db-status** -- Prints database migrations status
* **db-mark-applied** -- Marks all migrations as ran without actually running them