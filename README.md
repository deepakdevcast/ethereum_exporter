<a id="readme-top"></a>

[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">Ethereum Exporter</h3>

  <p align="center">
    An Ethereum Exporter to generate prometheus metrics
    <br />
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about">About</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About

This project is build to generate prometheus metrics for <b>multiple</b> ethereum rpc clients.

Here's why:
* This will helps you to monitor ethereum and get useful metrics
* You can scape these metrics using prometheus
* And eventually use grafana dashboard.


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites
* go

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo
   ```sh
   git clone https://github.com/deepakdevcast/ethereum_exporter.git
   ```
2. initialize .env | Config clients
   ```sh
    # the prometheus metrics endpoint
    METRICS_PORT=':7201'
    <!-- {[CLIENT1],[CLIENT2],...]} -->
    CLIENT_INFO='[{"client":"local","url":"https://localhost:8545"}]'
   ```
3. Run
   ```sh
    cd cmd/ethereum_exporter
    go run .
   ```

<!-- CONTACT -->
## Contact

Deepak Kumar - [@deepakkdev](https://www.linkedin.com/in/deepakkdev/)
