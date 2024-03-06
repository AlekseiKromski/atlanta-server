package postgres

import (
	"fmt"
)

// Migration - provide ability to create migration to storage if needed
type Migration struct {
	Sql  string
	Name string
}

var migrations = []*Migration{
	// migrations table
	&Migration{
		Name: "create_migrations",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.migrations
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			name text COLLATE pg_catalog."default" NOT NULL,
			created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT groups_migrations_PK PRIMARY KEY (id)
		)`,
	},

	// devices table
	&Migration{
		Name: "create_devices",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.devices
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			description text COLLATE pg_catalog."default" NOT NULL,
			status boolean NOT NULL,
			created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			deleted_at timestamp with time zone,
			CONSTRAINT device_PK PRIMARY KEY (id)
		)`,
	},

	// datapoints table
	&Migration{
		Name: "create_datapoints",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.datapoints
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			value character varying(50) COLLATE pg_catalog."default" NOT NULL,
			type character varying(50) COLLATE pg_catalog."default" NOT NULL,
			unit character varying(50) COLLATE pg_catalog."default" NOT NULL,
			measurement_time timestamp with time zone NOT NULL,
			created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			deviceuuid uuid NOT NULL,
			CONSTRAINT datapoints_PK PRIMARY KEY (id),
			CONSTRAINT device_fk FOREIGN KEY (deviceuuid)
				REFERENCES public.devices (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)`,
	},

	// groups table
	&Migration{
		Name: "create_groups",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.groups
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			name character varying(50) COLLATE pg_catalog."default" NOT NULL,
			created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			deleted_at timestamp with time zone DEFAULT NULL,
			CONSTRAINT groups_PK PRIMARY KEY (id)
		)`,
	},

	// insert default groups
	&Migration{
		Name: "insert_default_groups",
		Sql: `
		INSERT INTO groups (name) VALUES ('admin');
		INSERT INTO groups (name) VALUES ('manager');
		INSERT INTO groups (name) VALUES ('viewer');
		`,
	},

	// users table
	&Migration{
		Name: "create_users",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.users
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			username character varying(120) COLLATE pg_catalog."default" NOT NULL,
			first_name character varying(120) COLLATE pg_catalog."default" NOT NULL,
			second_name character varying(120) COLLATE pg_catalog."default" NOT NULL,
			image character varying(120) COLLATE pg_catalog."default" NOT NULL,
			email character varying(120) COLLATE pg_catalog."default" NOT NULL,
			password character varying(120) COLLATE pg_catalog."default" NOT NULL,
			created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			deleted_at timestamp with time zone DEFAULT NULL,
			CONSTRAINT users_PK PRIMARY KEY (id)
		)`,
	},

	// users table
	&Migration{
		Name: "create_groups_users",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.groups_users
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			userUuid uuid NOT NULL,
			groupUuid uuid NOT NULL,
		    CONSTRAINT groups_users_PK PRIMARY KEY (id),
			CONSTRAINT groups_users_user_FK FOREIGN KEY (userUuid)
				REFERENCES public.users (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
		   	CONSTRAINT groups_users_group_FK FOREIGN KEY (groupUuid)
				REFERENCES public.groups (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION
		)`,
	},

	// insert admin user with admin group relation
	&Migration{
		Name: "insert_admin_user",
		Sql: `
		INSERT INTO users (username, first_name, second_name, image, email, password) VALUES ('admin', 'admin', 'admin', 'default_user.png', 'admin@admin.com', '');
		INSERT INTO groups_users (useruuid, groupuuid)
		VALUES (
				(SELECT id FROM users WHERE username = 'admin' LIMIT 1),
				(SELECT id FROM groups WHERE name = 'admin' LIMIT 1)
			   )
		`,
	},

	// add flags field
	&Migration{
		Name: "alter_flags_filed",
		Sql: `
		ALTER TABLE datapoints ADD flags VARCHAR(255) DEFAULT NULL;
		`,
	},

	// add label field
	&Migration{
		Name: "alter_label_field",
		Sql: `
		ALTER TABLE datapoints ADD label VARCHAR(50) DEFAULT NULL;
		`,
	},
}

func (p *Postgres) migrations() error {
	// Check migrations table
	rows, err := p.db.Query("SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = 'migrations')")
	if err != nil {
		return fmt.Errorf("cannot send request to check migrations tables: %v", err)
	}
	defer rows.Close()

	var exists string
	for rows.Next() {
		err := rows.Scan(&exists)
		if err != nil {
			return fmt.Errorf("cannot read response from database: %v", err)
		}
	}

	migrated := map[string]string{}

	if exists == "true" {
		rows, err := p.db.Query("SELECT name, created_at FROM migrations ORDER BY created_at DESC")
		if err != nil {
			return fmt.Errorf("cannot send request to check migrations tables: %v", err)
		}
		defer rows.Close()

		var name string
		var created_at string
		for rows.Next() {
			err := rows.Scan(&name, &created_at)
			if err != nil {
				return fmt.Errorf("cannot read response from database: %v", err)
			}

			migrated[name] = created_at
		}
	}

	for _, migration := range migrations {

		if len(migrated[migration.Name]) != 0 {
			continue // skip creation if alredy exists
		}

		_, err := p.db.Exec(migration.Sql)
		if err != nil {
			return fmt.Errorf("cannot run migrations [%s]: %v", migration.Name, err)
		}

		query := "INSERT INTO migrations (name) VALUES ($1)"
		if _, err := p.db.Exec(query, migration.Name); err != nil {
			return fmt.Errorf("cannot create new migration record: %v", err)
		}

		p.Log("successful executed sql command", migration.Name)
	}

	return nil
}
