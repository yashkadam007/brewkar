# Coffee Tracker API Documentation

This document outlines the REST API endpoints for the Coffee Tracker application.

## Base URL

All endpoints are relative to the base URL: `https://api.brewkar.com/v1`

## Authentication

### Authentication Endpoints

#### POST /auth/register

Register a new user.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123",
  "displayName": "Coffee Lover"
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "user": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "email": "user@example.com",
      "displayName": "Coffee Lover",
      "createdAt": "2023-08-01T12:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Algorithm:**
1. Validate email and password requirements
2. Check if email is already registered
3. Hash password using bcrypt
4. Create new user record in database
5. Generate JWT token
6. Return user details and token

#### POST /auth/login

Authenticate a user.

**Request:**
```json
{
  "email": "user@example.com",
  "password": "securePassword123"
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "user": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "email": "user@example.com",
      "displayName": "Coffee Lover",
      "lastLoginAt": "2023-08-01T12:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Algorithm:**
1. Validate email format
2. Find user by email
3. Compare password hash
4. Update last login timestamp
5. Generate JWT token
6. Return user details and token

#### POST /auth/refresh

Refresh authentication token.

**Request:**
```json
{
  "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Algorithm:**
1. Validate refresh token
2. Check token expiry and blacklist status
3. Generate new access token and refresh token
4. Return new tokens

### Authentication Middleware

All other endpoints require a valid JWT token in the Authorization header:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## User Endpoints

#### GET /users/me

Get current user profile.

**Response:**
```json
{
  "status": "success",
  "data": {
    "user": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "email": "user@example.com",
      "displayName": "Coffee Lover",
      "bio": "Coffee enthusiast from Seattle",
      "avatarUrl": "https://example.com/avatar.jpg",
      "preferences": {
        "favoriteBrewMethods": ["aeropress", "french-press"],
        "defaultTemperature": "93",
        "temperatureUnit": "celsius",
        "weightUnit": "grams"
      },
      "createdAt": "2023-07-01T12:00:00Z",
      "updatedAt": "2023-08-01T12:00:00Z",
      "lastLoginAt": "2023-08-01T12:00:00Z"
    }
  }
}
```

#### PUT /users/me

Update current user profile.

**Request:**
```json
{
  "displayName": "Coffee Master",
  "bio": "Coffee enthusiast from Portland",
  "preferences": {
    "favoriteBrewMethods": ["aeropress", "v60"],
    "defaultTemperature": "95",
    "temperatureUnit": "celsius",
    "weightUnit": "grams"
  }
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "user": {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "displayName": "Coffee Master",
      "bio": "Coffee enthusiast from Portland",
      "preferences": {
        "favoriteBrewMethods": ["aeropress", "v60"],
        "defaultTemperature": "95",
        "temperatureUnit": "celsius",
        "weightUnit": "grams"
      },
      "updatedAt": "2023-08-01T13:00:00Z"
    }
  }
}
```

## Coffee Bean Endpoints

#### GET /beans

Get all coffee beans for the current user.

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)
- `sort`: Sort field (default: createdAt)
- `order`: Sort order (asc/desc, default: desc)
- `isActive`: Filter by active status (true/false)
- `search`: Search term for name/origin/roaster
- `roastLevel`: Filter by roast level

**Response:**
```json
{
  "status": "success",
  "data": {
    "beans": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "Ethiopia Yirgacheffe",
        "origin": "Ethiopia",
        "roaster": "Stumptown",
        "roastDate": "2023-07-15",
        "roastLevel": "light",
        "flavorNotes": ["floral", "citrus", "berry"],
        "beanSpecies": "arabica",
        "processingMethod": "washed",
        "altitude": "1800-2200m",
        "purchaseDate": "2023-07-20",
        "price": 18.99,
        "quantityGrams": 250,
        "imageUrls": ["https://example.com/bean1.jpg"],
        "isActive": true,
        "isFavorite": true,
        "createdAt": "2023-07-20T12:00:00Z",
        "updatedAt": "2023-07-20T12:00:00Z"
      }
      // More beans...
    ],
    "pagination": {
      "total": 42,
      "page": 1,
      "limit": 20,
      "pages": 3
    }
  }
}
```

