create table if not exists employees(
    id serial primary key,
    name varchar not null ,
    email varchar not null unique,
    age int not null
);

INSERT INTO employees (name, email, age) VALUES
    ('Alice Johnson', 'alice.johnson@company.com', 28),
    ('Bob Smith', 'bob.smith@company.com', 32),
    ('Carol Davis', 'carol.davis@company.com', 45),
    ('David Wilson', 'david.wilson@company.com', 29),
    ('Emma Brown', 'emma.brown@company.com', 35),
    ('Frank Miller', 'frank.miller@company.com', 41),
    ('Grace Taylor', 'grace.taylor@company.com', 27),
    ('Henry Clark', 'henry.clark@company.com', 38),
    ('Ivy Martinez', 'ivy.martinez@company.com', 31),
    ('Jack Anderson', 'jack.anderson@company.com', 26);