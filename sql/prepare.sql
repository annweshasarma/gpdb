LOAD 'pg_hint_plan';
SET pg_hint_plan.enable TO on;
SET pg_hint_plan.debug_print TO on;
SET client_min_messages TO LOG;

EXPLAIN (COSTS false) SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id;
EXPLAIN (COSTS false) SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > 10;

-- 9.1:PREPAREで指定したヒント句で実行計画が固定される
-- 9.2:PREPAREでヒント句を指定しても、実行計画は制御できない
/*+ NestLoop(t1 t2) */
PREPARE p1 AS SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id;
EXPLAIN (COSTS false) EXECUTE p1;

-- 9.1:PREPAREで指定したヒント句で実行計画が固定される
-- 9.2:パラメータがない場合は、1回目のEXPLAINで実行計画が決定する
/*+ NestLoop(t1 t2) */
PREPARE p2 AS SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id;
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p2;
EXPLAIN (COSTS false) EXECUTE p2;
EXPLAIN (COSTS false) EXECUTE p2;

-- 9.1:PREPAREで指定したヒント句で実行計画が固定される
-- 9.2:5回目のEXPLAINまでヒント句を指定しても、6回目以降は本来の実行計画に戻る
/*+ NestLoop(t1 t2) */
PREPARE p3 AS SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > $1;
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p3 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p3 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p3 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p3 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p3 (10);
EXPLAIN (COSTS false) EXECUTE p3 (10);
EXPLAIN (COSTS false) EXECUTE p3 (10);
EXPLAIN (COSTS false) EXECUTE p3 (10);

-- 9.1:PREPAREで指定したヒント句で実行計画が固定される
-- 9.2:6回目のEXPLAINまでヒント句を指定すると、7回目以降も実行計画が固定される
/*+ NestLoop(t1 t2) */
PREPARE p4 AS SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > $1;
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p4 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p4 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p4 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p4 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p4 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p4 (10);
EXPLAIN (COSTS false) EXECUTE p4 (10);
EXPLAIN (COSTS false) EXECUTE p4 (10);

-- 9.1:PREPAREで指定したヒント句で実行計画が固定される
-- 9.2:6回目のEXPLAINでヒント句を指定すると、7回目以降も実行計画を制御できる
/*+ NestLoop(t1 t2) */
PREPARE p5 AS SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > $1;
EXPLAIN (COSTS false) EXECUTE p5 (10);
EXPLAIN (COSTS false) EXECUTE p5 (10);
EXPLAIN (COSTS false) EXECUTE p5 (10);
EXPLAIN (COSTS false) EXECUTE p5 (10);
EXPLAIN (COSTS false) EXECUTE p5 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p5 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p5 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p5 (10);

-- 9.1:PREPAREで指定したヒント句で実行計画が固定される
-- 9.2:7回目以降のEXPLAINでヒント句を指定しても、以降も実行計画は制御できない
/*+ NestLoop(t1 t2) */
PREPARE p6 AS SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > $1;
EXPLAIN (COSTS false) EXECUTE p6 (10);
EXPLAIN (COSTS false) EXECUTE p6 (10);
EXPLAIN (COSTS false) EXECUTE p6 (10);
EXPLAIN (COSTS false) EXECUTE p6 (10);
EXPLAIN (COSTS false) EXECUTE p6 (10);
EXPLAIN (COSTS false) EXECUTE p6 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p6 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p6 (10);

-- 9.1:実行計画が固定されたあと、ANALYZEをすると1回目のEXECUTEで実行計画が固定される
-- 9.2:実行計画が固定されたあと、ANALYZEをすると1回目のEXECUTEで実行計画が固定される
/*+ NestLoop(t1 t2) */
PREPARE p7 AS SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > $1;
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);
EXPLAIN (COSTS false) EXECUTE p7 (10);
EXPLAIN (COSTS false) EXECUTE p7 (10);

TRUNCATE t1;
ANALYZE t1;
EXPLAIN (COSTS false) SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > 10;
EXPLAIN (COSTS false) EXECUTE p7 (10);
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);

INSERT INTO t1 SELECT i, i % 100 FROM (SELECT generate_series(1, 10000) i) t;
ANALYZE t1;
EXPLAIN (COSTS false) SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > 10;
/*+ HashJoin(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);
/*+ NestLoop(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p7 (10);

-- error case
/*+ NestLoop(t1 t2) */
EXPLAIN (COSTS false) EXECUTE p8 (10);
/*+ NestLoop(t1 t2) */
EXPLAIN (COSTS false) SELECT count(*) FROM t1, t2 WHERE t1.id = t2.id AND t1.id > 10;