#### POST /beans

Create a new coffee bean entry.

**Request:**
```json
{
  "name": "Colombia Huila",
  "origin": "Colombia",
  "roaster": "Blue Bottle",
  "roastDate": "2023-07-25",
  "roastLevel": "medium",
  "flavorNotes": ["chocolate", "caramel", "nutty"],
  "beanSpecies": "arabica",
  "processingMethod": "washed",
  "altitude": "1700-2000m",
  "purchaseDate": "2023-07-28",
  "price": 16.50,
  "quantityGrams": 250,
  "isFavorite": false
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "bean": {
      "id": "123e4567-e89b-12d3-a456-426614174001",
      "name": "Colombia Huila",
      "origin": "Colombia",
      "roaster": "Blue Bottle",
      "roastDate": "2023-07-25",
      "roastLevel": "medium",
      "flavorNotes": ["chocolate", "caramel", "nutty"],
      "beanSpecies": "arabica",
      "processingMethod": "washed",
      "altitude": "1700-2000m",
      "purchaseDate": "2023-07-28",
      "price": 16.50,
      "quantityGrams": 250,
      "imageUrls": [],
      "isActive": true,
      "isFavorite": false,
      "createdAt": "2023-08-01T14:00:00Z",
      "updatedAt": "2023-08-01T14:00:00Z"
    }
  }
}
```

**Algorithm:**
1. Validate request payload
2. Create new bean record in database
3. Associate with current user
4. Return created bean entry

#### GET /beans/:id

Get a specific coffee bean by ID.

**Response:**
```json
{
  "status": "success",
  "data": {
    "bean": {
      "id": "123e4567-e89b-12d3-a456-426614174001",
      "name": "Colombia Huila",
      "origin": "Colombia",
      "roaster": "Blue Bottle",
      "roastDate": "2023-07-25",
      "roastLevel": "medium",
      "flavorNotes": ["chocolate", "caramel", "nutty"],
      "beanSpecies": "arabica",
      "processingMethod": "washed",
      "altitude": "1700-2000m",
      "purchaseDate": "2023-07-28",
      "price": 16.50,
      "quantityGrams": 250,
      "imageUrls": [],
      "isActive": true,
      "isFavorite": false,
      "createdAt": "2023-08-01T14:00:00Z",
      "updatedAt": "2023-08-01T14:00:00Z"
    }
  }
}
```

#### PUT /beans/:id

Update a coffee bean entry.

**Request:**
```json
{
  "quantityGrams": 220,
  "isFavorite": true
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "bean": {
      "id": "123e4567-e89b-12d3-a456-426614174001",
      "name": "Colombia Huila",
      "quantityGrams": 220,
      "isFavorite": true,
      "updatedAt": "2023-08-01T15:00:00Z"
      // Other fields remain unchanged
    }
  }
}
```

#### DELETE /beans/:id

Delete (or soft delete) a coffee bean entry.

**Response:**
```json
{
  "status": "success",
  "data": null
}
```

**Algorithm:**
1. Find bean by ID and verify ownership
2. Perform soft delete (update isActive to false)
3. Return success response

#### POST /beans/:id/images

Upload images for a coffee bean.

**Request:** Multipart form data with image files

**Response:**
```json
{
  "status": "success",
  "data": {
    "imageUrls": [
      "https://storage.brewkar.com/beans/123e4567-e89b-12d3-a456-426614174001/image1.jpg",
      "https://storage.brewkar.com/beans/123e4567-e89b-12d3-a456-426614174001/image2.jpg"
    ]
  }
}
```

**Algorithm:**
1. Validate image files (size, type)
2. Upload to cloud storage (S3)
3. Update bean record with new image URLs
4. Return updated image URLs

