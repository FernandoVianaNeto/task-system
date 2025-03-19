Task System API 📝

    API for task management, allowing task creation, listing, updating, and deletion, along with JWT authentication.

📌 Technologies Used

    Golang (Gin Framework) – Backend structure

    GORM – ORM for MySQL

    JWT – Secure authentication

    Kafka – Asynchronous communication for events

    Docker – Environment containerization

    MockGen – Mocking for unit tests

    GoMock – Testing framework

📌 Installation and Configuration

    1 – Configure environment variables

Create a .env file in the project root with the following configurations:

    APP_PORT=3000
    MYSQL_USER=user
    MYSQL_PASSWORD=password
    MYSQL_HOST=mysql_go
    MYSQL_PORT=3306
    MYSQL_NAME=mydatabase
    PASSWORD_SECRET_HASH="SECRET_HASH"
    TASK_STATUS_UPDATED_TOPIC="task-status-updated.events"
    KAFKA_BROKERS_HOST=kafka

2️⃣ Start the environment with Docker

    First start the services:

    make docker/up

    The command above build and run the database container, http api container and task consumer container, with detached mode.

    The Api is now available in http://locahost:3000 with de following endpoints

📌 Endpoints

Authentication

🔹 Login'
    POST /auth

    Body:

    {
        "email": "user@email.com",
        "password": "password123"
    }

    Response:

    {
        "token": "eyJhbGciOiJIUzI1NiIsIn..."
    }

    OBS: Inside the token are 3 additional fields, containing informations about the user: "user_email", "user_role" and "user_uuid". These informations are used to validate, for example, which tasks can be show or if can delete or update some task.

Users

    🔹 Create user

    POST /user

    Body:

    {
        "name": "John Doe",
        "email": "john@email.com",
        "password": "password123",
        "role": "admin"
    }

Tasks

    🔹 Create task

    POST /task
    Authorization: Bearer {token}

    Body:

    {
        "title": "Review code",
        "summary": "Review PR #45 on GitHub"
    }

    🔹 List tasks

    GET /task
        Authorization: Bearer {token}

    🔹 Update task status

    PUT /task
        Authorization: Bearer {token}

    Body:

        {
            "task_uuid": "123e4567-e89b-12d3-a456-426614174000",
            "new_status": "completed"
        }

    OBS: Everytime the task has its status changed, an event will be produce, at the same time, a consumer will receive this event and log a message in the terminal with the task uuid, the user uuid and the time that it was performed

    🔹 Delete task

        DELETE /task/{uuid}
        Authorization: Bearer {token}

📌 Testing

1️⃣ Run unit tests

    make test

2️⃣ Generate mocks

    mockgen -source=internal/domain/repository/task.repository.go -package=domain_repository -destination=test/mocks/domain_repository/mock_task.repository.go

📌 Messaging (Kafka)

    The API publishes task update events to the Kafka topic task-status-updated.

📌 Published event

    {
    "user_uuid": "123e4567-e89b-12d3-a456-426614174000",
    "task_uuid": "321e4567-e89b-12d3-a456-426614174000",
    "new_status": "completed",
    "timestamp": "2025-03-15T12:34:56Z"
    }

    To test Kafka messages:

    kafkacat -b localhost:9092 -t task-status-updated -C