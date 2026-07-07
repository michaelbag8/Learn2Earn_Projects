# AI as Learning Amplifier

## Phase 1: Build Foundation (you first)

Wireless Routing Protocol  is a proactive routing protocol for ad hoc networks.You use WRP to decide how data moves from one node to another. In an ad hoc network, no central router exists. Each device forwards traffic. WRP tells each device which path to use.WRP keeps route information ready at all times.

Each node stores four tables.

- Distance table
- Routing table
- Link cost table
- Message retransmission list

## How WRP works in practice:
Each node exchanges routing updates with neighbors.
Each node computes the shortest path using hop count.
Each node updates its routing table continuously.

Example with 5 nodes:
```
A — B — D
 \     /
   C — E
```
## Phase 2: Strategic AI Use

# 1. Test Your Understanding
I test my understanding by answering 2 questions without AI .
### Question 1  
Why is WRP classified as proactive?

My answer:  
Wireless Routing Protocol maintains routing tables at all times. Nodes do not wait for route requests that is why they are classified as proactive routing protocol

### Question 2  
Why does WRP use four tables instead of one?

My answer:  
It separates neighbor knowledge, routing decisions, link costs, and retransmission tracking to prevent loops and improve reliability. if it uses one table it won't be reliable.

# 2. Explore Edge Cases
Pushing my  design into failure or extreme scenarios.
I ask targeted questions such as:

### What happens if two links fail at the same time?
If two links fail simultaneously, multiple nodes may mark routes as unreach destination. This increases update traffic and recalculation overhead.

### What if nodes move rapidly?
If nodes move rapidly, proactive routing increases control overhead and reduces efficiency.


# 3. Validate with AI
Using AI strategically to validate my answers

I ask AI the following questions:
- Is my explanation of link failure correct?
- Did I miss any major failure scenario?
- Is proactive routing still efficient in high mobility?
- Would reactive routing perform better in this edge case?

Validation result:  
My explanation was correct regarding link break handling. I missed update storm risk and scalability limits. AI highlighted routing overhead growth under high mobility.

# Phase 3: Real Application

Design a small smart-city network (1,000 IoT sensors, 50 traffic lights, 10 emergency vehicles). Decide protocols, justify choices, list failure points, then refine with AI feedback

![my_design](https://i.ibb.co/1GXyqC83/sc.png)

The 1,000 IoT sensors are divided into 10 zones to make the system easier to manage and scale. Instead of all sensors sending data directly to the control center, they form small wireless mesh groups. Inside each zone, sensors pass messages from one to another until the data reaches a gateway. The gateways act like collection hubs; they gather sensor data and send it securely over high-speed fiber or 5G to the central control center. 

The 50 traffic lights are organized into clusters, for example, 5 per cluster. Each cluster has a local controller that manages signal timing in its area. These controllers connect to the central traffic management system for city-wide coordination. 

The 10 emergency vehicles connect through a secure 5G network using encrypted tunnels back to the control center. When an ambulance or fire truck is dispatched, it sends its GPS location to the system. The control center calculates the best route and sends commands to traffic light controllers
the protocol is proactive , because traffic flows frequently between many nodes, the failure point is frequent updates under high mobility


# AI Model : Refined Design and Feedback

Your structure is strong.  
You separated sensors, traffic lights, and vehicles.  
You introduced zones and clustering.  
You reduced direct load on the control center.  

Now I refine your design and give focused feedback.  

# Refined Smart City Network Design  

## 1. IoT Sensors, 1,000 units  

### Design  

- Divide sensors into 10 zones, 100 sensors per zone.  
- Use a wireless mesh inside each zone.  
- Use Wireless Routing Protocol inside each zone.  
- Each zone has one gateway.  
- Gateways connect to control center via fiber or 5G.  

### Justification  

- Mesh reduces single point dependency at sensor level.  
- Zoning limits routing table size.  
- Proactive routing reduces delay for frequent reporting such as air quality, traffic density, flooding alerts.  
- Gateway aggregation reduces backbone congestion.  

### Failure points  

- Gateway failure isolates entire zone.  
- High routing overhead if sensors move or fail frequently.  
- Power failure in battery sensors.  
- Interference in dense wireless areas.  

### Improvement  

- Add backup gateway per zone.  
- Add energy monitoring for sensors.  
- Use redundant backhaul link.  

---

## 2. Traffic Lights, 50 units  

### Design  

- Divide into 10 clusters, 5 lights per cluster.  
- Each cluster has a local controller.  
- Use wired Ethernet or fiber between controller and lights.  
- Controllers connect to central system using IP routing.  

### Justification  

- Wired link increases reliability.  
- Local controller allows intersection-level autonomy.  
- Reduces latency during signal switching.  
- Central system handles optimization and coordination.  

### Failure points  

- Cluster controller failure disables 5 intersections.  
- Fiber cut disconnects cluster from city system.  
- Software bug in timing logic.  

### Improvement  

- Install fail-safe mode in traffic lights.  
- Allow default timing plan if controller loses connection.  
- Add redundant controller for critical intersections.  

---

## 3. Emergency Vehicles, 10 units  

### Design  

- Use 5G network connection.  
- Establish secure VPN tunnel to control center.  
- Send GPS every few seconds.  
- Control center pushes priority signal requests to traffic controllers.  

### Justification  

- High mobility requires wide-area cellular coverage.  
- 5G supports low latency communication.  
- Centralized routing ensures optimal city-wide coordination.  

### Failure points  

- 5G network outage.  
- GPS signal loss in tunnels.  
- Cyberattack on communication channel.  

### Improvement  

- Add fallback LTE channel.  
- Add offline routing map inside vehicle.  
- Use strong authentication and encryption.  

---

# Protocol Decision Feedback  

You selected proactive routing.  

### Refinement  

- Use proactive protocol such as Wireless Routing Protocol only inside sensor zones.  
- Do not use proactive protocol city-wide.  
- Use standard IP routing in backbone network.  
- Use cellular routing for vehicles.  

### Why not city-wide proactive  

- 1,000 nodes city-wide increases routing table size.  
- Control overhead becomes large.  
- Mobility from vehicles increases update frequency.  

---

# Expanded Failure Analysis  

Your original failure point was incomplete.  

You listed:  

- Frequent updates under high mobility.  

Additional failure risks:  

- Routing overhead explosion in dense mesh.  
- Gateway single point failure.  
- Backhaul congestion during peak hours.  
- Cybersecurity breaches.  
- Power outages.  
- Software synchronization errors.  

---

# Overall Feedback  

## Strengths  

- Clear zoning strategy.  
- Logical separation of systems.  
- Sensible use of mesh networking.  

## Areas to improve  

- Add redundancy planning.  
- Differentiate protocols per subsystem.  
- Expand failure analysis beyond mobility.  
- Consider cybersecurity explicitly.  

Your design shows good systems thinking.  
With redundancy and protocol layering refinement, it becomes strong for an academic submission. 



# Reflection:
### % human judgment vs. AI contribution
I did about 75 % manually on my own before using AI to validate my work and find issues and gave me feedback

### Could you defend decisions without AI?

Yes, i could defend my decisions without AI

### What will you still remember in 6 months?
I will still remember what wireless routing protocol is , how it works and what are the failure points

### Did AI make you sharper, or think for you?

AI exposed blind spots, and  I did not outsource my thinking to it because the core reasoning was mine
