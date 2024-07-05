CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    total_hours DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Index on user_id for faster queries
CREATE INDEX idx_tasks_user_id ON tasks (user_id);

-- Add foreign key constraint
ALTER TABLE tasks
ADD CONSTRAINT fk_tasks_user_id
FOREIGN KEY (user_id) REFERENCES users(id)
ON DELETE CASCADE;

-- Comments explaining the table structure
COMMENT ON TABLE tasks IS 'Table containing user tasks';

-- Comments for each column
COMMENT ON COLUMN tasks.id IS 'Primary key for tasks table';
COMMENT ON COLUMN tasks.user_id IS 'Foreign key referencing users table';
COMMENT ON COLUMN tasks.name IS 'Name of the task';
COMMENT ON COLUMN tasks.start_time IS 'Start time of the task';
COMMENT ON COLUMN tasks.end_time IS 'End time of the task';
COMMENT ON COLUMN tasks.total_hours IS 'Total hours spent on the task';
COMMENT ON COLUMN tasks.created_at IS 'Timestamp when the task was created';
COMMENT ON COLUMN tasks.updated_at IS 'Timestamp when the task was last updated';