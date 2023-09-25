# SaaS Summary

### 1. What's the different between IaaS, SaaS and PaaS

1. IaaS 

IaaS Stands for "Infrastrusture as a Service." This refers to the delivery of computing resources, such as servers and storage, over the internet. IaaS make those resources available to all customers on-demand, allowing them to scale and adjust their usage as needed.

2. SaaS

SaaS or "Software as a Service." Is refers to the delivery of software applications over the internet typically on subcription basis. Saas provide application that customers can accsess and use the software platform from anywhere, without needing to install it on their devices.

3. PaaS

PaaS is "platform as A service." this refers to the delivery of a platform for developing, runing, and managing applications over the internet.

#### Conclusion
1. Iaas provide the infrastructure.
2. SaaS provides the software,
3. PaaS provides the platform for building, deploying dan managing applications


### 2. Saas Platform Achitecture

Software as a Service (SaaS) is a software licensing and delivery model that provides software on a subscription basis and hosts it centrally, allowing users to access it through web browsers. SaaS is commonly used for various business applications, including office, messaging, management software, and virtualization, and is a part of cloud computing, alongside Infrastructure as a Service (IaaS), Platform as a Service (PaaS), and Desktop as a Service (DaaS).

SaaS is associated with Application Service Providers (ASPs) who offer web-based applications to business users. Unlike early internet-delivered software, modern SaaS applications are built on a single-instance, multi-tenant architecture, offering feature-rich experiences comparable to on-premise applications. Aggregators bundle SaaS offerings from different vendors into unified application platforms.

The SaaS provider hosts the application centrally, deploying patches and upgrades transparently over the internet. Many SaaS vendors provide APIs for developers to create composite applications, and they implement various security mechanisms to ensure data security during transmission and storage.

SaaS architecture typically involves a single application version and configuration for all customers, with scalability achieved through horizontal scaling. While some exceptions exist, most SaaS solutions employ multitenancy to manage a large customer base cost-effectively.

There are two main types of SaaS: Vertical SaaS, tailored to specific industries, and Horizontal SaaS, which focuses on software categories applicable to various industries.

Benefits of SaaS include reduced deployment costs, minimized IT infrastructure requirements, ease of integration, and shorter time-to-value intervals. It allows organizations of all sizes to shift from reactive IT cost centers to proactive, value-producing entities.

Transitioning to SaaS requires assessing its impact on existing IT assets, addressing data security standards, and evaluating reporting services. SaaS can change IT roles and responsibilities, potentially causing a shift from gatekeeper to a more collaborative role.

#### conclusion

Enterprises should consider the flexibility and risk-management implications of adding SaaS to their IT services portfolio. Successful integration and composition are critical for incorporating SaaS effectively into a service-centric IT infrastructure, and it is believed that the future of enterprise computing will include a combination of on-premise and cloud-based solutions.

### 3. SaaS (Software as a Service) Platform Architecture

Saas is way to deliver software, the provider of the software centrally hosts one or more applications an makes them available to customers over the internet. Saas architecture ia also on of the main pillars of cloud computing.

#### The Simplicity of SaaS Arcitecture Apps

Software applications architected as SaaS solution are typically accesed over the web trough various types of devices.

#### Lack of control in SaaS Architecture platforms 

An in-house or on-premise application will give an bussiness more control over it's behavior, for example, a Windows-based application might have more configuration otion than regular web.

#### Limited ecosystem

SaaS is growing trend as a software distribution channel. There are still many application that don't offer a hosted version.

#### Data Concerns

When selecting a SaaS product, and for example, with the advent of the GDPR, bussines must pay special attention in terms of where any SaaS implementation stores data in the cloud.

#### Performence 

An in-house, thick client, or on-premiese application will always run quicker than a product being delivered over the internet.

#### Key Components of a SaaS platform

Our Latest SaaS is to determine which key components or features users will expect as satndard. Features will understandaly by driven by market and user community demands.


#### Security

Protecting customer data in our SaaS platform is of the utmost importance, as such, our SaaS product will most likely serve hundreds, if not thousands of users.  Ensure our SaaS architecture takes this into account.

#### Privacy

With new regulations, such as GDPR, business is more accuntable than ever for ensuring user and data privacy is maintained.


#### Customization and Configuration

Factoring in extensibility, SaaS architecture is another important component for our consider. We can shipping a "White label" version of our SaaS product or by implementing plugin mechanism.

### Design Considerations for a SaaS Architecture Platform

#### Scalability

SaaS architectur must be able to scale and accomodate hundereds.

we can archieve by using Network Load Balancers to evenly distribute or optimize incoming traffic acoss multiple web servers.

#### Zero downtime and Service Level Agreements

With internal or on-premise software applications, users are more forgiving on the occasion that software must be taken offline for a given period of time whilst internal IT install a new kit or release a version update.

Users typically expect the product to be available almost all the time and with no interruption to service.

Take time to consider how we will factor upgrades, patches, or debugging and troubleshooting production issues into our SaaS applications architecture.

#### Multi-tenancy in SaaS architecture

SaaS product must have support multi-tenancy. Our product needs to be able to accommodate multiple users whilst at the same time, ensuring that user data, privacy and are all still being observed

### 4. How to build a cloud-based SaaS Application

### Choose Programing Language

First for all, we have to choose our programing language that we wanna use. In this case we use Phyton, because Python is great and many developers love it. Dynamic typing, meta programming, rapid prototyping. Everything’s possible with Python.

### Choose Database 

In this case the recommendations is making use of a document-oriented database. 

Document databases get their type information from the data itself. Thus every instance of data can be different from any other.

This allows more flexibility, especially when dealing with changes. And it often reduces database sizes.

In summary, the DOB concept offers a richer experience with modern programming techniques.

MongoDB is a document-oriented database that provides high performance, high availability, and easy scalability.

Besides performance, scalability is the most important factor for a global SaaS business.


### Set up MongoDB for the SaaS application

In this case is using AWS cloud and therefore have set up EC2 instances in others country such as America etc.

### Queuing system for your SaaS application


A message queuing system is an asynchronous communication protocol, enabling sender and receiver of a message not interacting at the same time.

Also known as Message Queuing (MSMQ) technology it enables web apps to run at different times and to communicate with various 3rd party integrations / APIs / and other services asynchronously.

A message (e.g. a query asking a 3rd party service via an API) is placed onto the queue. It’s stored there until the receiver retrieves it.

A message queue has limits regarding the size and amount of data transmitted in the queue. The great thing about modern queuing systems is that they can be scaled easily.

### Web Storage S3

Getting more user on board for our product will make easily wonder about our web storage. With the Amazon S3 storage service, we have a great, and highly scalable object storage installed.


### Content Delivery Network for your SaaS application

A content delivery network (CDN) is basically a system of distributed servers which enables, this is to serve content to our app users with high performance and high availability.


### Recap: SaaS application set up

With Python, MongoDB – as a great document-orientated database, RabbitMQ software-wise the basic setup is done. However, there is way more to think of.

### Final step start with software testing