## Recipe Endpoints

#### GET /recipes

Get all recipes for the current user.

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)
- `sort`: Sort field (default: createdAt)
- `order`: Sort order (asc/desc, default: desc)
- `search`: Search term for name/description
- `brewMethod`: Filter by brew method
- `isPublic`: Filter by public status (true/false)

**Response:**
```json
{
  "status": "success",
  "data": {
    "recipes": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174002",
        "name": "My Aeropress Recipe",
        "brewMethod": "aeropress",
        "coffeeDoseGrams": 17.0,
        "waterAmountGrams": 250.0,
        "grindSize": "medium-fine",
        "grinderSetting": "Timemore C2: 15 clicks",
        "waterTemperature": 92.5,
        "brewTimeSeconds": 90,
        "description": "My go-to Aeropress recipe for light roasts",
        "instructions": "1. Rinse filter\n2. Add coffee\n3. Add water\n4. Stir 10 times\n5. Press after 90 seconds",
        "flavorTags": ["bright", "balanced", "fruity"],
        "imageUrls": ["https://example.com/recipe1.jpg"],
        "isPublic": true,
        "isFavorite": true,
        "source": "Modified from James Hoffmann method",
        "createdAt": "2023-07-10T12:00:00Z",
        "updatedAt": "2023-07-15T12:00:00Z"
      }
      // More recipes...
    ],
    "pagination": {
      "total": 15,
      "page": 1,
      "limit": 20,
      "pages": 1
    }
  }
}
```

#### POST /recipes

Create a new recipe.

**Request:**
```json
{
  "name": "V60 Technique",
  "brewMethod": "v60",
  "coffeeDoseGrams": 22.0,
  "waterAmountGrams": 360.0,
  "grindSize": "medium",
  "grinderSetting": "Timemore C2: 20 clicks",
  "waterTemperature": 94.0,
  "brewTimeSeconds": 180,
  "description": "Classic V60 pour-over technique",
  "instructions": "1. Rinse filter\n2. Add coffee\n3. Bloom with 50g water for 30s\n4. Pour to 180g at 1:00\n5. Pour to 360g at 1:45\n6. Drawdown complete by 3:00",
  "flavorTags": ["clean", "balanced", "tea-like"],
  "isPublic": false,
  "source": "Modified James Hoffmann V60 method"
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "recipe": {
      "id": "123e4567-e89b-12d3-a456-426614174003",
      "name": "V60 Technique",
      "brewMethod": "v60",
      "coffeeDoseGrams": 22.0,
      "waterAmountGrams": 360.0,
      "grindSize": "medium",
      "grinderSetting": "Timemore C2: 20 clicks",
      "waterTemperature": 94.0,
      "brewTimeSeconds": 180,
      "description": "Classic V60 pour-over technique",
      "instructions": "1. Rinse filter\n2. Add coffee\n3. Bloom with 50g water for 30s\n4. Pour to 180g at 1:00\n5. Pour to 360g at 1:45\n6. Drawdown complete by 3:00",
      "flavorTags": ["clean", "balanced", "tea-like"],
      "imageUrls": [],
      "isPublic": false,
      "isFavorite": false,
      "source": "Modified James Hoffmann V60 method",
      "createdAt": "2023-08-01T16:00:00Z",
      "updatedAt": "2023-08-01T16:00:00Z"
    }
  }
}
```

**Algorithm:**
1. Validate request payload
2. Create new recipe record in database
3. Associate with current user
4. Return created recipe

#### Endpoints for GET /recipes/:id, PUT /recipes/:id, DELETE /recipes/:id, and POST /recipes/:id/images

These follow the same patterns as the corresponding bean endpoints.

#### GET /recipes/public

Get public recipes from all users.

**Query Parameters:**
- Same as GET /recipes, plus:
- `userId`: Filter by specific user
- `rating`: Filter by minimum overall rating

**Algorithm:**
1. Apply filters for public recipes only
2. Add popularity metrics (likes, comments, average ratings)
3. Paginate results
4. Return filtered recipes

