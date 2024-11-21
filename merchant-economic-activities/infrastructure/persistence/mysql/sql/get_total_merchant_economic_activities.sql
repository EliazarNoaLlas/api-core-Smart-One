SELECT COUNT(*)
FROM core_merchant_economic_activities merchant_economic
         INNER JOIN core_economic_activities activities ON merchant_economic.economic_activity_id = activities.id
WHERE merchant_economic.merchant_id=?
  AND merchant_economic.deleted_at IS NULL;