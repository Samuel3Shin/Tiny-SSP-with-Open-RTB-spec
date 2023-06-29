# Tiny-SSP-with-Open-RTB-spec

## Description
This project implements a tiny Supply-Side Platform (SSP) conforming to the OpenRTB 2.5 spec. It simulates the process of bidding for ad placements in an automated auction with [react-web-for-tiny-ssp repo](https://github.com/Samuel3Shin/react-web-for-tiny-ssp).

## System Architecture
The system comprises an SSP server, two Demand-Side Platform (DSP) servers (dsp1 and dsp2) and Log server to log impression pixel. 

### SSP Server
The SSP server carries out the following functions:

1. Receive ad call from a web browser
2. If the ad call is in a cache, return the cached response.
3. If the cache misses, send a bid request to the DSP partners
4. Receive bid responses from DSP partners
5. Run an auction to determine the highest bid
6. Return the winning ad to the browser
7. Trigger an impression pixel upon ad display
8. Store the impression data in MongoDB via a log server

### DSP Servers
Each DSP server carries out the following functions:

1. Receive bid request from the SSP server
2. Respond with a bid response, including the bid price and ad creatives
3. Trigger an impression pixel upon ad display

### Log server
The Log server carries out the following functions:

1. Web browser triggers an impression pixel upon ad display
2. Store the impression data(Ad ID, Timestamp) in MongoDB

## Improving QPS
Queries Per Second (QPS) is a key performance metric in digital advertising. In this project, I used goroutines to handle multiple requests concurrently. I also used cache to reduce latency for the same ad requests from web browsers. We can also leverage cloud auto-scaling capabilities to handle traffic spikes in the future.

## System Landscape in a Cloud Environment
The servers in this system (SSP, dsp1, dsp2, logserver) are designed to run independently (stateless) so they can be scaled horizontally. This architecture supports high availability and fault tolerance. If we assume to use AWS in the future, load balancers will distribute incoming ad requests across multiple SSP server instances, and the SSP server distributes bid requests across multiple DSP servers. This ensures that the system can handle high QPS while maintaining low latency.

## Usage
To use this project, you can use Docker Compose to build and run the system with the following commands from the root directory of this project:

Make sure you have installed Docker in your machine.

```bash
docker-compose build --no-cache
docker-compose up
```

## Testing
Each component (ssp, dsp1, dsp2 and logserver) includes unit test cases to ensure that it functions as expected. To test all the components, run the following commands from the root directory of this project.
```bash
go test ./...
```

## References
- [OpenRTB 2.5 spec](https://www.iab.com/wp-content/uploads/2016/03/OpenRTB-API-Specification-Version-2-5-FINAL.pdf)
- [Prebid's OpenRTB GitHub repo](https://github.com/prebid/openrtb)

## Contributions
Contributions to this project are welcome. Please follow the standard GitHub Pull Request process to propose changes to this repository.
