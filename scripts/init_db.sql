CREATE TABLE "todo" (
    id serial primary key,
    title varchar NOT NULL,
    description text NOT NULL,
    done bool
);
