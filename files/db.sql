CREATE TABLE public.users (
    id serial NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW(),
	deleted_at timestamptz NULL,
    username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    CONSTRAINT predictions_pk PRIMARY KEY (id)
);

CREATE TABLE public.predictions (
	id serial NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	updated_at timestamptz NOT NULL DEFAULT NOW(),
	deleted_at timestamptz NULL,
	user_id int NOT NULL,
	image_path varchar(255) NULL,
	prediction int NOT NULL,
	CONSTRAINT predictions_pk PRIMARY KEY (id)
);