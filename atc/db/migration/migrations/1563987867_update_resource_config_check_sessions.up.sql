BEGIN;
  ALTER TABLE resource_config_check_sessions
   DROP CONSTRAINT resource_config_check_sessions_resource_config_id_fkey;

  ALTER TABLE resource_config_check_sessions
   DROP CONSTRAINT resource_config_check_sessions_worker_base_resource_type_id_fke;

  TRUNCATE TABLE resource_config_check_sessions CASCADE;

  ALTER TABLE resource_config_check_sessions
    DROP COLUMN resource_config_id,
    DROP COLUMN worker_base_resource_type_id;

  ALTER TABLE resource_config_check_sessions ADD COLUMN resource_config_scope_id integer;

  ALTER TABLE resource_config_check_sessions
    ADD CONSTRAINT resource_config_check_session_resource_config_scope_id_fkey FOREIGN KEY (resource_config_scope_id) REFERENCES resource_config_scopes(id) ON DELETE CASCADE;

  CREATE UNIQUE INDEX resource_config_check_sessions_uniq
    ON resource_config_check_sessions (resource_config_scope_id);

COMMIT;
