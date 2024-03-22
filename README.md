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

<p align="right">(<a href="#readme-top">back to top</a>)</p>

- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
- [Malven's Grid Cheatsheet](https://grid.malven.co/)
- [Img Shields](https://shields.io)
- [GitHub Pages](https://pages.github.com)
- [Font Awesome](https://fontawesome.com)
- [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: docs/ERD.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com
[Golang]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Golang-url]: https://laravel.com
