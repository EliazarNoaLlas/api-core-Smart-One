INSERT INTO core_policies(id,
                          name,
                          description,
                          module_id,
                          merchant_id,
                          store_id,
                          level,
                          enable,
                          created_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
