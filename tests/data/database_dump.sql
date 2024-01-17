--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1 (Debian 16.1-1.pgdg120+1)
-- Dumped by pg_dump version 16.1 (Debian 16.1-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: posts; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.posts (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text,
    content text
);


ALTER TABLE public.posts OWNER TO admin;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.posts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.posts_id_seq OWNER TO admin;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.posts_id_seq OWNED BY public.posts.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    password bytea NOT NULL
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO admin;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.posts ALTER COLUMN id SET DEFAULT nextval('public.posts_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.posts (id, created_at, updated_at, deleted_at, title, content) FROM stdin;
1	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 1	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod
2	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 2	Lorem ipsum dolor sit amet consetetur sadipscing
3	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 3	Lorem ipsum dolor sit amet
4	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 4	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor
5	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 5	Lorem ipsum dolor sit amet consetetur sadipscing elitr
6	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 6	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna aliquyam
7	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 7	Lorem ipsum dolor sit amet
8	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 8	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
9	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 9	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor
10	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 10	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy
11	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 11	Lorem ipsum dolor sit amet consetetur
12	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 12	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut
13	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 13	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
14	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 14	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
15	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 15	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
16	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 16	Lorem ipsum dolor sit amet consetetur sadipscing elitr
17	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 17	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
18	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 18	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna
19	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 19	Lorem ipsum dolor sit amet consetetur sadipscing elitr
20	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 20	Lorem ipsum dolor sit amet consetetur sadipscing
21	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 21	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
22	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 22	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor
23	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 23	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
24	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 24	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna aliquyam
25	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 25	Lorem ipsum dolor sit amet consetetur sadipscing elitr
26	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 26	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
27	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 27	Lorem ipsum dolor sit amet consetetur sadipscing
28	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 28	Lorem ipsum dolor sit amet consetetur sadipscing
29	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 29	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut
30	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 30	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
31	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 31	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut
32	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 32	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
33	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 33	Lorem ipsum dolor sit amet
34	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 34	Lorem ipsum dolor sit amet consetetur
35	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 35	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor
36	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 36	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
37	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 37	Lorem ipsum dolor sit amet consetetur sadipscing elitr
38	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 38	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut
39	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 39	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut
40	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 40	Lorem ipsum dolor sit amet consetetur sadipscing
41	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 41	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed
42	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 42	Lorem ipsum dolor sit amet
43	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 43	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy
44	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 44	Lorem ipsum dolor sit amet consetetur sadipscing elitr
45	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 45	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
46	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 46	Lorem ipsum dolor sit amet consetetur sadipscing elitr
47	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 47	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
48	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 48	Lorem ipsum dolor sit amet consetetur
49	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 49	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna aliquyam
50	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 50	Lorem ipsum dolor sit amet consetetur
51	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 51	Lorem ipsum dolor sit amet consetetur sadipscing
52	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 52	Lorem ipsum dolor sit amet
53	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 53	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
54	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 54	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
55	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 55	Lorem ipsum dolor sit amet
56	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 56	Lorem ipsum dolor sit amet consetetur sadipscing elitr
57	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 57	Lorem ipsum dolor sit amet consetetur
58	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 58	Lorem ipsum dolor sit amet consetetur sadipscing
59	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 59	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed
60	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 60	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
61	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 61	Lorem ipsum dolor sit amet consetetur sadipscing elitr
62	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 62	Lorem ipsum dolor sit amet consetetur sadipscing elitr
63	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 63	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore
64	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 64	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut
65	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 65	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
66	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 66	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
67	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 67	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
68	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 68	Lorem ipsum dolor sit amet consetetur sadipscing
69	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 69	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod
70	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 70	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
71	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 71	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy
72	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 72	Lorem ipsum dolor sit amet consetetur
73	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 73	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed
74	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 74	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor
75	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 75	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna
76	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 76	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna aliquyam
77	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 77	Lorem ipsum dolor sit amet
78	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 78	Lorem ipsum dolor sit amet consetetur sadipscing
79	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 79	Lorem ipsum dolor sit amet
80	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 80	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
81	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 81	Lorem ipsum dolor sit amet consetetur sadipscing
82	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 82	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna aliquyam
83	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 83	Lorem ipsum dolor sit amet consetetur
84	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 84	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
85	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 85	Lorem ipsum dolor sit amet consetetur sadipscing elitr
86	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 86	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed
87	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 87	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam
88	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 88	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod
89	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 89	Lorem ipsum dolor sit amet consetetur
90	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 90	Lorem ipsum dolor sit amet consetetur
91	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 91	Lorem ipsum dolor sit amet consetetur
92	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 92	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
93	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 93	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua
94	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 94	Lorem ipsum dolor sit amet consetetur sadipscing elitr
95	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 95	Lorem ipsum dolor sit amet consetetur
96	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 96	Lorem ipsum dolor sit amet consetetur sadipscing elitr
97	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 97	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod tempor ut labore voluptua magna aliquyam
98	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 98	Lorem ipsum dolor sit amet
99	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 99	Lorem ipsum dolor sit amet consetetur sadipscing elitr sed diam nonumy eirmod
100	2024-01-16 16:55:57.10057+00	2024-01-16 16:55:57.10057+00	\N	Post 100	Lorem ipsum dolor sit amet consetetur
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.users (id, email, password) FROM stdin;
\.


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.posts_id_seq', 100, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