## Brew Log Endpoints

#### GET /brew-logs

Get all brew logs for the current user.

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)
- `sort`: Sort field (default: brewDate)
- `order`: Sort order (asc/desc, default: desc)
- `search`: Search term for notes
- `brewMethod`: Filter by brew method
- `beanId`: Filter by specific bean
- `recipeId`: Filter by specific recipe
- `minRating`: Filter by minimum overall rating

**Response:**
```json
{
  "status": "success",
  "data": {
    "brewLogs": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174004",
        "brewDate": "2023-08-01T08:30:00Z",
        "brewMethod": "aeropress",
        "beanId": "123e4567-e89b-12d3-a456-426614174001",
        "beanName": "Colombia Huila",
        "recipeId": "123e4567-e89b-12d3-a456-426614174002",
        "recipeName": "My Aeropress Recipe",
        "coffeeDoseGrams": 17.0,
        "waterAmountGrams": 250.0,
        "grindSize": "medium-fine",
        "grinderSetting": "Timemore C2: 14 clicks",
        "waterTemperature": 93.0,
        "brewTimeSeconds": 95,
        "notes": "Slightly finer grind than my usual recipe, resulted in more body",
        "tasteRating": 8,
        "aromaRating": 7,
        "bodyRating": 9,
        "acidityRating": 6,
        "overallRating": 8,
        "flavorNotes": ["chocolate", "nutty", "tangy"],
        "imageUrls": ["https://example.com/brew1.jpg"],
        "isPublic": false,
        "createdAt": "2023-08-01T08:45:00Z",
        "updatedAt": "2023-08-01T08:45:00Z"
      }
      // More brew logs...
    ],
    "pagination": {
      "total": 32,
      "page": 1,
      "limit": 20,
      "pages": 2
    }
  }
}
```

#### POST /brew-logs

Create a new brew log.

**Request:**
```json
{
  "brewDate": "2023-08-02T07:15:00Z",
  "brewMethod": "v60",
  "beanId": "123e4567-e89b-12d3-a456-426614174001",
  "recipeId": "123e4567-e89b-12d3-a456-426614174003",
  "coffeeDoseGrams": 22.0,
  "waterAmountGrams": 360.0,
  "grindSize": "medium",
  "grinderSetting": "Timemore C2: 19 clicks",
  "waterTemperature": 94.5,
  "brewTimeSeconds": 190,
  "notes": "Slightly hotter water than usual recipe, improved extraction",
  "tasteRating": 9,
  "aromaRating": 8,
  "bodyRating": 7,
  "acidityRating": 8,
  "overallRating": 9,
  "flavorNotes": ["caramel", "cherry", "balanced"],
  "isPublic": true
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "brewLog": {
      "id": "123e4567-e89b-12d3-a456-426614174005",
      "brewDate": "2023-08-02T07:15:00Z",
      "brewMethod": "v60",
      "beanId": "123e4567-e89b-12d3-a456-426614174001",
      "beanName": "Colombia Huila",
      "recipeId": "123e4567-e89b-12d3-a456-426614174003",
      "recipeName": "V60 Technique",
      "coffeeDoseGrams": 22.0,
      "waterAmountGrams": 360.0,
      "grindSize": "medium",
      "grinderSetting": "Timemore C2: 19 clicks",
      "waterTemperature": 94.5,
      "brewTimeSeconds": 190,
      "notes": "Slightly hotter water than usual recipe, improved extraction",
      "tasteRating": 9,
      "aromaRating": 8,
      "bodyRating": 7,
      "acidityRating": 8,
      "overallRating": 9,
      "flavorNotes": ["caramel", "cherry", "balanced"],
      "imageUrls": [],
      "isPublic": true,
      "createdAt": "2023-08-02T07:20:00Z",
      "updatedAt": "2023-08-02T07:20:00Z"
    }
  }
}
```

