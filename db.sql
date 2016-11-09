-- Copyright 2016 David Lavieri.  All rights reserved.
-- Use of this source code is governed by a MIT License
-- License that can be found in the LICENSE file.

-- PostgreSQL
-- change schema "public" if required

CREATE TABLE public.user (
  id serial PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  email VARCHAR(90) UNIQUE NOT NULL,
  password VARCHAR(512) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE public.todo (
  id serial PRIMARY KEY,
  user_id int4 NOT NULL,
  title VARCHAR(256) NOT NULL,
  completed BOOL NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE,
  updated_at TIMESTAMP WITH TIME ZONE
);

ALTER TABLE public.todo
ADD CONSTRAINT todo_user_id FOREIGN KEY (user_id) REFERENCES public.user (id);
