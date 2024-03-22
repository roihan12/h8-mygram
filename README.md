<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<div id="top"></div>

<!-- PROJECT LOGO -->
<div align="center">
<!--  mengarah ke repo  -->
  <a href="https://h8-mygram.fly.dev">
    <h1>MYGRAM API</h1>
  </a>
  <br/>
    <a href="https://www.codacy.com/gh/helmimuzkr/campyuk-be/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=helmimuzkr/campyuk-be&amp;utm_campaign=Badge_Grade">

  </a>
  <br/>
  <br/>
  <a href="https://h8-mygram.fly.dev">Go to API »</a>
</div>

# MYGRAM API

MyGram is an application where you can store photos and make comments on other people's photos. This application will be equipped with CRUD (Create, Read, Update, Delete) operations with tables and the following flow:

- Create: Users can upload their photos to the platform and add accompanying comments.
- Read: Users can view their own uploaded photos along with comments, as well as browse through photos uploaded by other users and their respective comments.
- Update: Users can edit the captions or details of their uploaded photos, as well as modify their comments.
- Delete: Users can delete their uploaded photos, as well as remove their comments from any photo.

This CRUD functionality will provide users with a seamless experience for managing their photos and interactions within the application.

# Features

## User:

- Register
- Login
- Get profile
- Edit profile
- Delete profile

<div>

<details>

| Feature User | Endpoint  | Param | JWT Token | Function                         |
| ------------ | --------- | ----- | --------- | -------------------------------- |
| POST         | /register | -     | NO        | Register new users (account).    |
| POST         | /login    | -     | NO        | Log in into account.             |
| GET          | /users    | -     | YES       | Get account information details. |
| PUT          | /users    | -     | YES       | Edit account details.            |
| DELETE       | /users    | -     | YES       | Delete account.                  |

</div>

## Photos :

- Add new Photo
- Show all list photos
- Edit photo
- Show detail photo
- Delete photo

<div>

<details>

| Feature Photos | Endpoint | Param    | JWT Token | Function           |
| -------------- | -------- | -------- | --------- | ------------------ |
| POST           | /photos  | -        | YES       | Add new photo.     |
| GET            | /photos  | -        | YES       | Get all photos.    |
| PUT            | /photos  | PHOTO ID | YES       | Edit photo.        |
| GET            | /photos  | PHOTO ID | YES       | Get photos details |
| DELETE         | /photos  | PHOTO ID | YES       | Delete photo       |

</details>

</div>

## Comments :

- Add new comment
- Show all my list comments
- Edit comment
- Show detail comment
- Delete comment

<div>

<details>

| Feature Comments | Endpoint  | Param      | JWT Token | Function            |
| ---------------- | --------- | ---------- | --------- | ------------------- |
| POST             | /comments | -          | YES       | Add new comment.    |
| GET              | /comments | -          | YES       | Get all my comment. |
| PUT              | /comments | COMMENT ID | YES       | Edit comment.       |
| GET              | /comments | COMMENT ID | YES       | Get comment details |
| DELETE           | /comments | COMMENT ID | YES       | Delete comment      |

</details>

</div>

## Social Medias :

- Add new social media
- Show all list social media
- Edit social media
- Show detail social media
- Delete social media

<div>

<details>

| Feature Comments | Endpoint      | Param          | JWT Token | Function                 |
| ---------------- | ------------- | -------------- | --------- | ------------------------ |
| POST             | /socialmedias | -              | YES       | Add new social media.    |
| GET              | /socialmedias | -              | YES       | Get all social media.    |
| PUT              | /socialmedias | SOCIALMEDIA ID | YES       | Edit social media.       |
| GET              | /socialmedias | SOCIALMEDIA ID | YES       | Get social media details |
| DELETE           | /socialmedias | SOCIALMEDIA ID | YES       | Delete social media      |

</details>

</div>

# API Documentation

[![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)](https://h8-mygram.fly.dev/docs/index.html) [![Postman Collection](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://documenter.getpostman.com/view/19247831/2sA35A966B)

# ERD

![ERD](https://res.cloudinary.com/drugrg7xz/image/upload/v1711121846/mygram_-_public_crmm6w.png "ERD")

# Run Locally

1. Clone the project

    ```bash
    $ git clone https://github.com/roihan12/h8-mygram.git
    ```

2. Create new database

3. Go to project directory

    ```bash
    $ cd mygram
    ```

4. Dont forget to activate the credential for third party api cloudinary,
5. Create env using env.example format and fill in all value
6. Download all packages and dependencies
    ```bash
    $ go mod tidy
    ```
7. Run the program
    ```bash
    $ go run .
    ```
8. Enjoy

### Built With

## 🛠 Tools

**Backend:** <br>
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

**Deployment:** <br>
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)

**Communication:**  
![GitHub](https://img.shields.io/badge/github%20Project-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