**Algorithm:**
1. Validate request payload
2. Verify bean and recipe IDs (if provided)
3. Create new brew log in database
4. Update bean quantity if tracking is enabled (optional feature)
5. Return created brew log with related entity names

#### Endpoints for GET /brew-logs/:id, PUT /brew-logs/:id, DELETE /brew-logs/:id, and POST /brew-logs/:id/images

These follow the same patterns as the corresponding bean endpoints.

## Analytics Endpoints

#### GET /analytics/brew-stats

Get statistical data about brew logs.

**Query Parameters:**
- `period`: Time period for analysis (week, month, year, all)
- `beanId`: Filter by specific bean (optional)
- `brewMethod`: Filter by brew method (optional)

**Response:**
```json
{
  "status": "success",
  "data": {
    "totalBrews": 87,
    "averageRating": 7.8,
    "ratingDistribution": {
      "1": 0,
      "2": 1,
      "3": 2,
      "4": 3,
      "5": 8,
      "6": 12,
      "7": 18,
      "8": 25,
      "9": 15,
      "10": 3
    },
    "brewMethodDistribution": {
      "aeropress": 32,
      "v60": 28,
      "french-press": 15,
      "espresso": 12
    },
    "favoriteGrindSize": "medium-fine",
    "averageBrewParameters": {
      "coffeeDoseGrams": 18.5,
      "waterAmountGrams": 280.3,
      "waterTemperature": 93.2,
      "brewTimeSeconds": 152
    },
    "timeOfDayDistribution": {
      "morning": 65,
      "afternoon": 20,
      "evening": 2
    },
    "recentTrend": {
      "dates": ["2023-07-01", "2023-07-08", "2023-07-15", "2023-07-22", "2023-07-29"],
      "ratings": [7.2, 7.5, 7.8, 8.1, 8.4]
    }
  }
}
```

**Algorithm:**
1. Apply filters for time period and other parameters
2. Calculate aggregate statistics:
   - Count total brews
   - Calculate average ratings
   - Group brews by method, time of day, rating
   - Calculate average parameters
   - Determine time trends
3. Return analytics object

#### GET /analytics/correlations

Get correlation analysis between brew parameters and ratings.

**Query Parameters:**
- `parameter`: Parameter to analyze (grindSize, waterTemperature, brewTime, all)
- `brewMethod`: Filter by brew method (optional)

**Response:**
```json
{
  "status": "success",
  "data": {
    "correlations": [
      {
        "parameter": "waterTemperature",
        "correlation": 0.65,
        "optimalRange": {
          "min": 93.0,
          "max": 95.0
        },
        "dataPoints": [
          {"value": 90.0, "rating": 6.5},
          {"value": 92.0, "rating": 7.2},
          {"value": 94.0, "rating": 8.7},
          {"value": 96.0, "rating": 7.8}
          // More data points...
        ]
      },
      {
        "parameter": "grindSize",
        "correlation": 0.42,
        "optimalRange": {
          "value": "medium-fine"
        },
        "dataPoints": [
          {"value": "extra-fine", "rating": 5.8},
          {"value": "fine", "rating": 7.4},
          {"value": "medium-fine", "rating": 8.6},
          {"value": "medium", "rating": 7.9},
          {"value": "coarse", "rating": 6.2}
        ]
      },
      {
        "parameter": "brewTimeSeconds",
        "correlation": 0.38,
        "optimalRange": {
          "min": 85,
          "max": 110
        },
        "dataPoints": [
          {"value": 60, "rating": 6.1},
          {"value": 90, "rating": 8.5},
          {"value": 120, "rating": 7.8},
          {"value": 150, "rating": 6.9}
          // More data points...
        ]
      }
    ]
  }
}
```

**Algorithm:**
1. Apply filters for brew method
2. For each parameter:
   - Calculate correlation coefficient with overall rating
   - Identify optimal ranges based on highest ratings
   - Group data into meaningful clusters
   - Create visualization-ready data points
3. Return correlation analysis

## Social Endpoints

