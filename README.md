## Overview

NimbusStore is a Go-based project designed to implement a custom key-value store. This project aims to explore and demonstrate fundamental as well as advanced distributed systems concepts through the lens of a practical, real-world application. Inspired by principles from the Raft consensus algorithm and other distributed systems techniques, NimbusStore includes several key features:

- **Basic CRUD Operations:** Supports Create, Read, Update, and Delete operations for efficient data storage and retrieval.
- **Replication:** Implements data replication across multiple nodes to ensure high availability and fault tolerance, a crucial aspect for distributed systems.
- **Consensus Algorithm (Inspired by Raft):** Utilizes concepts inspired by the Raft consensus algorithm to manage a replicated log, ensuring strong consistency and reliability in a cluster environment.
- **Sharding:** Incorporates sharding techniques to effectively distribute data across multiple nodes, enhancing scalability and load balancing.
- **Fault Tolerance:** Designed to handle node failures gracefully, NimbusStore maintains data integrity and availability, showcasing robust fault tolerance capabilities.
- **Scalability:** Engineered to scale horizontally, facilitating the addition of more nodes to the system seamlessly without significant downtime or performance degradation.
- **Simple Client API:** Offers a straightforward and user-friendly API for clients, abstracting the complexities of the underlying distributed system operations.

NimbusStore serves as both a robust storage solution and a learning project that encapsulates the challenges and solutions inherent in the design of distributed systems.
