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

CREATE TABLE "public"."answers" (
    id uuid PRIMARY KEY,
    question_id uuid NOT NULL,
    text varchar(100) NOT NULL,
    is_correct bool DEFAULT false
);

CREATE TABLE "public"."questions" (
    id uuid PRIMARY KEY,
    text varchar(100) NOT NULL
);

ALTER TABLE "public"."answers"
    ADD CONSTRAINT fk_answers_questions
        FOREIGN KEY (question_id) REFERENCES questions (id);

INSERT INTO "public"."questions" (id, text)
    VALUES ('47ce6c5a-b97f-4b66-a3b9-54eb3c018746', 'What is structure in GoLang?'),
        ('b054c21e-6d28-4522-b614-709544fcde38', 'You can pass the struct to/from a function as'),
        ('15c1d4aa-954c-4b4b-8855-b25efe833a34', 'Is it possible to create a pointer to a struct?');

INSERT INTO "public"."answers" (id, question_id, text, is_correct)
    VALUES ('b0bb5845-cff0-459c-acc7-1ab07db2c567', '47ce6c5a-b97f-4b66-a3b9-54eb3c018746', 'Interface', false),
        ('9706f4b0-7d60-44df-85b6-12705a55a7ec', '47ce6c5a-b97f-4b66-a3b9-54eb3c018746', 'Typed collection of fields', true),
        ('03d7d018-d541-4bda-aafe-a7c2fd32f962', '47ce6c5a-b97f-4b66-a3b9-54eb3c018746', 'Pointer', false),

        ('600eddca-39c0-4254-8559-0934522e0078', 'b054c21e-6d28-4522-b614-709544fcde38', 'Value', true),
        ('f804c343-e575-48d5-bfbc-72cf52727541', 'b054c21e-6d28-4522-b614-709544fcde38', 'Link', false),

        ('1239c5b9-b41a-419e-bf45-f567ef97af30', '15c1d4aa-954c-4b4b-8855-b25efe833a34', 'Yes', true),
        ('8dc31e52-7c66-4f3a-8374-74b8859373fb', '15c1d4aa-954c-4b4b-8855-b25efe833a34', 'No', false);
