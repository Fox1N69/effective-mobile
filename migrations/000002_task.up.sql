CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT, 
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
