<p align="center">
 <img src="https://user-images.githubusercontent.com/796136/50286124-6f7f3780-046f-11e9-9f45-e8fedd4f786d.png" height="75px" alt="RoadRunner">
</p>
<p align="center">
 <a href="https://packagist.org/packages/spiral/roadrunner"><img src="https://poser.pugx.org/spiral/roadrunner/version"></a>
	<a href="https://godoc.org/github.com/agui2200/roadrunner"><img src="https://godoc.org/github.com/agui2200/roadrunner?status.svg"></a>
	<a href="https://travis-ci.org/spiral/roadrunner"><img src="https://travis-ci.org/spiral/roadrunner.svg?branch=master"></a>
	<a href="https://goreportcard.com/report/github.com/agui2200/roadrunner"><img src="https://goreportcard.com/badge/github.com/agui2200/roadrunner"></a>
	<a href="https://scrutinizer-ci.com/g/spiral/roadrunner/?branch=master"><img src="https://scrutinizer-ci.com/g/spiral/roadrunner/badges/quality-score.png"></a>
	<a href="https://codecov.io/gh/spiral/roadrunner/"><img src="https://codecov.io/gh/spiral/roadrunner/branch/master/graph/badge.svg"></a>
</p>

RoadRunner is an open source (MIT licensed) high-performance PHP application server, load balancer and process manager.
It supports running as a service with the ability to extend its functionality on a per-project basis. 

RoadRunner includes PSR-7/PSR-17 compatible HTTP and HTTP/2 server and can be used to replace classic Nginx+FPM setup with much greater performance and flexibility.

Table of Contents 
-----------------
* Introduction
  * [About RoadRunner](https://github.com/agui2200/roadrunner/wiki/About-RoadRunner)
  * [Installation](https://github.com/agui2200/roadrunner/wiki/Installation)
  * [Quick Builds](https://github.com/agui2200/roadrunner/wiki/Quick-Builds)
  * [Configuration](https://github.com/agui2200/roadrunner/wiki/Configuration)
  * [License](https://github.com/agui2200/roadrunner/wiki/License)
* Using RoadRunner
  * [Environment Configuration](https://github.com/agui2200/roadrunner/wiki/Enviroment-Configuration)
  * [HTTPS and HTTP/2](https://github.com/agui2200/roadrunner/wiki/HTTPS-and-HTTP2)
  * [**PHP Workers**](https://github.com/agui2200/roadrunner/wiki/PHP-Workers)
  * [Debugging](https://github.com/agui2200/roadrunner/wiki/Debugging)
  * [Server Commands](https://github.com/agui2200/roadrunner/wiki/Server-Commands)
  * [RPC Integration](https://github.com/agui2200/roadrunner/wiki/RPC-Integration)
  * [Restarting Workers](https://github.com/agui2200/roadrunner/wiki/Restarting-Workers)
  * [IDE integration](https://github.com/agui2200/roadrunner/wiki/IDE-Integration)
  * [Error Handling](https://github.com/agui2200/roadrunner/wiki/Debug-And-Error-Handling)
  * [Production Usage](https://github.com/agui2200/roadrunner/wiki/Production-Usage)
* Integrations
   * [Laravel Framework](https://github.com/agui2200/roadrunner/wiki/Laravel-Framework) (alpha)
   * [Slim Framework](https://github.com/agui2200/roadrunner/issues/62) (pending verification)
   * [Symfony Framework](https://github.com/agui2200/roadrunner/wiki/Symfony-Framework)
   * [Yii2/3 Framework](https://github.com/agui2200/roadrunner/issues/78) (in progress)
* Server Customization
  * [Building Server](https://github.com/agui2200/roadrunner/wiki/Building-Server)
  * [Writing Services](https://github.com/agui2200/roadrunner/wiki/Writing-Services)
  * [HTTP Middlewares](https://github.com/agui2200/roadrunner/wiki/Middlewares)
* Additional Notes
  * [Event Listeners](https://github.com/agui2200/roadrunner/wiki/Event-Listeners)
  * [Standalone Usage](https://github.com/agui2200/roadrunner/wiki/Standalone-usage)
  * [AWS Lambda](https://github.com/agui2200/roadrunner/wiki/AWS-Lambda)
* Custom Builds
  * [GRPC Server](https://github.com/agui2200/php-grpc)

Features:
--------
- production ready
- PSR-7 HTTP server (file uploads, error handling, static files, hot reload, middlewares, event listeners)
- HTTPS and HTTP/2 support (including HTTP/2 Push)
- fully customizable server
- flexible environment configuration
- no external PHP dependencies, drop-in (based on [Goridge](https://github.com/spiral/goridge))
- load balancer, process manager and task pipeline
- frontend agnostic (Queue, PSR-7, [GRPC](https://github.com/agui2200/php-grpc), etc)
- works over TCP, unix sockets and standard pipes
- automatic worker replacement and safe PHP process destruction
- worker lifecycle management (create/allocate/destroy timeouts)
- payload context and body
- control over max jobs per worker
- protocol, worker and job level error management (including PHP errors)
- memory leak failswitch
- very fast (~250k rpc calls per second on Ryzen 1700X using 16 threads)
- works on Windows

Example:
--------

```php
<?php
ini_set('display_errors', 'stderr');
include "vendor/autoload.php";

$relay = new Spiral\Goridge\StreamRelay(STDIN, STDOUT);
$psr7 = new Spiral\RoadRunner\PSR7Client(new Spiral\RoadRunner\Worker($relay));

while ($req = $psr7->acceptRequest()) {
    try {
        $resp = new \Zend\Diactoros\Response();
        $resp->getBody()->write("hello world");

        $psr7->respond($resp);
    } catch (\Throwable $e) {
        $psr7->getWorker()->error((string)$e);
    }
}
```

Configuration can be located in `.rr.yaml` file:

```yaml
http:
  address: 0.0.0.0:8080
  workers:
    command: "php psr-worker.php"
    pool:
      numWorkers: 4
```

> Read more in [Wiki](https://github.com/agui2200/roadrunner/wiki/PHP-Workers).

License:
--------
The MIT License (MIT). Please see [`LICENSE`](./LICENSE) for more information. Maintained by [SpiralScout](https://spiralscout.com).
