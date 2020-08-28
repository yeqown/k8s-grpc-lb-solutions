# k8s-grpc-lb-solutions
LoadBalancer solutions in kubernetes, includes demo code and solution detail


### Solutions

* [gRPC Client LoadBalancing](./solutions/lb-client/README.md)
    
    * advantages：
        
        * No need to import `proxy`, no need to maintain an external component.
        
        * No delay to increase.
    
    * disadvantages:
        
        * `Client` changes and `LoadBalancing` strategy is simple as usually.
        
        * `Service` changes, [headless](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services) service required.

* [gRPC Server LoadBalancing proxy](./solutions/lb-proxy-mesh/README.md)
    
    * advantages：
        
        * Keep client simple and no changes. 
    
    * disadvantages:
        
        * Increasing delay.
        
        * Need to maintain `LoadBalancing` proxy 

### References

* https://www.youtube.com/watch?v=F2znfxn_5Hg
* https://linkerd.io/2/overview/
* http://yangxikun.github.io/golang/2019/10/19/golang-grpc-client-side-lb.html
* https://github.com/grpc/grpc/blob/master/doc/load-balancing.md