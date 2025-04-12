# Coffee Tracker Data Models

This document outlines the database schema for the Coffee Tracker application using PostgreSQL.

## Core Entities

### User

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    display_name TEXT NOT NULL,
    bio TEXT,
    avatar_url TEXT,
    preferences JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_login_at TIMESTAMP WITH TIME ZONE
);

-- Index for email searching
CREATE INDEX idx_users_email ON users(email);
```

**Rules & Constraints:**
- Email must be unique and valid format
- Password must be securely hashed (bcrypt)
- Display name must be between 3-50 characters
- Preferences JSON can store user preferences like favorite brew methods, UI settings, etc.

### Coffee Bean

```sql
CREATE TABLE coffee_beans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    origin TEXT,
    roaster TEXT,
    roast_date DATE,
    roast_level TEXT CHECK (roast_level IN ('light', 'medium', 'medium-dark', 'dark', 'unknown')),
    flavor_notes TEXT[],
    bean_species TEXT CHECK (bean_species IN ('arabica', 'robusta', 'blend', 'other', 'unknown')),
    processing_method TEXT,
    altitude TEXT,
    purchase_date DATE,
    price DECIMAL(10, 2),
    quantity_grams INTEGER,
    image_urls TEXT[],
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    is_active BOOLEAN DEFAULT TRUE,
    is_favorite BOOLEAN DEFAULT FALSE
);

-- Indexes for common queries
CREATE INDEX idx_coffee_beans_user_id ON coffee_beans(user_id);
CREATE INDEX idx_coffee_beans_roast_level ON coffee_beans(roast_level);
CREATE INDEX idx_coffee_beans_is_active ON coffee_beans(is_active);
```

**Rules & Constraints:**
- Bean must belong to a user
- Roast level must be one of predefined values
- Bean species must be one of predefined values
- Quantity in grams must be positive
- Array of image URLs for multiple images of packaging/beans

### Recipe

```sql
CREATE TABLE recipes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    brew_method TEXT NOT NULL,
    coffee_dose_grams DECIMAL(6, 2) NOT NULL,
    water_amount_grams DECIMAL(6, 2) NOT NULL,
    grind_size TEXT NOT NULL,
    grinder_setting TEXT,
    water_temperature DECIMAL(4, 1),
    brew_time_seconds INTEGER,
    description TEXT,
    instructions TEXT,
    flavor_tags TEXT[],
    image_urls TEXT[],
    is_public BOOLEAN DEFAULT FALSE,
    is_favorite BOOLEAN DEFAULT FALSE,
    source TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Indexes for common queries
CREATE INDEX idx_recipes_user_id ON recipes(user_id);
CREATE INDEX idx_recipes_brew_method ON recipes(brew_method);
CREATE INDEX idx_recipes_is_public ON recipes(is_public);
```

**Rules & Constraints:**
- Recipe must belong to a user
- Coffee dose and water amount must be positive
- Brew time must be positive
- Flavor tags array for categorization
- Public flag controls visibility in community

### Brew Log

```sql
CREATE TABLE brew_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    recipe_id UUID REFERENCES recipes(id) ON DELETE SET NULL,
    bean_id UUID REFERENCES coffee_beans(id) ON DELETE SET NULL,
    brew_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    brew_method TEXT NOT NULL,
    coffee_dose_grams DECIMAL(6, 2) NOT NULL,
    water_amount_grams DECIMAL(6, 2) NOT NULL,
    grind_size TEXT NOT NULL,
    grinder_setting TEXT,
    water_temperature DECIMAL(4, 1),
    brew_time_seconds INTEGER,
    notes TEXT,
    taste_rating INTEGER CHECK (taste_rating BETWEEN 1 AND 10),
    aroma_rating INTEGER CHECK (aroma_rating BETWEEN 1 AND 10),
    body_rating INTEGER CHECK (body_rating BETWEEN 1 AND 10),
    acidity_rating INTEGER CHECK (acidity_rating BETWEEN 1 AND 10),
    overall_rating INTEGER CHECK (overall_rating BETWEEN 1 AND 10),
    flavor_notes TEXT[],
    image_urls TEXT[],
    is_public BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Indexes for common queries
