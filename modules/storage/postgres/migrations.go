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

	// insert admin user with admin group relation
	&Migration{
		Name: "insert_admin_user",
		Sql: `
		INSERT INTO users (username, first_name, second_name, image, email, password) VALUES ('admin', 'admin', 'admin', 'default_user.png', 'admin@admin.com', '');
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

	// create roles table
	&Migration{
		Name: "create_roles_table",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.roles
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			name VARCHAR(50) NOT NULL,
		    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			deleted_at timestamp with time zone DEFAULT NULL,
			CONSTRAINT roles_PK PRIMARY KEY (id)
		)
		`,
	},

	// create default role
	&Migration{
		Name: "create_default_role",
		Sql: `
		INSERT INTO roles (id, name) VALUES ('9349e8e0-9f69-4a97-a47f-85d8d55a4776','default')
		`,
	},

	// add role field to user with default role
	&Migration{
		Name: "alter_users_role_field",
		Sql: `
		ALTER TABLE users ADD role uuid NOT NULL DEFAULT '9349e8e0-9f69-4a97-a47f-85d8d55a4776';
		`,
	},

	// add constraint to user -> role FK
	&Migration{
		Name: "alter_users_role_field_fk",
		Sql: `
			ALTER TABLE users ADD CONSTRAINT user_role_FK FOREIGN KEY (role)
			REFERENCES public.roles (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION;
		`,
	},

	// create endpoints table
	&Migration{
		Name: "create_enpoints_table",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.endpoints
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			urn VARCHAR(300) NOT NULL,
		    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			deleted_at timestamp with time zone DEFAULT NULL,
			CONSTRAINT endpoints_PK PRIMARY KEY (id)
		)
		`,
	},

	// Create permission reference between endpoint and role
	&Migration{
		Name: "create_roles_endpoints",
		Sql: `
		CREATE TABLE IF NOT EXISTS public.roles_endpoints
		(
			id uuid NOT NULL DEFAULT gen_random_uuid(),
			roleUuid uuid NOT NULL,
			endpointUuid uuid NOT NULL,
		    CONSTRAINT roles_endpoints_PK PRIMARY KEY (id),
			CONSTRAINT roles_endpoints_role_FK FOREIGN KEY (roleUuid)
				REFERENCES public.roles (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
		   	CONSTRAINT roles_endpoints_endpoint_FK FOREIGN KEY (endpointUuid)
				REFERENCES public.endpoints (id) MATCH SIMPLE
				ON UPDATE NO ACTION
				ON DELETE NO ACTION,
		    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
			deleted_at timestamp with time zone DEFAULT NULL
		)
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
