INSERT INTO guacamole_connection (
            connection_name,
            parent_id,
            protocol,
            max_connections,
            max_connections_per_user,
            proxy_hostname,
            proxy_port,
            proxy_encryption_method,
            connection_weight,
            failover_only
        )
        VALUES (
            'testing',
            null,
            'rdp',
            null,
            null,
            null,
            null,
            null,
            null,
            0
        )