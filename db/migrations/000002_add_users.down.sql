-- Drop the unique constraint on "accounts" table
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "unique_owner_currency";

-- Drop the foreign key constraint on "accounts" table
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drop the "users" table
DROP TABLE IF EXISTS "users";