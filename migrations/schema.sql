--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4 (Debian 13.4-1.pgdg100+1)
-- Dumped by pg_dump version 13.4 (Ubuntu 13.4-1.pgdg20.04+1)

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
-- Name: container_inventories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.container_inventories (
    id integer NOT NULL,
    container_id integer,
    item_id integer,
    rate real DEFAULT 1.0
);


ALTER TABLE public.container_inventories OWNER TO postgres;

--
-- Name: containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.containers (
    id integer NOT NULL,
    name character varying(100),
    description text,
    is_floor boolean
);


ALTER TABLE public.containers OWNER TO postgres;

--
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer NOT NULL,
    name character varying(100),
    description text,
    value integer,
    category smallint
);


ALTER TABLE public.items OWNER TO postgres;

--
-- Name: player_current_inventories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.player_current_inventories (
    id integer NOT NULL,
    player_id integer NOT NULL,
    item_id integer NOT NULL,
    is_bound boolean DEFAULT false,
    is_equipped boolean DEFAULT false
);


ALTER TABLE public.player_current_inventories OWNER TO postgres;

--
-- Name: player_current_inventories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.player_current_inventories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.player_current_inventories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: players; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.players (
    id integer NOT NULL,
    name character varying(20) NOT NULL,
    password character(64) NOT NULL,
    race smallint DEFAULT 0,
    pronouns smallint DEFAULT 0,
    saved_room_id integer,
    current_room_id integer
);


ALTER TABLE public.players OWNER TO postgres;

--
-- Name: players_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.players ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.players_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: room_connections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.room_connections (
    id integer NOT NULL,
    from_id integer,
    to_id integer,
    direction smallint,
    locked_by_id integer,
    locked_by_flag character varying(100)
);


ALTER TABLE public.room_connections OWNER TO postgres;

--
-- Name: room_connections_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.room_connections ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.room_connections_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: room_containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.room_containers (
    id integer NOT NULL,
    container_id integer NOT NULL,
    room_id integer NOT NULL
);


ALTER TABLE public.room_containers OWNER TO postgres;

--
-- Name: room_containers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.room_containers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.room_containers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: room_current_inventories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.room_current_inventories (
    id integer NOT NULL,
    room_container_id integer NOT NULL,
    item_id integer NOT NULL,
    visible_to_id integer NOT NULL
);


ALTER TABLE public.room_current_inventories OWNER TO postgres;

--
-- Name: room_current_inventories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.room_current_inventories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.room_current_inventories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: rooms; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rooms (
    id integer NOT NULL,
    description text,
    zone_id integer,
    is_safe boolean DEFAULT false NOT NULL
);


ALTER TABLE public.rooms OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: zones; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.zones (
    id integer NOT NULL,
    name character varying(100),
    description text
);


ALTER TABLE public.zones OWNER TO postgres;

--
-- Name: container_inventories container_inventories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.container_inventories
    ADD CONSTRAINT container_inventories_pkey PRIMARY KEY (id);


--
-- Name: containers containers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT containers_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: player_current_inventories player_current_inventories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_current_inventories
    ADD CONSTRAINT player_current_inventories_pkey PRIMARY KEY (id);


--
-- Name: players players_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT players_pkey PRIMARY KEY (id);


--
-- Name: room_connections room_connections_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_connections
    ADD CONSTRAINT room_connections_pkey PRIMARY KEY (id);


--
-- Name: room_containers room_containers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_containers
    ADD CONSTRAINT room_containers_pkey PRIMARY KEY (id);


--
-- Name: room_current_inventories room_current_inventories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_current_inventories
    ADD CONSTRAINT room_current_inventories_pkey PRIMARY KEY (id);


--
-- Name: rooms rooms_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT rooms_pkey PRIMARY KEY (id);


--
-- Name: zones zones_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zones
    ADD CONSTRAINT zones_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: container_inventories fk_container; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.container_inventories
    ADD CONSTRAINT fk_container FOREIGN KEY (container_id) REFERENCES public.containers(id);


--
-- Name: room_containers fk_container; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_containers
    ADD CONSTRAINT fk_container FOREIGN KEY (container_id) REFERENCES public.containers(id);


--
-- Name: players fk_current_room; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT fk_current_room FOREIGN KEY (current_room_id) REFERENCES public.rooms(id);


--
-- Name: room_connections fk_from; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_connections
    ADD CONSTRAINT fk_from FOREIGN KEY (from_id) REFERENCES public.rooms(id);


--
-- Name: container_inventories fk_item; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.container_inventories
    ADD CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- Name: player_current_inventories fk_item; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_current_inventories
    ADD CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- Name: room_current_inventories fk_item; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_current_inventories
    ADD CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- Name: room_connections fk_locked_by; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_connections
    ADD CONSTRAINT fk_locked_by FOREIGN KEY (locked_by_id) REFERENCES public.items(id);


--
-- Name: player_current_inventories fk_player; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_current_inventories
    ADD CONSTRAINT fk_player FOREIGN KEY (player_id) REFERENCES public.players(id);


--
-- Name: room_containers fk_room; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_containers
    ADD CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES public.rooms(id);


--
-- Name: room_current_inventories fk_room_container; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_current_inventories
    ADD CONSTRAINT fk_room_container FOREIGN KEY (room_container_id) REFERENCES public.room_containers(id);


--
-- Name: players fk_saved_room; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.players
    ADD CONSTRAINT fk_saved_room FOREIGN KEY (saved_room_id) REFERENCES public.rooms(id);


--
-- Name: room_connections fk_to; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_connections
    ADD CONSTRAINT fk_to FOREIGN KEY (to_id) REFERENCES public.rooms(id);


--
-- Name: room_current_inventories fk_visible_to; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.room_current_inventories
    ADD CONSTRAINT fk_visible_to FOREIGN KEY (visible_to_id) REFERENCES public.players(id);


--
-- Name: rooms fk_zone; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rooms
    ADD CONSTRAINT fk_zone FOREIGN KEY (zone_id) REFERENCES public.zones(id);


--
-- PostgreSQL database dump complete
--