#### GET /social/feed

Get social feed with public brew logs and recipes.

**Query Parameters:**
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)
- `contentType`: Type of content (brews, recipes, all)
- `following`: Show only content from followed users (true/false)

**Response:**
```json
{
  "status": "success",
  "data": {
    "feed": [
      {
        "type": "brew",
        "id": "123e4567-e89b-12d3-a456-426614174006",
        "userId": "223e4567-e89b-12d3-a456-426614174000",
        "userName": "CoffeeExpert",
        "userAvatar": "https://example.com/avatar1.jpg",
        "content": {
          "brewMethod": "chemex",
          "beanName": "Ethiopian Natural",
          "overallRating": 9,
          "flavorNotes": ["blueberry", "chocolate", "floral"],
          "imageUrl": "https://example.com/brew2.jpg"
        },
        "likes": 12,
        "comments": 3,
        "createdAt": "2023-08-01T10:00:00Z",
        "isLiked": false
      },
      {
        "type": "recipe",
        "id": "123e4567-e89b-12d3-a456-426614174007",
        "userId": "323e4567-e89b-12d3-a456-426614174000",
        "userName": "BaristaChamp",
        "userAvatar": "https://example.com/avatar2.jpg",
        "content": {
          "name": "Ultimate V60 Technique",
          "brewMethod": "v60",
          "description": "My competition-winning V60 recipe",
          "imageUrl": "https://example.com/recipe2.jpg"
        },
        "likes": 45,
        "comments": 8,
        "createdAt": "2023-07-28T14:00:00Z",
        "isLiked": true
      }
      // More feed items...
    ],
    "pagination": {
      "total": 324,
      "page": 1,
      "limit": 20,
      "pages": 17
    }
  }
}
```

**Algorithm:**
1. Determine content types to include
2. Filter by followed users if requested
3. Combine and sort brew logs and recipes by date
4. Enhance with user details, like counts and comment counts
5. Mark items liked by current user
6. Paginate results
7. Return feed items

#### POST /social/follow/:userId

Follow a user.

**Response:**
```json
{
  "status": "success",
  "data": {
    "following": true
  }
}
```

#### DELETE /social/follow/:userId

Unfollow a user.

**Response:**
```json
{
  "status": "success",
  "data": {
    "following": false
  }
}
```

#### GET /social/following

Get users being followed by current user.

**Response:**
```json
{
  "status": "success",
  "data": {
    "following": [
      {
        "id": "223e4567-e89b-12d3-a456-426614174000",
        "displayName": "CoffeeExpert",
        "avatarUrl": "https://example.com/avatar1.jpg",
        "followedAt": "2023-07-15T12:00:00Z"
      },
      {
        "id": "323e4567-e89b-12d3-a456-426614174000",
        "displayName": "BaristaChamp",
        "avatarUrl": "https://example.com/avatar2.jpg",
        "followedAt": "2023-07-20T12:00:00Z"
      }
      // More followed users...
    ]
  }
}
```

#### GET /social/followers

Get users following the current user.

**Response:** Similar to the following endpoint.

#### POST /social/like

Like a brew log or recipe.

**Request:**
```json
{
  "contentType": "recipe",
  "contentId": "123e4567-e89b-12d3-a456-426614174007"
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "liked": true,
    "likeCount": 46
  }
}
```

#### DELETE /social/like/:id

Remove a like.

**Response:**
```json
{
  "status": "success",
  "data": {
    "liked": false,
    "likeCount": 45
  }
}
```

#### POST /social/comments

Add a comment to a brew log or recipe.

