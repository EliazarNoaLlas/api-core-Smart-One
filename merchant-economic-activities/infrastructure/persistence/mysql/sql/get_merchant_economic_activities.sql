SELECT merchant_economic.id                   AS merchant_economic_id,
       merchant_economic.sequence             AS merchant_economic_sequence,
       merchant_economic.created_at           AS merchant_economic_created_at,
       activities.id                          AS activity_id,
       activities.cuui_id                     AS activity_cuui_id,
       activities.description                 AS activity_description,
       activities.status                      AS activity_status,
       activities.created_at                  AS activity_created_at
FROM core_merchant_economic_activities merchant_economic
         INNER JOIN core_economic_activities activities ON merchant_economic.economic_activity_id = activities.id
WHERE merchant_economic.merchant_id = ?
  AND merchant_economic.deleted_at IS NULL
ORDER BY merchant_economic.created_at DESC
LIMIT ? OFFSET ?;