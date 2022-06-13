create table log_rate_change
(
	trace_id text
		constraint log_rate_change_pk
			primary key,
	origin_currency text,
	destination_currency text,
	rate numeric,
	expired_at timestamptz,
	created_at timestamptz default now()
);
