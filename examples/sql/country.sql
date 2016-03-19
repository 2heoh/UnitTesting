CREATE TABLE country
(
  id serial NOT NULL,
  parent_id bigint,
  point_osm_id bigint,
  polygon_osm_id bigint,
  shp_gid text,
  src text,
  group_type text,
  type text,
  name text,
  alt_names text[],
  lat text,
  lon text,
  fias_id text,
  is_capital integer,
  path text,
  info text,
  place_type text,
  fias_type_ru text,
  official_status_ru text,
  msg text,
  updated integer,
  action_log_id bigint,
  parents_ids bigint[],
  CONSTRAINT address_pkey PRIMARY KEY (id)
) WITH (OIDS=FALSE);
ALTER TABLE address OWNER TO address;