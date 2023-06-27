# Tiny-SSP-with-Open-RTB-spec

## Description
This project implements a tiny Supply-Side Platform (SSP) conforming to the OpenRTB 2.5 spec. It simulates the process of bidding for ad placements in an automated auction. 

The main goal of the project is to address two important considerations in SSP operation: improving Queries Per Second (QPS) and effectively designing the system landscape for servers in a cloud environment.

## System Architecture
The system comprises an SSP server and two Demand-Side Platform (DSP) servers (dsp1 and dsp2). 

### SSP Server
The SSP server carries out the following functions:

1. Receive ad call from a web browser
2. Send bid request to the DSP partners
3. Receive bid responses from DSP partners
4. Run an auction to determine the highest bid
5. Return the winning ad to the browser
6. Trigger an impression pixel upon ad display
7. Store the impression data in MongoDB via a log server

### DSP Servers
Each DSP server carries out the following functions:

1. Receive bid request from the SSP server
2. Respond with a bid response, including the bid price and ad creatives
3. Trigger an impression pixel upon ad display

## Improving QPS
Queries Per Second (QPS) is a key performance metric in digital advertising. In this project, we improve QPS by optimizing both the SSP server and the DSP servers for high concurrency and low latency. We use lightweight protocols, async I/O operations, and efficient data structures to ensure that we can process a high number of queries per second. We also leverage cloud auto-scaling capabilities to handle traffic spikes.

## System Landscape in a Cloud Environment
The servers in this system (SSP, dsp1, dsp2, logserver) are designed to run independently and can be scaled horizontally. This architecture supports high availability and fault tolerance. Load balancers distribute incoming ad requests across multiple SSP server instances, and the SSP server distributes bid requests across multiple DSP servers. This ensures that the system can handle high QPS while maintaining low latency.

## Usage
To use this project, you can use Docker Compose to build and run the system with the following commands:

```bash
docker-compose build --no-cache
docker-compose up
```

## Testing
Each component (ssp, dsp1, dsp2 and logserver) includes unit test cases to ensure that it functions as expected.
```bash
go test ./...
```

## References
- [OpenRTB 2.5 spec](https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf)
- [Prebid's OpenRTB GitHub repo](https://github.com/prebid/openrtb)

## Contributions
Contributions to this project are welcome. Please follow the standard GitHub Pull Request process to propose changes to this repository.
