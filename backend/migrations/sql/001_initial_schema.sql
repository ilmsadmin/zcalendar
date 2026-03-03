-- ====================================================================
-- Zplus Calendar – Initial Schema Migration
-- ====================================================================

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ── Users ───────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS users (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email           VARCHAR(255) NOT NULL UNIQUE,
    username        VARCHAR(50)  NOT NULL UNIQUE,
    password_hash   VARCHAR(255) NOT NULL,
    full_name       VARCHAR(100),
    avatar_url      VARCHAR(500),
    role            VARCHAR(20) DEFAULT 'user',
    is_active       BOOLEAN DEFAULT TRUE,
    preferences     JSONB DEFAULT '{}',

    -- OAuth fields
    oauth_provider  VARCHAR(20),
    oauth_id        VARCHAR(255),

    last_login_at   TIMESTAMP WITH TIME ZONE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);

-- ── Categories ──────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS categories (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name       VARCHAR(100) NOT NULL,
    color      VARCHAR(7) DEFAULT '#3B82F6',
    icon       VARCHAR(50),
    type       VARCHAR(20) NOT NULL, -- task, note, both
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_categories_user_id ON categories(user_id);

-- ── Tags ────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS tags (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name       VARCHAR(50) NOT NULL,
    color      VARCHAR(7) DEFAULT '#6B7280',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_tags_user_id ON tags(user_id);

-- ── Tasks ───────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS tasks (
    id               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id          UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id      UUID REFERENCES categories(id) ON DELETE SET NULL,
    parent_task_id   UUID REFERENCES tasks(id) ON DELETE CASCADE,

    title            VARCHAR(500) NOT NULL,
    description      TEXT,

    -- Scheduling
    due_date         DATE,
    due_time         TIME,
    start_date       DATE,
    start_time       TIME,

    -- Lunar date reference
    lunar_day        INT,
    lunar_month      INT,
    lunar_year       INT,
    is_lunar_based   BOOLEAN DEFAULT FALSE,

    -- Status & Priority
    status           VARCHAR(20) DEFAULT 'todo',
    priority         VARCHAR(10) DEFAULT 'medium',

    -- Recurrence
    is_recurring        BOOLEAN DEFAULT FALSE,
    recurrence_rule     VARCHAR(255),
    recurrence_end_date DATE,

    -- Progress
    progress         INT DEFAULT 0,
    completed_at     TIMESTAMP WITH TIME ZONE,

    -- Meta
    sort_order       INT DEFAULT 0,
    estimated_hours  DECIMAL(5,2),
    actual_hours     DECIMAL(5,2),

    created_at       TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at       TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at       TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_tasks_user_id ON tasks(user_id);
CREATE INDEX idx_tasks_category_id ON tasks(category_id);
CREATE INDEX idx_tasks_parent_task_id ON tasks(parent_task_id);
CREATE INDEX idx_tasks_due_date ON tasks(due_date);
CREATE INDEX idx_tasks_status ON tasks(status);
CREATE INDEX idx_tasks_deleted_at ON tasks(deleted_at);

-- ── Task Tags (join table) ─────────────────────────────────────────
CREATE TABLE IF NOT EXISTS task_tags (
    task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    tag_id  UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, tag_id)
);

-- ── Reminders ───────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS reminders (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id       UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    task_id       UUID REFERENCES tasks(id) ON DELETE CASCADE,
    title         VARCHAR(255) NOT NULL,
    description   TEXT,
    remind_at     TIMESTAMP WITH TIME ZONE NOT NULL,
    type          VARCHAR(20) DEFAULT 'once', -- once, daily, weekly, monthly, yearly
    is_active     BOOLEAN DEFAULT TRUE,
    is_sent       BOOLEAN DEFAULT FALSE,
    sent_at       TIMESTAMP WITH TIME ZONE,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_reminders_user_id ON reminders(user_id);
CREATE INDEX idx_reminders_remind_at ON reminders(remind_at);
CREATE INDEX idx_reminders_is_sent ON reminders(is_sent);

-- ── Vietnamese Holidays ─────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS vietnamese_holidays (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    holiday_date DATE NOT NULL,
    is_lunar    BOOLEAN DEFAULT FALSE,
    lunar_day   INT,
    lunar_month INT,
    is_official BOOLEAN DEFAULT TRUE,
    year        INT,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_vietnamese_holidays_date ON vietnamese_holidays(holiday_date);
CREATE INDEX idx_vietnamese_holidays_year ON vietnamese_holidays(year);

-- ── User Sessions ───────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS user_sessions (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id       UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_id      VARCHAR(255) NOT NULL UNIQUE,
    refresh_token VARCHAR(500) NOT NULL,
    user_agent    VARCHAR(500),
    ip_address    VARCHAR(45),
    is_active     BOOLEAN DEFAULT TRUE,
    expires_at    TIMESTAMP WITH TIME ZONE NOT NULL,
    last_used_at  TIMESTAMP WITH TIME ZONE,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_user_sessions_user_id ON user_sessions(user_id);
CREATE INDEX idx_user_sessions_token_id ON user_sessions(token_id);
CREATE INDEX idx_user_sessions_expires_at ON user_sessions(expires_at);
