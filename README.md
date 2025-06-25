# Synapse

![CI/CD Status](https://img.shields.io/github/actions/workflow/status/YOUR_USERNAME/synapse/go.yml?branch=main&style=for-the-badge)
![Go Report Card](https://goreportcard.com/badge/github.com/YOUR_USERNAME/synapse?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/YOUR_USERNAME/synapse?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/YOUR_USERNAME/synapse?style=for-the-badge)

A flexible, self-hostable webhook-to-notification relay. Synapse receives webhooks from any source and intelligently forwards them to your configured channels like Telegram, Slack, and more.

---

## 🤔 What is Synapse?

In modern development, we rely on countless services that communicate via webhooks (e.g., Git pushes, user sign-ups, monitoring alerts). The problem? Each service has its own JSON format. Synapse acts as a powerful middleware that catches these webhooks, allows you to extract the data you need, format it into a human-readable message, and route it to one or more notification channels.

It's your personal, centralized notification hub.

## ✨ Key Features

* **Project-Based Webhooks**: Create isolated "projects" for each message source, each with its own unique webhook URL.
* **Multi-Channel Support**: Forward messages to multiple channels simultaneously. Easily extensible to support new services (Telegram, Slack, Webhooks, etc.).
* **Advanced Message Templating**: Don't just forward ugly JSON! Use Go's `text/template` syntax to create beautiful, custom message formats.
* **Dynamic Variable Extraction**: Pull data from any part of the incoming JSON payload using `gjson` path syntax. No need to write custom code for every webhook format.
* **Flexible Routing Strategies**:
    * **Send to All**: Broadcast a message to all configured channels.
    * **Failover**: Send messages in a prioritized order, stopping as soon as one channel succeeds. Perfect for ensuring deliverability.
* **Configurable Execution Modes**:
    * **Asynchronous (Default)**: Responds instantly to the webhook source and processes the message in the background for maximum performance.
    * **Synchronous**: Waits for the message to be sent and returns the final delivery result to the caller.
* **Message History & Logs**: A complete history of received messages and their delivery status for easy debugging and auditing.
* **Modern Web UI**: A clean and intuitive interface built with Vue.js and Naive UI to manage your projects, channels, and routes.

##  diagrama de arquitectura

```mermaid
graph TD
    A[External Service<br/>(e.g., GitHub, Stripe)] -- JSON Payload --> B{Synapse Webhook Endpoint<br/>(Gin)};
    B -- 1. Save Message (Pending) --> C[MySQL Database];
    B -- 2. Push Message ID --> D{Async Queue<br/>(Go Channel)};
    E[Dispatcher Worker] -- 3. Consume from Queue --> D;
    E -- 4. Get Rules & Message --> C;
    E -- 5. Render Template & Send --> F[Notifier];
    F -- Telegram --> G[Telegram API];
    F -- Slack, etc. --> H[...];
    
    U[User] <--> X[Web UI<br/>(Vue.js + Naive UI)];
    X <--> Y{Synapse API<br/>(Gin)};
    Y <--> C;
```

## 🛠️ Tech Stack

  * **Backend**: Go, Gin, GORM, Viper, JWT
  * **Frontend**: Vue.js (v3), Pinia, Vue Router, Naive UI, Axios
  * **Database**: MySQL
  * **Containerization**: Docker, Docker Compose

## 🚀 Getting Started

Follow these steps to get your own Synapse instance up and running locally.

### Prerequisites

  * [Git](https://git-scm.com/)
  * [Go](https://go.dev/) (v1.21+)
  * [Node.js](https://nodejs.org/) (v18+)
  * [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

### 1\. Installation

Clone the repository:

```bash
git clone [https://github.com/YOUR_USERNAME/synapse.git](https://github.com/YOUR_USERNAME/synapse.git)
cd synapse
```

### 2\. Configuration

Copy the example configuration file.

```bash
cp config.yaml.example config.yaml
```

Now, edit `config.yaml` and fill in your details, especially the `database` connection settings and a strong `jwt_secret`.

### 3\. Database Setup

Ensure your MySQL server is running. Then, connect to it and execute the SQL script to create all necessary tables.

```bash
# Using a MySQL client, execute the contents of this file:
# db/schema.sql
```

### 4\. Running the Application

The easiest way to run the entire stack (backend, frontend, and database) is with Docker Compose.

```bash
docker-compose up -d --build
```

That's it\!

  * The frontend will be available at `http://localhost:5173`
  * The backend API will be available at `http://localhost:8080`

### For Local Development (without Docker)

If you prefer to run the services directly on your host machine:

```bash
# Terminal 1: Run the Backend
cd backend # Assuming your Go code is in a backend folder
go run cmd/server/main.go

# Terminal 2: Run the Frontend
cd frontend # Assuming your Vue code is in a frontend folder
npm install
npm run dev
```

## kullanım

1.  Navigate to `http://localhost:5173`.

2.  **Register** a new user and log in.

3.  Go to the **Channels** page and add a new channel (e.g., your Telegram Bot Token and Chat ID).

4.  Go to the **Projects** page and create a new project. A unique webhook URL will be generated for you.

5.  On the project's management page, create a **Routing Rule** to link your project to the channel you just created.

6.  Use a tool like `curl` or Postman to send a test payload to your webhook URL:

    ```bash
    curl -X POST \
      -H "Content-Type: application/json" \
      -d '{ "user": "John Doe", "action": "signed_up", "level": "critical" }' \
      http://localhost:8080/webhook/YOUR_UNIQUE_WEBHOOK_KEY
    ```

7.  Check your notification channel. You should see the message\! To customize it, edit your routing rule and add a message template.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome\! Feel free to check the [issues page](https://www.google.com/search?q=https://github.com/YOUR_USERNAME/synapse/issues).

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](https://www.google.com/search?q=LICENSE) file for details.