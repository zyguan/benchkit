create table if not exists `bench_result_tag` (
	`id`  varchar(40),
	`tag` varchar(255),
	primary key (`id`, `tag`) /*T![clustered_index] CLUSTERED */
);
