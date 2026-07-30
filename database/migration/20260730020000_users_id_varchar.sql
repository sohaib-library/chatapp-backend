-- +goose Up
-- Drop all FK constraints that reference users.id or conversations.id
ALTER TABLE conversation_members DROP CONSTRAINT IF EXISTS conversation_members_conversation_id_fkey;
ALTER TABLE conversation_members DROP CONSTRAINT IF EXISTS conversation_members_user_id_fkey;
ALTER TABLE messages             DROP CONSTRAINT IF EXISTS messages_conversation_id_fkey;
ALTER TABLE messages             DROP CONSTRAINT IF EXISTS messages_sender_id_fkey;

-- Change all UUID columns to VARCHAR(36)
ALTER TABLE users                ALTER COLUMN id              TYPE VARCHAR(36) USING id::VARCHAR;
ALTER TABLE users                ALTER COLUMN id              DROP DEFAULT;

ALTER TABLE conversations        ALTER COLUMN id              TYPE VARCHAR(36) USING id::VARCHAR;
ALTER TABLE conversations        ALTER COLUMN id              DROP DEFAULT;

ALTER TABLE conversation_members ALTER COLUMN conversation_id TYPE VARCHAR(36) USING conversation_id::VARCHAR;
ALTER TABLE conversation_members ALTER COLUMN user_id         TYPE VARCHAR(36) USING user_id::VARCHAR;

ALTER TABLE messages             ALTER COLUMN id              TYPE VARCHAR(36) USING id::VARCHAR;
ALTER TABLE messages             ALTER COLUMN id              DROP DEFAULT;
ALTER TABLE messages             ALTER COLUMN conversation_id TYPE VARCHAR(36) USING conversation_id::VARCHAR;
ALTER TABLE messages             ALTER COLUMN sender_id       TYPE VARCHAR(36) USING sender_id::VARCHAR;

-- Re-add FK constraints
ALTER TABLE conversation_members
    ADD CONSTRAINT conversation_members_conversation_id_fkey
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE;

ALTER TABLE conversation_members
    ADD CONSTRAINT conversation_members_user_id_fkey
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE messages
    ADD CONSTRAINT messages_conversation_id_fkey
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE;

ALTER TABLE messages
    ADD CONSTRAINT messages_sender_id_fkey
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE conversation_members DROP CONSTRAINT IF EXISTS conversation_members_conversation_id_fkey;
ALTER TABLE conversation_members DROP CONSTRAINT IF EXISTS conversation_members_user_id_fkey;
ALTER TABLE messages             DROP CONSTRAINT IF EXISTS messages_conversation_id_fkey;
ALTER TABLE messages             DROP CONSTRAINT IF EXISTS messages_sender_id_fkey;

ALTER TABLE users                ALTER COLUMN id              TYPE UUID USING id::UUID;
ALTER TABLE users                ALTER COLUMN id              SET DEFAULT gen_random_uuid();

ALTER TABLE conversations        ALTER COLUMN id              TYPE UUID USING id::UUID;
ALTER TABLE conversations        ALTER COLUMN id              SET DEFAULT gen_random_uuid();

ALTER TABLE conversation_members ALTER COLUMN conversation_id TYPE UUID USING conversation_id::UUID;
ALTER TABLE conversation_members ALTER COLUMN user_id         TYPE UUID USING user_id::UUID;

ALTER TABLE messages             ALTER COLUMN id              TYPE UUID USING id::UUID;
ALTER TABLE messages             ALTER COLUMN id              SET DEFAULT gen_random_uuid();
ALTER TABLE messages             ALTER COLUMN conversation_id TYPE UUID USING conversation_id::UUID;
ALTER TABLE messages             ALTER COLUMN sender_id       TYPE UUID USING sender_id::UUID;

ALTER TABLE conversation_members
    ADD CONSTRAINT conversation_members_conversation_id_fkey
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE;

ALTER TABLE conversation_members
    ADD CONSTRAINT conversation_members_user_id_fkey
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE messages
    ADD CONSTRAINT messages_conversation_id_fkey
    FOREIGN KEY (conversation_id) REFERENCES conversations(id) ON DELETE CASCADE;

ALTER TABLE messages
    ADD CONSTRAINT messages_sender_id_fkey
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE;