**Request:**
```json
{
  "contentType": "recipe",
  "contentId": "123e4567-e89b-12d3-a456-426614174007",
  "commentText": "I tried this recipe and it was amazing! I adjusted the water temperature to 95°C and it worked perfectly.",
  "parentCommentId": null
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "comment": {
      "id": "123e4567-e89b-12d3-a456-426614174008",
      "userId": "123e4567-e89b-12d3-a456-426614174000",
      "userName": "Coffee Lover",
      "userAvatar": "https://example.com/avatar.jpg",
      "contentType": "recipe",
      "contentId": "123e4567-e89b-12d3-a456-426614174007",
      "parentCommentId": null,
      "commentText": "I tried this recipe and it was amazing! I adjusted the water temperature to 95°C and it worked perfectly.",
      "createdAt": "2023-08-02T09:00:00Z"
    }
  }
}
```

#### GET /social/comments

Get comments for a brew log or recipe.

**Query Parameters:**
- `contentType`: Type of content (brew, recipe)
- `contentId`: ID of the content
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 20)

**Response:**
```json
{
  "status": "success",
  "data": {
    "comments": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174008",
        "userId": "123e4567-e89b-12d3-a456-426614174000",
        "userName": "Coffee Lover",
        "userAvatar": "https://example.com/avatar.jpg",
        "contentType": "recipe",
        "contentId": "123e4567-e89b-12d3-a456-426614174007",
        "parentCommentId": null,
        "commentText": "I tried this recipe and it was amazing! I adjusted the water temperature to 95°C and it worked perfectly.",
        "createdAt": "2023-08-02T09:00:00Z",
        "replies": [
          {
            "id": "123e4567-e89b-12d3-a456-426614174009",
            "userId": "323e4567-e89b-12d3-a456-426614174000",
            "userName": "BaristaChamp",
            "userAvatar": "https://example.com/avatar2.jpg",
            "contentType": "recipe",
            "contentId": "123e4567-e89b-12d3-a456-426614174007",
            "parentCommentId": "123e4567-e89b-12d3-a456-426614174008",
            "commentText": "Thanks for trying it! The higher temperature works well for darker roasts.",
            "createdAt": "2023-08-02T09:30:00Z",
            "replies": []
          }
        ]
      }
      // More comments...
    ],
    "pagination": {
      "total": 8,
      "page": 1,
      "limit": 20,
      "pages": 1
    }
  }
}
```

## Challenge & Gamification Endpoints

#### GET /challenges

Get active challenges.

**Response:**
```json
{
  "status": "success",
  "data": {
    "challenges": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174010",
        "title": "Brew Around the World",
        "description": "Try brewing with beans from 5 different countries",
        "startDate": "2023-08-01T00:00:00Z",
        "endDate": "2023-08-31T23:59:59Z",
        "imageUrl": "https://example.com/challenge1.jpg",
        "progress": {
          "status": "in_progress",
          "completed": 3,
          "total": 5,
          "details": [
            {"country": "Ethiopia", "completed": true},
            {"country": "Colombia", "completed": true},
            {"country": "Brazil", "completed": true},
            {"country": "Kenya", "completed": false},
            {"country": "Guatemala", "completed": false}
          ]
        },
        "participants": 156,
        "completions": 42
      }
      // More challenges...
    ]
  }
}
```

#### GET /badges

Get user's earned badges.

**Response:**
```json
{
  "status": "success",
  "data": {
    "badges": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174011",
        "name": "Bean Explorer",
        "description": "Try 10 different coffee beans",
        "imageUrl": "https://example.com/badge1.jpg",
        "earnedAt": "2023-07-20T12:00:00Z"
      },
      {
        "id": "123e4567-e89b-12d3-a456-426614174012",
        "name": "Brew Master",
        "description": "Log 50 brew sessions",
        "imageUrl": "https://example.com/badge2.jpg",
        "earnedAt": "2023-07-25T12:00:00Z"
      }
      // More badges...
    ]
  }
}
```

## File Upload Endpoints

#### POST /upload/image

Upload an image to be used in the application.

**Request:** Multipart form data with image file

**Response:**
```json
{
  "status": "success",
  "data": {
    "imageUrl": "https://storage.brewkar.com/uploads/123e4567-e89b-12d3-a456-426614174000/image.jpg"
  }
}
```

