# Infrastructure Setup Documentation

This repository contains the infrastructure setup for a microservices-based application. Each folder is dedicated to a specific component of the architecture. Below is an overview of the repository structure and the roles of each component.

## Repository Structure

### 1. **api-gateway-transformer**
   This folder contains the API Gateway responsible for routing, load balancing, and request transformations. It serves as the entry point for client requests and manages communication with backend services.

   **Key Features:**
   - Centralized routing for all microservices
   - Authentication and authorization handling
   - Request transformation and validation

   **Files:**
   - `Dockerfile`: Defines the container for the API Gateway.
   - `gateway-config.yaml`: Configuration file for routing and middleware.
   - `main.go`: Entry point for the Go-based API Gateway (if applicable).

---

### 2. **backend-optimus**
   This folder contains the backend services (microservices) responsible for the application's core logic and business operations. Each service is modular and serves a specific function.

   **Subfolders:**
   - `auth-service/`: Manages user authentication and token generation (e.g., JWT).
   - `profile-service/`: Handles user profile data.
   - `video-service/`: Manages video streaming and buffering.

   **Files:**
   - `Dockerfile`: Container setup for each service.
   - `*.yaml`: Kubernetes deployment configurations for each service.
   - `.env`: Environment variables for service configurations.

---

### 3. **ci-cd-spiderman**
   This folder contains CI/CD pipeline configurations for automating build, test, and deployment processes.

   **Key Features:**
   - Automated container builds
   - Continuous integration and testing
   - Deployment to Kubernetes clusters

   **Files:**
   - `github-actions/`: Contains GitHub Actions workflows for CI/CD pipelines.
     - `deploy.yml`: Pipeline for building and deploying containers to Kubernetes.
     - `test.yml`: Pipeline for running unit and integration tests.
   - `docker-compose.yml`: Configuration for local testing and multi-container setup.

---

### 4. **frontend-ironman**
   This folder contains the front-end application responsible for the user interface. Built using modern frameworks like React or Angular, it interacts with the backend via the API Gateway.

   **Key Features:**
   - Responsive UI for web users
   - Integration with backend APIs
   - State management and component reusability

   **Files:**
   - `Dockerfile`: Container setup for the front-end application.
   - `package.json`: Manages front-end dependencies and scripts.
   - `src/`: Source code for the application.

---

## Infrastructure Deployment Steps

1. **Set Up Docker:**
   - Ensure Docker is installed and running on your system.
   - Build and run the containers for all services:
     ```bash
     docker-compose up --build
     ```

2. **Deploy to Kubernetes:**
   - Ensure `kubectl` and Kubernetes are configured.
   - Apply the deployment configurations:
     ```bash
     kubectl apply -f backend-optimus/auth-service/auth-service.yaml
     kubectl apply -f backend-optimus/profile-service/profile-service.yaml
     kubectl apply -f backend-optimus/video-service/video-service.yaml
     kubectl apply -f api-gateway-transformer/gateway-config.yaml
     ```

3. **CI/CD Pipeline:**
   - Ensure GitHub Actions are configured with the necessary secrets (e.g., Kubernetes credentials, Docker registry access).
   - Push your changes to the repository to trigger the pipeline.

4. **Access the Application:**
   - The front-end application will be accessible at the API Gateway’s exposed IP or domain.
   - Example: `http://<your-gateway-ip>:<port>`

---

## Notes

- **Environment Variables:** Ensure all `.env` files are populated with the correct values before deploying.
- **Scaling:** Use Kubernetes Horizontal Pod Autoscaler (HPA) for auto-scaling services based on load.
- **Monitoring:** Integrate monitoring tools like Prometheus and Grafana for better observability.

---

## Future Enhancements

- Add Terraform scripts for Infrastructure as Code (IaC).
- Integrate a service mesh like Istio for advanced traffic management.
- Implement caching mechanisms (e.g., Redis) for improved performance.

---

For any issues or questions, please contact the project maintainer.

