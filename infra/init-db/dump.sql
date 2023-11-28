CREATE TABLE "public"."users" (
      id uuid PRIMARY KEY,
      email varchar(255) NOT NULL UNIQUE,
      is_active boolean NOT NULL DEFAULT true
);

CREATE TABLE "public"."articles" (
      id uuid PRIMARY KEY,
    title varchar(255) NOT NULL,
      body text NULL,
      created_at timestamp without time zone NOT NULL,
      updated_at timestamp without time zone NOT NULL
);