CREATE INDEX idx_brew_logs_user_id ON brew_logs(user_id);
CREATE INDEX idx_brew_logs_recipe_id ON brew_logs(recipe_id);
CREATE INDEX idx_brew_logs_bean_id ON brew_logs(bean_id);
CREATE INDEX idx_brew_logs_brew_date ON brew_logs(brew_date);
CREATE INDEX idx_brew_logs_overall_rating ON brew_logs(overall_rating);
```

**Rules & Constraints:**
- Brew log must belong to a user
- Recipe and bean references are optional but should be present when possible
- All ratings must be between 1-10
- Brew parameters (dose, water, time) must be positive
- Contains all parameters even if using a recipe (to track deviations)

### Social & Community

```sql
CREATE TABLE follows (
    follower_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    followed_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (follower_id, followed_id)
);

CREATE TABLE likes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_type TEXT NOT NULL CHECK (content_type IN ('recipe', 'brew_log')),
    content_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content_type TEXT NOT NULL CHECK (content_type IN ('recipe', 'brew_log')),
    content_id UUID NOT NULL,
    parent_comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
    comment_text TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Indexes for social features
CREATE INDEX idx_likes_content ON likes(content_type, content_id);
CREATE INDEX idx_comments_content ON comments(content_type, content_id);
CREATE INDEX idx_comments_parent ON comments(parent_comment_id);
```

**Rules & Constraints:**
- Content type must be one of predefined values
- A user cannot follow themselves
- Comment text cannot be empty

### Gamification

```sql
CREATE TABLE challenges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    criteria JSONB NOT NULL,
    image_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE user_challenges (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    status TEXT NOT NULL CHECK (status IN ('in_progress', 'completed', 'failed')),
    progress JSONB,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, challenge_id)
);

CREATE TABLE badges (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    criteria JSONB NOT NULL,
    image_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE user_badges (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    badge_id UUID NOT NULL REFERENCES badges(id) ON DELETE CASCADE,
    earned_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, badge_id)
);
```

**Rules & Constraints:**
- Challenge end date must be after start date
- Challenge criteria stored as JSON for flexibility
- Badge criteria stored as JSON for flexibility

## Data Relationships

### One-to-Many Relationships
- User → Coffee Beans (one user has many coffee beans)
- User → Recipes (one user creates many recipes)
- User → Brew Logs (one user records many brew logs)
- Recipe → Brew Logs (one recipe can be used for many brew logs)
- Coffee Bean → Brew Logs (one coffee bean can be used in many brew logs)

### Many-to-Many Relationships
- Users ↔ Users (followers/following)
- Users ↔ Recipes/Brew Logs (likes)
- Users ↔ Badges (through user_badges)
- Users ↔ Challenges (through user_challenges)

## Indexes and Performance Considerations

1. All foreign keys have corresponding indexes
2. Text fields that will be searched frequently have indexes
3. Common filter conditions (is_active, is_public) have indexes
4. Date fields used for sorting have indexes

## Data Types and Validation Rules

1. **IDs**: UUID v4 for globally unique identifiers
2. **Timestamps**: Always use TIMESTAMP WITH TIME ZONE
3. **Enumerated Types**: Implemented as TEXT with CHECK constraints
4. **Floating Point**: Use DECIMAL for precision where needed
5. **Arrays**: Used for simple lists (flavor_notes, image_urls)
6. **JSON/JSONB**: Used for flexible/schemaless data (preferences, criteria)

## Soft Delete Strategy

For entities where historical data is important (beans, recipes, brew logs), use the is_active flag rather than physical deletion.

## Additional Considerations

1. **Versioning**: Recipe versions could be implemented for tracking changes/improvements
2. **Archiving**: Consider archiving old brew logs or beans after a certain time period
3. **Data Migration**: Plan for schema evolution with minimal disruption
4. **Data Partitioning**: For larger tables like brew_logs, consider partitioning by date 