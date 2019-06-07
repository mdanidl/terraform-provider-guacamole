INSERT INTO guacamole_user (
            entity_id,
            password_hash,
            password_salt,
            password_date,
            disabled,
            expired,
            access_window_start,
            access_window_end,
            valid_from,
            valid_until,
            timezone,
            full_name,
            email_address,
            organization,
            organizational_role
        )
        VALUES (
            2,
            x'C8CE46C593D226A315DDA25C6D96892ADAFB4D6EC77DA4FCE9A09B5366D68BAE',
            x'C5B2D7759A8FB58F3658F55CC830BA09C4164A4E8ED2C8AB3D76D8904299807B',
            '2019-01-21 19:57:10.188',
            0,
            0,
            null,
            null,
            null,
            null,
            null,
            'myfullname',
            null,
            null,
            null
        )