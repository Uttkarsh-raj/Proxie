# Proxie
<br>
<center>
Proxie, a proxy server in Golang, shields clients from direct server exposure. It features advanced logging for comprehensive network traffic monitoring and analysis, crucial for troubleshooting and security auditing. Additionally, Proxie's rate-limiting mechanism ensures fair resource allocation, enhancing server performance and stability while mitigating the risk of abuse and vulnerabilities.
</center>
<br><br>

<!--TABLE OF CONTENTS-->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a> 
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a> 
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
  </details>

<!--About the Project-->
  
## About The Project
<center>
Proxie is a simple proxy server made in golang which prevents the client getting exposed to the server. Additionally , Proxie offers advanced logging functionalities, enabling administrators to monitor and analyze network traffic comprehensively. This feature is crucial for troubleshooting, compliance, and security auditing purposes.Proxie implements a rate-limiting mechanism to prevent abuse and ensure fair resource allocation. This feature enhances server performance and stability, mitigating the risk of service disruptions and security vulnerabilities.</center>
  

### Built With
  - **Golang** - An open-source programming language
  - **Gin-Gonic** - A web framework written in Go
<br><br>

<img height="100px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"/>



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!--GETTING STARTED-->

## Getting Started

To get started with your Golang application, follow these steps:

1. **Install Golang**: Download and install Golang from the [official website](https://golang.org/dl/).

2. **Set Up Your Workspace**: Create a new directory for your project and set your `GOPATH` environment variable to point to this directory.

3. **Initialize Your Project**: Inside your project directory, run the following command to initialize a new Go module:

   ```
   go mod init github.com/your-username/project-name
   ```
   After installing Golang, you can start running your Go project.
4. **Run without Debugging**: In your terminal, navigate to the directory containing your main Go file (usually named `main.go`). Then, run the following command to build and execute your Go application:
   ```
   go run main.go
   ```
   This command will compile and execute your Go program without generating a binary file.



### Installation 

Below is an example of how you can instruct your audience on installing and setting up your app.This template doesn't rely on any external dependencies or services.

1. Clone the repo 
  ```
  git clone https://github.com/Uttkarsh-raj/MiniCache
  ```

2. Install the packages 
  ```
  go mod tidy
  ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Routes

- **Get "/"**
  * Request as
  ```
  curl -v http://localhost:7000/?url=<your_url_here>
  ```
  * Response as
      - Sucess : Respective output
      - Failure: "error" : Error
  
<!--USAGE EXAMPLES-->

## Usage
The Proxie proxy server project offers a versatile solution for various networking scenarios where proxy functionality is required. Some potential use cases include:

1. **Anonymity and Privacy**:  Proxie can be used to hide the client's IP address from the server, providing anonymity and privacy.
2. **Rate Limiting**: Limits the number of requests a client can make to the server within a specified timeframe, preventing excessive traffic that could lead to service disruptions or security vulnerabilities.
3. **Monitoring and Logging**: Proxie's logging feature enables administrators to monitor network traffic and analyze usage patterns for auditing and troubleshooting.

With its customizable features and easy integration, Proxie can serve as a reliable proxy server solution for a variety of networking needs.

## Screenshots:
<br>
<center>
<img width="1000" src="https://github.com/Uttkarsh-raj/Proxie/assets/106571927/3acaa7b6-464d-4ad4-97f9-0a6fad8a2607"></img>
<br> 
<img width="1000" src="https://github.com/Uttkarsh-raj/Proxie/assets/106571927/0861ba55-5b71-4868-9f9b-7c492faf4057"></img>
<img width="1000" src="https://github.com/Uttkarsh-raj/Proxie/assets/106571927/2df4d53c-ddd6-460f-811a-499491d0fd12"></img>
<br>
<img width="1000" src="https://github.com/Uttkarsh-raj/Proxie/assets/106571927/bf2b36cc-f2cd-4837-8182-8383867c5c99"></img>
</center>
<br>
<!-- ROADMAP -->

## Roadmap

- [x] Add Changelog
- [x] Add back to top links
- [x] Add Additional Templates w/ Examples
- [ ] Add "components" document to easily copy & paste sections of the readme
- [ ] Multi-language Support
  - [ ] Hindi
  - [ ] English

  
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!--CONTRIBUTING-->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire ,and create.Any contributions you make are *greatly appreciated*.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact
Uttkarsh Raj - https://github.com/Uttkarsh-raj <br>

Project Link: https://github.com/Uttkarsh-raj/MiniCache

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
- [Malven's Grid Cheatsheet](https://grid.malven.co/)
- [Img Shields](https://shields.io)
- [GitHub Pages](https://pages.github.com)
- [Font Awesome](https://fontawesome.com)
- [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
