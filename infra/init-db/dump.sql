CREATE TABLE "public"."users" (
      id uuid PRIMARY KEY,
      email varchar(255) NOT NULL UNIQUE,
      is_active boolean NOT NULL DEFAULT true
);
