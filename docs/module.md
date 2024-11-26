# API Documentation

## Overview

This document provides an overview of the API endpoints available in the application.

## Endpoints

### 1. Get All Items

- **URL:** `/api/items`
- **Method:** `GET`
- **Description:** Retrieves a list of all items.
- **Response:**
    - `200 OK`: Returns a JSON array of items.

### 2. Get Item by ID

- **URL:** `/api/items/{id}`
- **Method:** `GET`
- **Description:** Retrieves a single item by its ID.
- **Parameters:**
    - `id` (required): The ID of the item.
- **Response:**
    - `200 OK`: Returns the item details.
    - `404 Not Found`: Item not found.

### 3. Create New Item

- **URL:** `/api/items`
- **Method:** `POST`
- **Description:** Creates a new item.
- **Request Body:**
    - `name` (required): The name of the item.
    - `description` (optional): The description of the item.
- **Response:**
    - `201 Created`: Item successfully created.
    - `400 Bad Request`: Invalid input.

### 4. Update Item

- **URL:** `/api/items/{id}`
- **Method:** `PUT`
- **Description:** Updates an existing item.
- **Parameters:**
    - `id` (required): The ID of the item.
- **Request Body:**
    - `name` (optional): The name of the item.
    - `description` (optional): The description of the item.
- **Response:**
    - `200 OK`: Item successfully updated.
    - `400 Bad Request`: Invalid input.
    - `404 Not Found`: Item not found.

### 5. Delete Item

- **URL:** `/api/items/{id}`
- **Method:** `DELETE`
- **Description:** Deletes an item by its ID.
- **Parameters:**
    - `id` (required): The ID of the item.
- **Response:**
    - `200 OK`: Item successfully deleted.
    - `404 Not Found`: Item not found.

## Error Codes

- `200 OK`: The request was successful.
- `201 Created`: The resource was successfully created.
- `400 Bad Request`: The request was invalid or cannot be served.
- `404 Not Found`: The requested resource could not be found.

## Example Requests

### Get All Items

```bash
curl -X GET http://localhost:3000/api/items
```

### Get Item by ID

```bash
curl -X GET http://localhost:3000/api/items/1
```

### Create New Item

```bash
curl -X POST http://localhost:3000/api/items -d '{"name": "NewItem", "description": "New item description"}' -H "Content-Type: application/json"
```

### Update Item

```bash
curl -X PUT http://localhost:3000/api/items/1 -d '{"name": "UpdatedItem", "description": "Updated item description"}' -H "Content-Type: application/json"
```

### Delete Item

```bash
curl -X DELETE http://localhost:3000/api/items/1
```
