<div align="center">
  <a href="https://koyeb.com">
    <img src="https://www.koyeb.com/static/images/icons/koyeb.svg" alt="Logo" width="80" height="80">
  </a>
  <h3 align="center">Koyeb Serverless Platform</h3>
  <p align="center">
    Deploy a Go gRPC server and gateway on Koyeb
    <br />
    <a href="https://koyeb.com">Learn more about Koyeb</a>
    ·
    <a href="https://koyeb.com/docs">Explore the documentation</a>
    ·
    <a href="https://koyeb.com/tutorials">Discover our tutorials</a>
  </p>
</div>

## About Koyeb and the Go gRPC server and gateway example applications

Koyeb is a developer-friendly serverless platform to deploy apps globally. No-ops, servers, or infrastructure management.  This repository contains two companion applications: A gRPC API application and a gateway that can be used to accept and translate HTTP requests.

This example application is designed to show how you can create a gRPC API and make it accessible to web clients by deploying it on Koyeb.  You can find out more about these two applications in the [associated tutorial](https://www.koyeb.com/tutorials/build-a-grpc-api-using-go-and-grpc-gateway).

## Getting Started

Follow the steps below to deploy and run the gRCP API server (`orders-service`) and the HTTP gateway (`gateway-service`) applications on your Koyeb account.

### Requirements

You need a Koyeb account to successfully deploy and run this application. If you don't already have an account, you can sign-up for free [here](https://app.koyeb.com/auth/signup).

### Deploy using the Koyeb button

The fastest way to deploy the two applications is to click the **Deploy to Koyeb** buttons below.

#### Deploy the orders-service gRPC API

First, deploy the gRPC API service with the following button:

[![Deploy to Koyeb](https://www.koyeb.com/static/images/deploy/button.svg)](https://app.koyeb.com/deploy?name=orders-service&type=git&repository=koyeb%2Fexample-go-grpc-gateway&branch=main&builder=dockerfile&target=orders-service&ports=50051%3Bhttp%3B%2F)

#### Deploy the HTTP gateway

Afterwards, deploy the HTTP gateway service with this button:

[![Deploy to Koyeb](https://www.koyeb.com/static/images/deploy/button.svg)](https://app.koyeb.com/deploy?name=gateway-service&type=git&repository=koyeb%2Fexample-go-grpc-gateway&branch=main&builder=dockerfile&args=ORDER_SERVICE_ADDRESS&target=gateway-service&env%5BORDER_SERVICE_ADDRESS%5D=orders-service.orders-service.koyeb%3A50051&ports=8080%3Bhttp%3B%2F)

The required environment variable for the HTTP gateway application is:

* `ORDER_SERVICE_ADDRESS`: The internal service address where the `orders-service` application is deployed.  If using the deploy button above, this will be: `https://orders-service.orders-service.koyeb:50051`.  Be sure to modify it if you deployed with a different app or service name.

Clicking on these buttons brings you to the Koyeb App creation page with everything pre-set to launch the applications.  Fill out the required information in the environment variables and then click **Apply** to launch the applications.

_To modify this application example, you will need to fork this repository. Check out the [fork and deploy](#fork-and-deploy-to-koyeb) instructions._

### Fork and deploy to Koyeb

If you want to customize and enhance this application, you need to fork this repository.

If you used the **Deploy to Koyeb** button, you can simply link your service to your forked repository to be able to push changes.  Alternatively, you can manually create the application as described below.

On the [Koyeb Control Panel](//app.koyeb.com/apps), on the **Overview** tab, click the **Create Web Service** button to begin.

For the `orders-service` application:

1. Select **GitHub** as the deployment method.
2. Select your project from the GitHub repository list and choose the appropriate branch.
3. In the **Builder** section, select **Dockerfile**.  Click the **override** toggle associated with the **Target** option and enter `orders-service` in the field.
4. Expand the **Exposed ports** sections and set the port to **50051**.  De-select the **Public** option to mark the service as internal.
5. Set the App and Service name to `orders-service` and click **Deploy**.

Next, deploy the `gateway-service` application with the following steps:

1. Select **GitHub** as the deployment method.
2. Select your project from the GitHub repository list and choose the appropriate branch.
3. In the **Builder** section, select **Dockerfile**.  Click the **override** toggle associated with the **Target** option and enter `gateway-service` in the field.  Click the **override** toggle associated with the **Args** option and set the value to `["ORDER_SERVICE_ADDRESS"]`.
4. Expand **Environment variable** section and add a new environment variable called `ORDER_SERVICE_ADDRESS` with the private address where your order service can be reached.  It should follow this format: `<SERVICE_NAME>.<APP_NAME>.koyeb:50051`.  If you followed the above steps, this should be: `orders-service.orders-service.koyeb:50051`.
5. Expand the **Exposed ports** section and set the port to **8080**.
8. Set the App name to `gateway-service`. This determines where the application will be deployed to. For example, `https://gateway-service-<YOUR_USERNAME>.koyeb.app`.
9. Click **Deploy** to begin the deployment process.

By visiting the URL for the HTTP gateway you should be able to access the API.

## Contributing

If you have any questions, ideas or suggestions regarding this application sample, feel free to open an [issue](//github.com//koyeb/example-go-grpc-gateway/issues) or fork this repository and open a [pull request](//github.com/koyeb/example-go-grpc-gateway/pulls).

## Contact

[Koyeb](https://www.koyeb.com) - [@gokoyeb](https://twitter.com/gokoyeb) - [Slack](http://slack.koyeb.com/)
