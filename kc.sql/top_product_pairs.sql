-- Counting the most popular product pairs
-- TODO: find a way to use less joins

WITH ordered_goods AS (
  SELECT
    DISTINCT order_id,
    unnest(product_ids) AS product_id
  FROM orders
  WHERE order_id NOT IN (SELECT order_id FROM user_actions WHERE action = 'cancel_order')
), 
order_goods_rows AS (
  SELECT *, row_number() OVER (PARTITION BY order_id) AS row_id
  FROM (SELECT * FROM ordered_goods) AS tmp 
  ORDER BY order_id
),
goods_pairs AS (
  SELECT
    t1.product_id AS product1_id,
    t2.product_id AS product2_id
  FROM
    order_goods_rows t1
    INNER JOIN
    order_goods_rows t2 USING (order_id)
  WHERE t1.row_id < t2.row_id
)

SELECT
  paired_goods AS pair,
  COUNT(paired_goods) AS count_pair
FROM (
    SELECT array_sort(array[p1.name, p2.name]) AS paired_goods 
    FROM
      goods_pairs 
      INNER JOIN products p1 ON p1.product_id = goods_pairs.product1_id 
      INNER JOIN products p2 ON p2.product_id = goods_pairs.product2_id
  ) tmp
GROUP BY paired_goods 
ORDER BY count_pair DESC, pair
