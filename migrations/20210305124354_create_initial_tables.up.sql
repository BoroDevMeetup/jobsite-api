CREATE TABLE users
(
    id SERIAL,
    slack_user_id character varying(50),
    slack_access_token character varying(255),
    first_name character varying(50),
    last_name character varying(50),
    email character varying(255) NOT NULL,
    email_verified_at date,
    created_at date,
    updated_at date,
    PRIMARY KEY (id),
    CONSTRAINT users_email_uniq UNIQUE (email)
);

CREATE TABLE jobs
(
    id SERIAL,
    user_id bigint,
    title character varying(255) NOT NULL,
    bio text,
    facebook character varying(255),
    twitter character varying(255),
    linked_in character varying(255),
    apply_link character varying(255),
    job_type character varying(255),
    xp character varying(255),
    remote boolean,
    job_description text,
    city character varying(255),
    state character varying(4),
    company_name character varying(255),
    tech_stack character varying(255),
    PRIMARY KEY (id),
    CONSTRAINT jobs_user_id_foreign FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);

CREATE TABLE companies
(
    id SERIAL,
    user_id bigint,
    name character varying(255),
    city character varying(255),
    state character varying(4),
    size integer,
    logo_url character varying(255),
    facebook character varying(255),
    twitter character varying(255),
    github character varying(255),
    linked_in character varying(255),
    description text,
    short_description character varying(255),
    PRIMARY KEY (id),
    CONSTRAINT companies_user_id_foreign FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);