**Algorithm:**
1. Validate image file (size, type)
2. Generate unique filename
3. Upload to cloud storage (S3)
4. Return public URL

## External API Access

The Phase 3 roadmap mentions enabling API access for external coffee-related applications and smart device integrations.

### API Key Management

#### POST /api-keys

Generate a new API key.

**Request:**
```json
{
  "name": "My Coffee App Integration",
  "scopes": ["beans:read", "recipes:read", "brew-logs:read", "brew-logs:write"]
}
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "apiKey": {
      "id": "123e4567-e89b-12d3-a456-426614174013",
      "key": "bk_api_123456789abcdef",
      "name": "My Coffee App Integration",
      "scopes": ["beans:read", "recipes:read", "brew-logs:read", "brew-logs:write"],
      "createdAt": "2023-08-02T10:00:00Z",
      "lastUsed": null
    }
  }
}
```

#### GET /api-keys

List all API keys for the current user.

**Response:**
```json
{
  "status": "success",
  "data": {
    "apiKeys": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174013",
        "name": "My Coffee App Integration",
        "scopes": ["beans:read", "recipes:read", "brew-logs:read", "brew-logs:write"],
        "createdAt": "2023-08-02T10:00:00Z",
        "lastUsed": "2023-08-02T11:00:00Z"
      }
      // More API keys...
    ]
  }
}
```

#### DELETE /api-keys/:id

Revoke an API key.

**Response:**
```json
{
  "status": "success",
  "data": null
}
```

## Error Handling

All API endpoints follow a consistent error response format:

```json
{
  "status": "error",
  "error": {
    "code": "RESOURCE_NOT_FOUND",
    "message": "The requested resource was not found",
    "details": {
      "resourceType": "bean",
      "resourceId": "123e4567-e89b-12d3-a456-426614174000"
    }
  }
}
```

Common error codes:
- `INVALID_REQUEST`: The request format is invalid
- `AUTHENTICATION_REQUIRED`: Authentication is required
- `INVALID_CREDENTIALS`: Invalid email or password
- `RESOURCE_NOT_FOUND`: The requested resource was not found
- `PERMISSION_DENIED`: The user does not have permission to access the resource
- `VALIDATION_ERROR`: The request data failed validation
- `RATE_LIMIT_EXCEEDED`: The user has exceeded the rate limit for this endpoint
- `SERVER_ERROR`: An unexpected server error occurred

---

# Implementation Notes

## Caching Strategy

The API implements Redis caching for:
1. Public recipes and brew logs 
2. User profiles
3. Analytics data
4. Social feed queries

Cache invalidation occurs when:
- A user updates their profile
- A user creates or updates a recipe/brew log
- A user likes or comments on content

## Rate Limiting

All endpoints use Redis-based rate limiting:
- 60 requests per minute for authenticated endpoints
- 30 requests per minute for authentication endpoints
- 120 requests per minute for GET operations
- 10,000 requests per day per API key for external API access

## Performance Optimization

1. **Pagination**: All list endpoints support pagination
2. **Partial Responses**: Support for selecting specific fields to reduce payload size
3. **Compression**: gzip/Brotli for all responses
4. **Database Optimization**: 
   - Efficient indexing for common queries
   - Connection pooling
   - Query optimization

## Security Measures

1. **Authentication**: JWT-based authentication with refresh tokens
2. **Authorization**: Role-based access control
3. **Input Validation**: All request data is validated
4. **HTTPS**: All API traffic is encrypted
5. **CORS**: Proper configuration for web client access
6. **API Keys**: Scoped access for external integrations
7. **Rate Limiting**: Prevent abuse and DoS attacks

## Monitoring and Logging

1. **Request Logging**: Log all API requests and responses
2. **Error Tracking**: Log and monitor all errors
3. **Performance Metrics**: Track response times and throughput
4. **Usage Analytics**: Monitor endpoint usage and user activity

## Versioning Strategy

The API uses a prefix versioning strategy (`/v1`) to allow for future API changes without breaking existing clients. 