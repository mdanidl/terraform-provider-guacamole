INSERT IGNORE INTO guacamole_user_permission (
            entity_id,
            permission,
            affected_user_id
        )
        SELECT DISTINCT
            permissions.entity_id,
            permissions.permission,
            affected_user.user_id
        FROM
             (
                SELECT 1         AS entity_id,
                       'READ'             AS permission,
                       'uname' AS affected_name
             UNION ALL
                SELECT 1         AS entity_id,
                       'UPDATE'             AS permission,
                       'uname' AS affected_name
             UNION ALL
                SELECT 1         AS entity_id,
                       'DELETE'             AS permission,
                       'uname' AS affected_name
             UNION ALL
                SELECT 1         AS entity_id,
                       'ADMINISTER'             AS permission,
                       'uname' AS affected_name
             UNION ALL
                SELECT 2         AS entity_id,
                       'READ'             AS permission,
                       'uname' AS affected_name
             )
        AS permissions
        JOIN guacamole_entity affected_entity ON
                affected_entity.name = permissions.affected_name
            AND affected_entity.type = 'USER'
        JOIN guacamole_user affected_user ON affected_user.entity_id = affected_en
tity.entity_id